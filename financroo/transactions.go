package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/openbanking-sandbox/client/transactions"
)

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			token string
			resp  *transactions.GetTransactionsOK
			err   error
		)

		if token, err = c.Cookie("token"); err != nil {
			c.String(http.StatusUnauthorized, "failed to get token from cookie: %w", err)
			return
		}

		if resp, err = s.BankClient.Transactions.GetTransactions(transactions.NewGetTransactionsParams().WithAuthorization(token), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get transactions: %+v", err))
			return
		}

		c.JSON(200, gin.H{
			"transactions": resp.Payload.Data.Transaction,
		})
	}
}
