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

	"github.com/cloudentity/acp/pkg/swagger/client"
	"github.com/cloudentity/acp/pkg/swagger/client/openbanking"
	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/dgrijalva/jwt-go"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

	parts := strings.Split(config.TokenURL.Path, "/")
	if len(parts) < 2 {
		return acpClient, errors.New("can't get tenant/server from token url")
	}
	acpClient.tenant = parts[1]
	acpClient.server = parts[2]

	if hc, err = newHttpClient(config); err != nil {
		return acpClient, err
	}

	cc := clientcredentials.Config{
		ClientID:  config.ClientID,
		Scopes:    []string{"accounts"},
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

func (a *AcpWebClient) AuthorizeURL(intentID string, challenge string, scopes []string, nonce string) (string, error) {
	var (
		buf           bytes.Buffer
		signedRequest string
		err           error
	)

	request := Request{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
		},
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
	logrus.Infof("request: %+v", request)

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

	buf.WriteString(a.AuthURL.String())
	if strings.Contains(a.AuthURL.String(), "?") {
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

	if response, err = a.httpClient.PostForm(a.TokenURL.String(), values); err != nil {
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
