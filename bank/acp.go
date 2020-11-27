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
	httptransport "github.com/go-openapi/runtime/client"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type AcpClient struct {
	client *client.Web
	tenant string
	server string
}

func NewAcpClient(config Config) (AcpClient, error) {
	var (
		acpClient = AcpClient{}
		hc        *http.Client
		err       error
	)

	if hc, err = newHTTPClient(config); err != nil {
		return acpClient, err
	}
	parts := strings.Split(config.TokenURL.Path, "/")
	if len(parts) < 2 {
		return acpClient, errors.New("can't get tenant/server from issuer url")
	}
	acpClient.tenant = parts[1]
	acpClient.server = parts[2]

	cc := clientcredentials.Config{
		ClientID:  config.ClientID,
		Scopes:    []string{"introspect_openbanking_tokens"},
		TokenURL:  config.TokenURL.String(),
		AuthStyle: oauth2.AuthStyleInParams,
	}

	acpClient.client = client.New(httptransport.NewWithClient(
		config.TokenURL.Host,
		"/",
		[]string{config.TokenURL.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return acpClient, nil
}

func (a *AcpClient) Introspect(token string) (*models.IntrospectOpenbankingAccountAccessConsentResponse, error) {
	var (
		resp *openbanking.OpenbankingAccountAccessConsentIntrospectOK
		err  error
	)

	if resp, err = a.client.Openbanking.OpenbankingAccountAccessConsentIntrospect(openbanking.NewOpenbankingAccountAccessConsentIntrospectParams().
		WithTid(a.tenant).WithAid(a.server).WithToken(&token), nil); err != nil {
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
