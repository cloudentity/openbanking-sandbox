package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			accounts InternalAccounts
			err      error
		)

		subject := "subject" // todo

		if accounts, err = s.BankClient.GetInternalAccounts(subject); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get accounts from bank: %+v", err))
			return
		}

		logrus.Infof("accounts: %+v", accounts)
		// Fetch the consents from the new ACP API (listConsentsByAccounts)

		c.HTML(http.StatusOK, "app.tmpl", gin.H{})
	}
}
