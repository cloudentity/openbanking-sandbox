package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"

	obc "github.com/cloudentity/openbanking-sandbox/client"
	"github.com/cloudentity/openbanking-sandbox/client/account_access"
	"github.com/cloudentity/openbanking-sandbox/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type AcpAccountAccessClient struct {
	*obc.Openbanking
}

func NewAcpAccountAccessClient(config SystemClientConfig) (AcpAccountAccessClient, error) {
	var (
		c   = AcpAccountAccessClient{}
		err error
		hc  *http.Client
		u   *url.URL
	)

	if hc, err = newHTTPClient(config.HTTPClientConfig); err != nil {
		return c, err
	}

	cc := clientcredentials.Config{
		ClientID:  config.ClientID,
		Scopes:    []string{"accounts"},
		TokenURL:  config.TokenURL,
		AuthStyle: oauth2.AuthStyleInParams,
	}

	if u, err = url.Parse(config.TokenURL); err != nil {
		return c, errors.Wrapf(err, "failed to parse token url")
	}

	c.Openbanking = obc.New(httptransport.NewWithClient(
		u.Host,
		fmt.Sprintf("/%s/%s/open-banking/v3.1/aisp", config.TenantID, config.ServerID),
		[]string{u.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return c, nil
}

func (a *AcpAccountAccessClient) RegisterAccountAccessConsent(permissions []string) (*models.OBReadConsentResponse1, error) {
	var (
		response *account_access.CreateAccountAccessConsentsCreated
		request  = &models.OBReadConsent1{
			Data: &models.OBReadConsent1Data{
				Permissions: permissions,
			},
		}
		err error
	)

	if response, err = a.Openbanking.AccountAccess.CreateAccountAccessConsents(account_access.NewCreateAccountAccessConsentsParams().
		WithOBReadConsent1Param(request), nil); err != nil {
		return nil, err
	}

	return response.Payload, nil
}

type LoginClient struct {
	Config     LoginClientConfig
	httpClient *http.Client
}

func NewLoginClient(config LoginClientConfig) (LoginClient, error) {
	var (
		client = LoginClient{}
		pool   *x509.CertPool
		data   []byte
		err    error
	)

	if pool, err = x509.SystemCertPool(); err != nil {
		return client, err
	}

	if config.RootCA != "" {
		if data, err = ioutil.ReadFile(config.RootCA); err != nil {
			return client, fmt.Errorf("failed to read http client root ca: %w", err)
		}

		pool.AppendCertsFromPEM(data)
	}

	client.httpClient = &http.Client{
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
				RootCAs:    pool,
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	client.Config = config

	return client, nil
}

func (a *LoginClient) Userinfo(token string) (body map[string]interface{}, err error) {
	var (
		request  *http.Request
		response *http.Response
		bs       []byte
	)

	if request, err = http.NewRequest("GET", a.Config.UserinfoURL, nil); err != nil {
		return body, fmt.Errorf("error while building request: %w", err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	if response, err = a.httpClient.Do(request); err != nil {
		return body, fmt.Errorf("error while calling userinfo endpoint: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return body, errors.New(fmt.Sprintf("ACP responded with status code: %v", response.Status))
	}

	if bs, err = ioutil.ReadAll(response.Body); err != nil {
		return body, fmt.Errorf("error during decoding exchange body: %w", err)
	}

	if err = json.Unmarshal(bs, &body); err != nil {
		return body, err
	}

	return body, nil
}

type Token struct {
	AccessToken  string
	RefreshToken string
}

func newHTTPClient(config HTTPClientConfig) (*http.Client, error) {
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
				Certificates: []tls.Certificate{cert},
				MinVersion:   tls.VersionTLS12,
			},
		},
	}, nil
}

type OpenbankingClient struct {
	*obc.Openbanking
}

func NewOpenbankingClient(config BankConfig) (OpenbankingClient, error) {
	var (
		c   = OpenbankingClient{}
		hc  = &http.Client{}
		u   *url.URL
		err error
	)

	if u, err = url.Parse(config.URL); err != nil {
		return c, errors.Wrapf(err, "failed to parse bank url")
	}

	c.Openbanking = obc.New(httptransport.NewWithClient(
		u.Host,
		"/",
		[]string{u.Scheme},
		hc,
	), nil)

	return c, nil
}
