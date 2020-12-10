package acpclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/cloudentity/acp-client-go/client"
	o2 "github.com/cloudentity/acp-client-go/client/oauth2"
	"github.com/cloudentity/acp-client-go/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

// Client provides a client to the ACP API
type Client struct {
	*client.Acp
	c *http.Client

	// Tenant id read from the IssuerURL
	TenantID string

	// Authorization server id read from the IssuerURL
	ServerID string
}

// ACP client configuration
type Config struct {
	// ClientID is the application's ID.
	ClientID string

	// ClientSecret is the application's secret.
	ClientSecret string

	// IssuerURL is the authorization server's url
	// example: https://localhost:8443/default/default
	IssuerURL *url.URL

	// Scope specifies optional requested permissions.
	Scopes []string

	// Default HttpClient timeout.
	// Ignored if HttpClient is provided.
	Timeout time.Duration

	// Optional path to the file with certificate for tls authentication.
	// Ignored if HttpClient is provided.
	CertFile string

	// Optional path to the file with private key for tls authentication.
	// Ignored if HttpClient is provided.
	KeyFile string

	// Optional path to the file with root CAs.
	// Ignored if HttpClient is provided.
	RootCA string

	// HttpClient is the client to use. Default will be used if not provided.
	HttpClient *http.Client
}

func (c *Config) newHTTPClient() (*http.Client, error) {
	var (
		pool *x509.CertPool
		cert tls.Certificate
		data []byte
		err  error
	)

	if cert, err = tls.LoadX509KeyPair(c.CertFile, c.KeyFile); err != nil {
		return nil, errors.Wrapf(err, "failed to read certificate and private key")
	}

	if pool, err = x509.SystemCertPool(); err != nil {
		return nil, errors.Wrapf(err, "failed to read system root CAs")
	}

	if c.RootCA != "" {
		if data, err = ioutil.ReadFile(c.RootCA); err != nil {
			return nil, fmt.Errorf("failed to read http client root ca: %w", err)
		}

		pool.AppendCertsFromPEM(data)
	}

	return &http.Client{
		Timeout: c.Timeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			ResponseHeaderTimeout: c.Timeout,
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

func New(cfg Config) (c Client, err error) {
	paths := strings.Split(cfg.IssuerURL.Path, "/")

	if len(paths) < 2 {
		return c, errors.New("invalid issuer url")
	}

	c.TenantID = paths[1]
	c.ServerID = paths[2]

	if c.c, err = cfg.newHTTPClient(); err != nil {
		return c, err
	}

	cc := clientcredentials.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       cfg.Scopes,
		TokenURL:     fmt.Sprintf("%s/oauth2/token", cfg.IssuerURL),
	}

	c.Acp = client.New(httptransport.NewWithClient(
		cfg.IssuerURL.Host,
		"/",
		[]string{cfg.IssuerURL.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, c.c)),
	), nil)

	return c, nil
}

func (c *Client) IntrospectToken(token string) (*models.IntrospectResponse, error) {
       var (
               resp *o2.IntrospectOK
               err  error
       )

       if resp, err = c.Oauth2.Introspect(o2.NewIntrospectParams().
               WithTid(c.TenantID).
               WithAid(c.ServerID).
               WithToken(&token), nil); err != nil {
               return nil, err
       }

       return resp.Payload, nil
}
