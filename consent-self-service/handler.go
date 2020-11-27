package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			accounts           InternalAccounts
			consentsByAccounts *models.ListConsentsByAccountsResponse
			err                error
		)

		subject := "subject" // todo

		if accounts, err = s.BankClient.GetInternalAccounts(subject); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get accounts from bank: %+v", err))
			return
		}
		logrus.Infof("accounts: %+v", accounts)

		accountIDs := make([]string, len(accounts.Accounts))
		for i, a := range accounts.Accounts {
			accountIDs[i] = a.ID
		}

		if consentsByAccounts, err = s.Client.GetAccountAccessConsent(accountIDs); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get account access consent from acp: %+v", err))
			return
		}
		logrus.Infof("consentsByAccounts: %+v", consentsByAccounts)

		c.HTML(http.StatusOK, "app.tmpl", gin.H{})
	}
}
