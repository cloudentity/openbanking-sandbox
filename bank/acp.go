package main

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudentity/acp/pkg/swagger/client"
	"github.com/cloudentity/acp/pkg/swagger/client/logins"
	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/cloudentity/go-tools/httpx"
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
		u         *url.URL
		err       error
	)

	if u, err = url.Parse(config.TokenURL); err != nil {
		return acpClient, err
	}

	parts := strings.Split(u.Path, "/")
	acpClient.tenant = parts[1]

	if hc, err = httpx.NewClient(httpx.ClientConfig{
		Timeout:            config.Timeout,
		InsecureSkipVerify: config.InsecureSkipVerify,
		RootCA:             config.RootCA,
	}); err != nil {
		return acpClient, err
	}

	cc := clientcredentials.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"manage_scope_grants"},
		TokenURL:     config.TokenURL,
	}

	acpClient.client = client.New(httptransport.NewWithClient(
		u.Host,
		"/",
		[]string{u.Scheme},
		cc.Client(context.WithValue(context.Background(), oauth2.HTTPClient, hc)),
	), nil)

	return acpClient, nil
}

func (a *AcpClient) GetScopeGrant(r LoginRequest) (*models.ScopeGrantSessionResponse, error) {
	var (
		response *logins.GetScopeGrantRequestOK
		err      error
	)

	if response, err = a.client.Logins.GetScopeGrantRequest(logins.NewGetScopeGrantRequestParams().
		WithTid(a.tenant).WithLoginID(r.ID), nil); err != nil {
		return nil, err
	}

	return response.Payload, nil
}

func (a *AcpClient) ApproveScopeGrant(r LoginRequest, approvedScopes []string) (string, error) {
	var (
		response *logins.AcceptScopeGrantRequestOK
		err      error
	)

	acceptScopeGrant := &models.AcceptScopeGrant{
		GrantedScopes: approvedScopes,
		ID:            r.ID,
		LoginState:    r.State,
	}

	if response, err = a.client.Logins.AcceptScopeGrantRequest(logins.NewAcceptScopeGrantRequestParams().
		WithTid(a.tenant).WithLoginID(r.ID).WithAcceptScopeGrant(acceptScopeGrant), nil); err != nil {
		return "", err
	}

	return response.Payload.RedirectTo, nil
}

func (a *AcpClient) RejectScopeGrant(r LoginRequest, rejectReason string, rejectStatusCode int64) (string, error) {
	var (
		response *logins.RejectScopeGrantRequestOK
		err      error
	)

	rejectScopeGrant := &models.RejectScopeGrant{
		ID:         r.ID,
		LoginState: r.State,
		Error:      rejectReason,
		StatusCode: rejectStatusCode,
	}

	if response, err = a.client.Logins.RejectScopeGrantRequest(logins.NewRejectScopeGrantRequestParams().
		WithTid(a.tenant).WithLoginID(r.ID).WithRejectScopeGrant(rejectScopeGrant), nil); err != nil {

		return "", nil
	}

	return response.Payload.RedirectTo, nil
}
