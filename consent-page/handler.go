package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/models"

	"github.com/cloudentity/acp-client-go/client/openbanking"
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

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			loginRequest = NewLoginRequest(c)
			response     *openbanking.GetAccountAccessConsentSystemOK

			accounts InternalAccounts
			err      error
		)

		if err = loginRequest.Validate(); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if response, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
			openbanking.NewGetAccountAccessConsentSystemParams().
				WithTid(s.Client.TenantID).
				WithLoginID(loginRequest.ID),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get account access consent: %+v", err))
			return
		}

		if accounts, err = s.BankClient.GetInternalAccounts(response.Payload.Subject); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get accounts from bank: %+v", err))
			return
		}

		// TODO accounts ids should be masked
		c.HTML(http.StatusOK, "consent.tmpl", gin.H{
			"login_request": loginRequest,
			"subject":       response.Payload.Subject,
			"accounts":      accounts.Accounts,
			"permissions":   response.Payload.Permissions,
		})
	}
}

func (s *Server) Post() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			loginRequest = NewLoginRequest(c)
			redirect     string
			err          error
		)

		if err = loginRequest.Validate(); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		switch c.PostForm("action") {
		case "confirm":
			var (
				consent *openbanking.GetAccountAccessConsentSystemOK
				accept  *openbanking.AcceptAccountAccessConsentSystemOK
			)

			if consent, err = s.Client.Openbanking.GetAccountAccessConsentSystem(
				openbanking.NewGetAccountAccessConsentSystemParams().
					WithTid(s.Client.TenantID).
					WithLoginID(loginRequest.ID),
				nil,
			); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to get account access consent: %+v", err))
				return
			}

			// it's the consent page responsibility to decide which scopes should be granted
			grantScopes := make([]string, len(consent.Payload.RequestedScopes))
			for i, r := range consent.Payload.RequestedScopes {
				grantScopes[i] = r.RequestedName
			}

			if accept, err = s.Client.Openbanking.AcceptAccountAccessConsentSystem(
				openbanking.NewAcceptAccountAccessConsentSystemParams().
					WithTid(s.Client.TenantID).
					WithLoginID(loginRequest.ID).
					WithAcceptAccountAccessConsent(&models.AcceptAccountAccessConsentRequest{
						GrantedScopes: grantScopes,
						AccountIDs:    c.PostFormArray("account_ids"),
						LoginState:    loginRequest.State,
					}),
				nil,
			); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to accept account access consent: %+v", err))
				return
			}

			redirect = accept.Payload.RedirectTo

			logrus.Infof("account access consent accepted, redirect to: %s", redirect)
		case "deny":
			var reject *openbanking.RejectAccountAccessConsentSystemOK

			if reject, err = s.Client.Openbanking.RejectAccountAccessConsentSystem(
				openbanking.NewRejectAccountAccessConsentSystemParams().
					WithTid(s.Client.TenantID).
					WithLoginID(loginRequest.ID).
					WithRejectAccountAccessConsent(&models.RejectAccountAccessConsentRequest{
						ID:         loginRequest.ID,
						LoginState: loginRequest.State,
						Error:      "rejected",
						StatusCode: 403,
					}),
				nil,
			); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to reject account access consent: %+v", err))
				return
			}

			redirect = reject.Payload.RedirectTo

			logrus.Infof("account access consent denied, redirect to: %s", redirect)
		default:
			c.String(http.StatusBadRequest, "invalid action")
			return
		}

		c.Redirect(http.StatusFound, redirect)
	}
}
