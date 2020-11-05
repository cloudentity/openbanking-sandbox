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
			loginRequest = NewLoginRequest(c)
			response     *models.GetAccountAccessConsentResponse
			err          error
		)

		if err = loginRequest.Validate(); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if response, err = s.Client.GetAccountAccessConsent(loginRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get account access consent: %+v", err))
			return
		}

		// at this point, we know that user is authenticated and bank should fetch list of accounts for this user
		accountIDs := []string{
			"27 1140 2004 0000 3002 0135 5387",
			"17 2240 1401 0000 000 0155 1312",
		}

		c.HTML(http.StatusOK, "consent.tmpl", gin.H{
			"login_request": loginRequest,
			"subject":       response.Subject,
			"account_ids":   accountIDs,
			"permissions":   response.Permissions,
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
			if redirect, err = s.Client.ApproveAccountAccessConsent(loginRequest, c.PostFormArray("account_ids")); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to accept account access consent: %+v", err))
				return
			}

			logrus.Infof("account access consent accepted, redirect to: %s", redirect)
		case "deny":
			if redirect, err = s.Client.RejectAccountAccessConsent(loginRequest, "rejected", 403); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to reject account access consent: %+v", err))
				return
			}

			logrus.Infof("account access consent denied, redirect to: %s", redirect)
		default:
			c.String(http.StatusBadRequest, "invalid action")
			return
		}

		c.Redirect(http.StatusFound, redirect)
	}
}
