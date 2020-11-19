package main

import (
	"fmt"
	"github.com/cloudentity/acp/pkg/openbanking/client/accounts"
	"github.com/cloudentity/acp/pkg/openbanking/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BankAccount struct {
	*models.OBAccount6
	BankId string
}

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			accountsData []BankAccount
			err          error
		)

		var resp *accounts.GetAccountsOK

		token, _ := c.Cookie("token")
		if resp, err = s.BankClient.Accounts.GetAccounts(accounts.NewGetAccountsParams().WithAuthorization(token), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
			return
		}

		accountsData = make([]BankAccount, len(resp.Payload.Data.Account))
		for i, a := range resp.Payload.Data.Account {
			accountsData[i] = BankAccount{
				OBAccount6: a,
				BankId:     "gobank",
			}
		}

		c.JSON(200, gin.H{
			"accounts": accountsData,
		})
	}
}
