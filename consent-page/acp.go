package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/cloudentity/acp/pkg/swagger/client"
	"github.com/cloudentity/acp/pkg/swagger/client/openbanking"
	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/gin-gonic/gin"
	httptransport "github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type LoginRequest struct {
	ID    string
	State string
}

func NewLoginRequest(c *gin.Context) LoginRequest {
	return LoginRequest{ID: c.Query("login_id"), State: c.Query("login_state")}
}

func (l *LoginRequest) Validate() error {
	if l.ID == "" || l.State == "" {
		return errors.New("login_id / login_state missing")
	}

	return nil
}

type AcpClient struct {
	client *client.Web
	tenant string
}

func NewAcpClient(config Config) (AcpClient, error) {
	var (
		acpClient = AcpClient{}
		hc        *http.Client
		err       error
	)

	parts := strings.Split(config.IssuerURL.Path, "/")
	if len(parts) == 0 {
		return acpClient, errors.New("can't get tenant from issuer url")
	}
	acpClient.tenant = parts[1]

	if hc, err = newHttpClient(config); err != nil {
		return acpClient, err
	}

	cc := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"manage_openbanking_consents"},
		TokenURL:     fmt.Sprintf("%s/oauth/token", config.IssuerURL.String()),
	}

	acpClient.client = client.New(httptransport.NewWithClient(
		config.IssuerURL.Host,
		"/",
		[]string{config.IssuerURL.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return acpClient, nil
}

func (a *AcpClient) GetAccountAccessConsent(r LoginRequest) (*models.GetAccountAccessConsentResponse, error) {
	var (
		response *openbanking.GetAccountAccessConsentSystemOK
		err      error
	)

	if response, err = a.client.Openbanking.GetAccountAccessConsentSystem(openbanking.NewGetAccountAccessConsentSystemParams().
		WithTid(a.tenant).WithLoginID(r.ID), nil); err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (a *AcpClient) ApproveAccountAccessConsent(r LoginRequest, accountIDs []string, grantScopes []string) (string, error) {
	var (
		response *openbanking.AcceptAccountAccessConsentSystemOK
		err      error
	)

	accept := &models.AcceptAccountAccessConsentRequest{
		GrantedScopes: grantScopes,
		AccountIDs:    accountIDs,
		LoginState:    r.State,
	}

	if response, err = a.client.Openbanking.AcceptAccountAccessConsentSystem(openbanking.NewAcceptAccountAccessConsentSystemParams().
		WithTid(a.tenant).WithLoginID(r.ID).WithAcceptAccountAccessConsent(accept), nil); err != nil {
		return "", err
	}

	return response.Payload.RedirectTo, nil
}

func (a *AcpClient) RejectAccountAccessConsent(r LoginRequest, rejectReason string, rejectStatusCode int64) (string, error) {
	var (
		response *openbanking.RejectAccountAccessConsentSystemOK
		err      error
	)

	reject := &models.RejectAccountAccessConsentRequest{
		ID:         r.ID,
		LoginState: r.State,
		Error:      rejectReason,
		StatusCode: rejectStatusCode,
	}

	if response, err = a.client.Openbanking.RejectAccountAccessConsentSystem(openbanking.NewRejectAccountAccessConsentSystemParams().
		WithTid(a.tenant).WithLoginID(r.ID).WithRejectAccountAccessConsent(reject), nil); err != nil {
		return "", nil
	}

	return response.Payload.RedirectTo, nil
}

func newHttpClient(config Config) (*http.Client, error) {
	var (
		pool *x509.CertPool
		data []byte
		err  error
	)

	if pool, err = x509.SystemCertPool(); err != nil {
		return nil, err
	}

	if config.RootCA != "" {
		if data, err = ioutil.ReadFile(config.RootCA); err != nil {
			return nil, fmt.Errorf("failed to read http client root ca: %+v", err)
		}

		pool.AppendCertsFromPEM(data)
	}

	return &http.Client{
		Timeout: config.Timeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			ResponseHeaderTimeout: config.Timeout,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.InsecureSkipVerify, // nolint
				RootCAs:            pool,
			},
		},
	}, nil
}
