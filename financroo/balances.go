package main

import (
	"fmt"
	"github.com/cloudentity/acp/pkg/openbanking/client/balances"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp *balances.GetBalancesOK
			err  error
		)

		token, _ := c.Cookie("token")
		if resp, err = s.BankClient.Balances.GetBalances(balances.NewGetBalancesParams().WithAuthorization(token), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get balances: %+v", err))
			return
		}

		c.JSON(200, gin.H{
			"balances": resp.Payload.Data.Balance,
		})
	}
}
