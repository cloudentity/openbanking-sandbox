package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudentity/openbanking-sandbox/client/accounts"
	"github.com/cloudentity/openbanking-sandbox/client/balances"
	"github.com/cloudentity/openbanking-sandbox/client/transactions"
	"github.com/cloudentity/openbanking-sandbox/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func (s *Server) WithUser(c *gin.Context) (User, error) {
	var (
		user   User
		claims map[string]interface{}
		sub    string
		ok     bool
		err    error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if token == "" {
		return user, errors.New("no access token provided")
	}

	if claims, err = s.LoginClient.Userinfo(token); err != nil {
		return user, errors.Wrapf(err, "invalid token")
	}

	if sub, ok = claims["sub"].(string); !ok {
		return user, errors.New("sub claim is missing")
	}

	if user, err = s.UserRepo.Get(sub); err != nil {
		return user, errors.Wrapf(err, "failed to get user")
	}

	return user, nil
}

func (s *Server) GetClientWithToken(bank ConnectedBank) (OpenbankingClient, Token, error) {
	var (
		clients Clients
		client  OpenbankingClient
		ok      bool
		token   Token
		err     error
	)

	if clients, ok = s.Clients[BankID(bank.BankID)]; !ok {
		return client, token, fmt.Errorf("can't get client for a bank: %s", bank.BankID)
	}

	logrus.Infof("XXX renew token: %s", bank.RefreshToken)
	if token, err = clients.AcpWebClient.RenewToken(bank.RefreshToken); err != nil {
		return client, token, errors.Wrapf(err, "can't renew access token for a bank: %s", bank.BankID)
	}
	logrus.Infof("XXX new token: %s", token.AccessToken)

	return clients.BankClient, token, nil
}

type Account struct {
	*models.OBAccount6
	BankID string `json:"BankId"`
}

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp         *accounts.GetAccountsOK
			client       OpenbankingClient
			token        Token
			accountsData = []Account{}
			user         User
			err          error
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
		}

		// todo parallel
		for i, b := range user.Banks {
			if client, token, err = s.GetClientWithToken(b); err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				return
			}
			if err = s.UserRepo.Set(user); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			}

			if resp, err = client.Accounts.GetAccounts(accounts.NewGetAccountsParams().WithAuthorization(token.AccessToken), nil); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get accounts: %+v", err))
				return
			}

			user.Banks[i].RefreshToken = token.RefreshToken

			for _, a := range resp.Payload.Data.Account {
				accountsData = append(accountsData, Account{
					OBAccount6: a,
					BankID:     b.BankID,
				})
			}
		}

		c.JSON(200, gin.H{
			"accounts": accountsData,
		})
	}
}

type Balance struct {
	models.OBReadBalance1DataBalanceItems0
	BankID string `json:"BankId"`
}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp         *balances.GetBalancesOK
			client       OpenbankingClient
			token        Token
			balancesData = []Balance{}
			user         User
			err          error
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		// todo parallel
		for i, b := range user.Banks {
			if client, token, err = s.GetClientWithToken(b); err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				return
			}
			if err = s.UserRepo.Set(user); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			}

			if resp, err = client.Balances.GetBalances(balances.NewGetBalancesParams().WithAuthorization(token.AccessToken), nil); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get balances: %+v", err))
				return
			}

			user.Banks[i].RefreshToken = token.RefreshToken

			for _, a := range resp.Payload.Data.Balance {
				balancesData = append(balancesData, Balance{
					OBReadBalance1DataBalanceItems0: *a,
					BankID:                          b.BankID,
				})
			}
		}

		c.JSON(200, gin.H{
			"balances": balancesData,
		})
	}
}

type Transaction struct {
	models.OBTransaction6
	BankID string `json:"BankId"`
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			resp             *transactions.GetTransactionsOK
			client           OpenbankingClient
			token            Token
			transactionsData = []Transaction{}
			user             User
			err              error
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
		}

		// todo parallel
		for i, b := range user.Banks {
			if client, token, err = s.GetClientWithToken(b); err != nil {
				c.String(http.StatusUnauthorized, err.Error())
				return
			}
			if err = s.UserRepo.Set(user); err != nil {
				c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			}

			if resp, err = client.Transactions.GetTransactions(transactions.NewGetTransactionsParams().WithAuthorization(token.AccessToken), nil); err != nil {
				c.String(http.StatusUnauthorized, fmt.Sprintf("failed to call bank get transactions: %+v", err))
				return
			}

			user.Banks[i].RefreshToken = token.RefreshToken

			for _, a := range resp.Payload.Data.Transaction {
				transactionsData = append(transactionsData, Transaction{
					OBTransaction6: *a,
					BankID:         b.BankID,
				})
			}
		}

		c.JSON(200, gin.H{
			"transactions": transactionsData,
		})
	}
}
