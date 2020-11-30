package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func (s *Server) ListConsents() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			consentsByAccounts *models.ListConsentsByAccountsResponse
			err                error
		)

		if consentsByAccounts, err = s.FetchAccounts(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, &consentsByAccounts)
	}
}

func (s *Server) FetchAccounts(c *gin.Context) (*models.ListConsentsByAccountsResponse, error) {
	var (
		accounts           InternalAccounts
		consentsByAccounts *models.ListConsentsByAccountsResponse
		err                error
	)

	subject := "subject" // todo introspect token

	if accounts, err = s.BankClient.GetInternalAccounts(subject); err != nil {
		return nil, fmt.Errorf("failed to get accounts from bank: %w", err)
	}

	accountIDs := make([]string, len(accounts.Accounts))
	for i, a := range accounts.Accounts {
		accountIDs[i] = a.ID
	}

	if consentsByAccounts, err = s.Client.GetAccountAccessConsent(accountIDs); err != nil {
		return nil, fmt.Errorf("failed to get account access consent from acp: %wv", err)
	}

	return consentsByAccounts, nil
}

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
