package main

import (
	"fmt"
	"net/http"

	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			loginRequest  = NewLoginRequest(c)
			grantResponse *models.ScopeGrantSessionResponse
			err           error
		)

		if err = loginRequest.Validate(); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if grantResponse, err = s.Client.GetScopeGrant(loginRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get scope grant: %+v", err))
			return
		}

		c.HTML(http.StatusOK, "consent.tmpl", gin.H{
			"login_request":    loginRequest,
			"client_id":        grantResponse.ClientID,
			"requested_scopes": grantResponse.RequestedScopes,
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
		case "Approve":
			if redirect, err = s.Client.ApproveScopeGrant(loginRequest, c.PostFormArray("scopes")); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to accept scope grant: %+v", err))
				return
			}

			logrus.Infof("scope grant accepted, redirect to: %s", redirect)
		case "Reject":
			if redirect, err = s.Client.RejectScopeGrant(loginRequest, "rejected", 403); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to reject scope grant: %+v", err))
				return
			}

			logrus.Infof("scope grant denied, redirect to: %s", redirect)
		default:
			c.String(http.StatusBadRequest, "invalid action")
			return
		}

		c.Redirect(http.StatusFound, redirect)
	}
}
