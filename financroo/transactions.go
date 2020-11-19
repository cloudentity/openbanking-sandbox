package main

import (
	"fmt"
	"github.com/cloudentity/acp/pkg/openbanking/client/transactions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)

		var resp *transactions.GetTransactionsOK

		token, _ := c.Cookie("token")
		if resp, err = s.BankClient.Transactions.GetTransactions(transactions.NewGetTransactionsParams().WithAuthorization(token), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get transactions: %+v", err))
			return
		}
		c.JSON(200, gin.H{
			"transactions": resp.Payload.Data.Transaction,
		})
	}
}
