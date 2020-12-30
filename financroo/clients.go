package main

import (
	"bytes"
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
	"strings"
	"time"

	obc "github.com/cloudentity/openbanking-sandbox/client"
	"github.com/cloudentity/openbanking-sandbox/client/account_access"
	"github.com/cloudentity/openbanking-sandbox/models"
	"github.com/dgrijalva/jwt-go"
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

type AcpWebClient struct {
	Config     WebClientConfig
	httpClient *http.Client
	signingKey interface{}
}

func NewAcpWebClient(config WebClientConfig) (client AcpWebClient, err error) {
	if client.httpClient, err = newHTTPClient(config.HTTPClientConfig); err != nil {
		return client, err
	}

	client.Config = config

	if client.signingKey, err = config.GetSigningKey(); err != nil {
		return client, err
	}

	return client, nil
}

type Request struct {
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Claims       Claims `json:"claims"`
	Nonce        string `json:"nonce"`
}

func (r Request) Valid() error {
	return nil
}

type Claim struct {
	Essential bool   `json:"essential"`
	Value     string `json:"value"`
}

type Claims struct {
	IDToken  map[string]Claim `json:"id_token"`
	Userinfo map[string]Claim `json:"userinfo"`
}

func SignRequest(request Request, signingKey interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, request)
	return token.SignedString(signingKey)
}

func (a *AcpWebClient) LoginURL(intentID string, challenge string, scopes []string, nonce string) (string, error) {
	var (
		buf           bytes.Buffer
		signedRequest string
		err           error
	)

	request := Request{
		ClientID:     a.Config.ClientID,
		Scope:        strings.Join(scopes, " "),
		ResponseType: "code",
		RedirectURI:  a.Config.RedirectURL,
		Nonce:        nonce,
		Claims: Claims{
			Userinfo: map[string]Claim{
				"openbanking_intent_id": {
					Essential: true,
					Value:     intentID,
				},
			},
			IDToken: map[string]Claim{
				"openbanking_intent_id": {
					Essential: true,
					Value:     intentID,
				},
				"acr": {
					Essential: true,
					Value:     "urn:openbanking:psd2:sca",
				},
			},
		},
	}

	if signedRequest, err = SignRequest(request, a.signingKey); err != nil {
		return "", errors.Wrapf(err, "failed to sign request")
	}

	queryParams := url.Values{
		"response_type":         {"code"},
		"client_id":             {a.Config.ClientID},
		"redirect_uri":          {a.Config.RedirectURL},
		"scope":                 {strings.Join(scopes, " ")},
		"code_challenge":        {challenge},
		"code_challenge_method": {"S256"},
		"request":               {signedRequest},
	}

	authorizeURL := a.Config.AuthorizeURL
	buf.WriteString(authorizeURL)
	if strings.Contains(authorizeURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}

	buf.WriteString(queryParams.Encode())
	return buf.String(), nil
}

type Token struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	TokenType    string  `json:"token_type"`
	Scope        string  `json:"scope"`
	ExpiresIn    int     `json:"expires_in"`
	IDToken      *string `json:"id_token"`
}

func (a *AcpWebClient) Exchange(code string, verifier string) (token Token, err error) {
	var (
		response *http.Response
		body     []byte
	)

	values := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {a.Config.ClientID},
		"redirect_uri":  {a.Config.RedirectURL},
		"code_verifier": {verifier},
	}

	if response, err = a.httpClient.PostForm(a.Config.TokenURL, values); err != nil {
		return token, fmt.Errorf("error while obtaining token: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return token, errors.New(fmt.Sprintf("ACP responded with status code: %v", response.Status))
	}

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return token, fmt.Errorf("error during decoding exchange body: %w", err)
	}

	if err = json.Unmarshal(body, &token); err != nil {
		return token, fmt.Errorf("error during parsing token response: %w", err)
	}

	return token, nil
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

func (a *AcpWebClient) RenewToken(refreshToken string) (token Token, err error) {
	var (
		response *http.Response
		body     []byte
	)

	values := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
		"client_id":     {a.Config.ClientID},
	}

	if response, err = a.httpClient.PostForm(a.Config.TokenURL, values); err != nil {
		return token, fmt.Errorf("error while obtaining token: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return token, errors.New(fmt.Sprintf("ACP responded with status code: %v", response.Status))
	}

	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return token, fmt.Errorf("error during decoding exchange body: %w", err)
	}

	if err = json.Unmarshal(body, &token); err != nil {
		return token, fmt.Errorf("error during parsing token response: %w", err)
	}

	return token, nil
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
