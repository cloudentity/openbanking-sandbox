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

	obc "github.com/cloudentity/acp/pkg/openbanking/client"
	"github.com/cloudentity/acp/pkg/swagger/client"
	"github.com/cloudentity/acp/pkg/swagger/client/openbanking"
	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/dgrijalva/jwt-go"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type AcpClient struct {
	client *client.Web
	tenant string
	server string
}

func NewAcpMTLSClient(config Config) (AcpClient, error) {
	var (
		acpClient = AcpClient{}
		hc        *http.Client
		err       error
	)

	parts := strings.Split(config.IssuerURL.Path, "/")
	if len(parts) < 2 {
		return acpClient, errors.New("can't get tenant/server from issuer url")
	}
	acpClient.tenant = parts[1]
	acpClient.server = parts[2]

	if hc, err = newHttpClient(config); err != nil {
		return acpClient, err
	}

	cc := clientcredentials.Config{
		ClientID:  config.ClientID,
		Scopes:    []string{"accounts"},
		TokenURL:  fmt.Sprintf("%s/oauth2/token", config.IssuerURL.String()),
		AuthStyle: oauth2.AuthStyleInParams,
	}

	acpClient.client = client.New(httptransport.NewWithClient(
		config.IssuerURL.Host,
		"/",
		[]string{config.IssuerURL.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return acpClient, nil
}

func (a *AcpClient) RegisterAccountAccessConsent(permissions []string) (string, error) {
	var (
		response *openbanking.CreateAccountAccessConsentRequestCreated
		request  = &models.AccountAccessConsentRequest{
			Data: &models.AccountAccessConsentRequestData{
				Permissions: permissions,
			},
		}
		err error
	)

	if response, err = a.client.Openbanking.CreateAccountAccessConsentRequest(openbanking.NewCreateAccountAccessConsentRequestParams().
		WithTid(a.tenant).WithAid(a.server).WithRequest(request), nil); err != nil {
		return "", err
	}

	return response.Payload.Data.ConsentID, nil
}

type AcpWebClient struct {
	Config
	httpClient *http.Client
	signingKey interface{}
}

func NewAcpWebClient(config Config) (client AcpWebClient, err error) {
	if client.httpClient, err = newHttpClient(config); err != nil {
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
	IdToken  map[string]Claim `json:"id_token"`
	Userinfo map[string]Claim `json:"userinfo"`
}

func SignRequest(request Request, signingKey interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, request)
	return token.SignedString(signingKey)
}

func (a *AcpWebClient) AuthorizeURL(intentID string, challenge string, scopes []string, nonce string) (string, error) {
	var (
		buf           bytes.Buffer
		signedRequest string
		err           error
	)

	request := Request{
		ClientID:     a.ClientID,
		Scope:        strings.Join(scopes, " "),
		ResponseType: "code",
		RedirectURI:  a.RedirectURL.String(),
		Nonce:        nonce,
		Claims: Claims{
			Userinfo: map[string]Claim{
				"openbanking_intent_id": {
					Essential: true,
					Value:     intentID,
				},
			},
			IdToken: map[string]Claim{
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
		"client_id":             {a.ClientID},
		"redirect_uri":          {a.RedirectURL.String()},
		"scope":                 {strings.Join(scopes, " ")},
		"code_challenge":        {challenge},
		"code_challenge_method": {"S256"},
		"request":               {signedRequest},
	}

	authorizeURL := fmt.Sprintf("%s/oauth2/authorize", a.IssuerURL.String())
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
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	Scope       string  `json:"scope"`
	ExpiresIn   int     `json:"expires_in"`
	IDToken     *string `json:"id_token"`
}

func (a *AcpWebClient) Exchange(code string, verifier string) (token Token, err error) {
	var (
		response *http.Response
		body     []byte
	)

	values := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {a.ClientID},
		"redirect_uri":  {a.RedirectURL.String()},
		"code_verifier": {verifier},
	}

	if response, err = a.httpClient.PostForm(fmt.Sprintf("%s/oauth2/token", a.IssuerURL.String()), values); err != nil {
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

func (a *AcpWebClient) Userinfo(token string) (body map[string]interface{}, err error) {
	var (
		request  *http.Request
		response *http.Response
		bs       []byte
	)

	if request, err = http.NewRequest("GET", fmt.Sprintf("%s/userinfo", a.Config.IssuerURL.String()), nil); err != nil {
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

func newHttpClient(config Config) (*http.Client, error) {
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
				RootCAs:      pool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}, nil
}

type OpenbankingClient struct {
	*obc.Openbanking
}

func NewOpenbankingClient(config Config) OpenbankingClient {
	var (
		c  = OpenbankingClient{}
		hc = &http.Client{}
	)

	c.Openbanking = obc.New(httptransport.NewWithClient(
		config.BankURL.Host,
		"/",
		[]string{config.BankURL.Scheme},
		hc,
	), nil)

	return c
}
