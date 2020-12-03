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

	httptransport "github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/cloudentity/acp/pkg/swagger/client"
	o2 "github.com/cloudentity/acp/pkg/swagger/client/oauth2"
	"github.com/cloudentity/acp/pkg/swagger/client/openbanking"
	"github.com/cloudentity/acp/pkg/swagger/models"
)

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

	parts := strings.Split(config.SystemTokenURL.Path, "/")
	if len(parts) == 0 {
		return acpClient, errors.New("can't get tenant from issuer url")
	}
	acpClient.tenant = parts[1]

	if hc, err = newHTTPClient(config); err != nil {
		return acpClient, err
	}

	cc := clientcredentials.Config{
		ClientID:     config.SystemClientID,
		ClientSecret: config.SystemClientSecret,
		Scopes:       []string{"manage_openbanking_consents"},
		TokenURL:     config.SystemTokenURL.String(),
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	acpClient.client = client.New(httptransport.NewWithClient(
		config.SystemTokenURL.Host,
		"/",
		[]string{config.SystemTokenURL.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return acpClient, nil
}

func (a *AcpClient) GetAccountAccessConsent(accountIDs []string) (*models.ListAccountAccessConsentsWithClient, error) {
	var (
		response *openbanking.ListConsentsByAccountsOK
		err      error
	)

	request := &models.ListConsentsByAccountsRequest{
		AccountIDs: accountIDs,
	}

	if response, err = a.client.Openbanking.ListConsentsByAccounts(openbanking.NewListConsentsByAccountsParams().
		WithTid(a.tenant).WithListConsentsByAccountsRequest(request), nil); err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (a *AcpClient) RevokeAccountAccessConsent(id string) error {
	var err error

	if _, err = a.client.Openbanking.RevokeOpenbankingConsent(openbanking.NewRevokeOpenbankingConsentParams().
		WithTid(a.tenant).WithConsentID(id), nil); err != nil {
		return err
	}

	return nil
}

type AcpIntrospectClient struct {
	client *client.Web
	tenant string
	server string
}

func NewAcpIntrospectClient(config Config) (AcpIntrospectClient, error) {
	var (
		acpClient = AcpIntrospectClient{}
		hc        *http.Client
		err       error
	)

	parts := strings.Split(config.IntrospectTokenURL.Path, "/")
	if len(parts) < 2 {
		return acpClient, errors.New("can't get tenant/server from token url")
	}
	acpClient.tenant = parts[1]
	acpClient.server = parts[2]

	if hc, err = newHTTPClient(config); err != nil {
		return acpClient, err
	}

	cc := clientcredentials.Config{
		ClientID:     config.IntrospectClientID,
		ClientSecret: config.IntrospectClientSecret,
		Scopes:       []string{"introspect_tokens"},
		TokenURL:     config.IntrospectTokenURL.String(),
	}

	acpClient.client = client.New(httptransport.NewWithClient(
		config.IntrospectTokenURL.Host,
		"/",
		[]string{config.IntrospectTokenURL.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return acpClient, nil
}

func (i *AcpIntrospectClient) IntrospectToken(token string) (*models.IntrospectResponse, error) {
	var (
		resp *o2.IntrospectOK
		err  error
	)

	if resp, err = i.client.Oauth2.Introspect(o2.NewIntrospectParams().
		WithTid(i.tenant).
		WithAid(i.server).
		WithToken(&token), nil); err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func newHTTPClient(config Config) (*http.Client, error) {
	var (
		pool *x509.CertPool
		cert tls.Certificate
		data []byte
		err  error
	)

	if cert, err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile); err != nil {
		return nil, err
	}

	if pool, err = x509.SystemCertPool(); err != nil {
		return nil, err
	}

	if config.RootCA != "" {
		if data, err = ioutil.ReadFile(config.RootCA); err != nil {
			return nil, fmt.Errorf("failed to read http client root ca: %w", err)
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
				RootCAs:      pool,
				MinVersion:   tls.VersionTLS12,
				Certificates: []tls.Certificate{cert},
			},
		},
	}, nil
}
