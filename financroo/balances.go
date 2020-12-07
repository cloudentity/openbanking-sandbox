package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp/pkg/openbanking/client/balances"
)

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp  *balances.GetBalancesOK
			token string
			err   error
		)

		if token, err = c.Cookie("token"); err != nil {
			c.String(http.StatusUnauthorized, "failed to get token from cookie: %w", err)
			return
		}

		if resp, err = s.BankClient.Balances.GetBalances(balances.NewGetBalancesParams().WithAuthorization(token), nil); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get balances: %+v", err))
			return
		}

		c.JSON(200, gin.H{
			"balances": resp.Payload.Data.Balance,
		})
	}
}
