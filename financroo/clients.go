package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	obc "github.com/cloudentity/openbanking-sandbox/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"

	acpclient "github.com/cloudentity/acp-client-go"
)

func NewAcpClient(c Config, cfg BankConfig) (acpclient.Client, error) {
	var (
		issuerURL, authorizeURL, redirectURL *url.URL
		client                               acpclient.Client
		err                                  error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", c.ACPURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s/api/callback", c.UIURL)); err != nil {
		return client, err
	}

	requestObjectExpiration := time.Minute * 10
	config := acpclient.Config{
		ClientID:                    cfg.AcpClient.ClientID,
		IssuerURL:                   issuerURL,
		AuthorizeURL:                authorizeURL,
		RedirectURL:                 redirectURL,
		RequestObjectSigningKeyFile: cfg.AcpClient.KeyFile,
		RequestObjectExpiration:     &requestObjectExpiration,
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

func NewLoginClient(c Config) (acpclient.Client, error) {
	var (
		issuerURL *url.URL
		client    acpclient.Client
		err       error
	)

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, c.Login.TenantID, c.Login.ServerID)); err != nil {
		return client, err
	}

	config := acpclient.Config{
		ClientID:  c.Login.ClientID,
		IssuerURL: issuerURL,
		Timeout:   c.Login.Timeout,
		RootCA:    c.Login.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
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
