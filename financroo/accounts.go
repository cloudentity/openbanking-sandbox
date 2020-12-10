package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/openbanking-sandbox/client/accounts"
	"github.com/openbanking-sandbox/models"
)

type BankAccount struct {
	*models.OBAccount6
	BankID string `json:"BankId"`
}

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			token        string
			resp         *accounts.GetAccountsOK
			accountsData []BankAccount
			err          error
		)

		if token, err = c.Cookie("token"); err != nil {
			c.String(http.StatusUnauthorized, "failed to get token from cookie: %w", err)
			return
		}

		if resp, err = s.BankClient.Accounts.GetAccounts(accounts.NewGetAccountsParams().WithAuthorization(token), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
			return
		}

		accountsData = make([]BankAccount, len(resp.Payload.Data.Account))
		for i, a := range resp.Payload.Data.Account {
			accountsData[i] = BankAccount{
				OBAccount6: a,
				BankID:     "gobank",
			}
		}

		c.JSON(200, gin.H{
			"accounts": accountsData,
		})
	}
}
