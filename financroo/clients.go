package main

import (
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
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
)

func NewAcpClient(c Config, cfg BankConfig) (acpclient.Client, error) {
	var (
		tokenURL, issuerURL, authorizeURL, userinfoURL, redirectURL *url.URL
		client                                                      acpclient.Client
		err                                                         error
	)

	if tokenURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/token", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", c.ACPURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if userinfoURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/userinfo", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s/api/callback", c.UIURL)); err != nil {
		return client, err
	}

	config := acpclient.Config{
		ClientID:                    cfg.AcpClient.ClientID,
		IssuerURL:                   issuerURL,
		TokenURL:                    tokenURL,
		AuthorizeURL:                authorizeURL,
		UserinfoURL:                 userinfoURL,
		RedirectURL:                 redirectURL,
		RequestObjectSigningKeyFile: cfg.AcpClient.KeyFile,
		Scopes:                      []string{"accounts", "openid", "offline_access"},
		Timeout:                     cfg.AcpClient.Timeout,
		CertFile:                    cfg.AcpClient.CertFile,
		KeyFile:                     cfg.AcpClient.KeyFile,
		RootCA:                      cfg.AcpClient.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
}

type LoginClientConfig struct {
	RootCA      string
	Timeout     time.Duration
	UserinfoURL string
}

func (c *Config) ToLoginClientConfig() LoginClientConfig {
	return LoginClientConfig{
		RootCA:      c.Login.RootCA,
		Timeout:     c.Login.Timeout,
		UserinfoURL: fmt.Sprintf("%s/%s/%s/userinfo", c.ACPInternalURL, c.Login.TenantID, c.Login.ServerID),
	}
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
