package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/sirupsen/logrus"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func (s *Server) ListClients() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			clients *models.Clients
			err     error
		)

		if clients, err = s.FetchClients(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, &clients)
	}
}

func (s *Server) FetchClients(c *gin.Context) (*models.Clients, error) {
	var (
		clients *models.Clients
		err     error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if _, err = s.IntrospectClient.IntrospectToken(token); err != nil {
		return nil, fmt.Errorf("failed to introspect client: %w", err)
	}

	if clients, err = s.Client.ListClients(); err != nil {
		return nil, fmt.Errorf("failed to get clients from acp: %wv", err)
	}
	logrus.Infof("XXX clients: %+v", clients)

	return clients, nil
}

/*
func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                 = c.Param("id")
			consentsByAccounts *models.ListConsentsByAccountsResponse
			canBeRevoked       bool
			err                error
		)

		if consentsByAccounts, err = s.FetchAccounts(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		for _, c := range consentsByAccounts.Consents {
			if c.ConsentID == id {
				canBeRevoked = true
				break
			}
		}

		if !canBeRevoked {
			c.String(http.StatusBadRequest, "user is not authorized to revoke this consent")
		}

		if err = s.Client.RevokeAccountAccessConsent(id); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}
*/
