package main

import (
	"fmt"
	"github.com/cloudentity/openbanking-sandbox/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/client/openbanking"
)

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *openbanking.OpenbankingAccountAccessConsentIntrospectOK
			userAccounts          []models.OBAccount6
			err                   error
		)

		token := c.GetHeader("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")

		if introspectionResponse, err = s.Client.Openbanking.OpenbankingAccountAccessConsentIntrospect(
			openbanking.NewOpenbankingAccountAccessConsentIntrospectParams().
				WithTid(s.Client.TenantID).
				WithAid(s.Client.ServerID).
				WithToken(&token),
			nil,
		); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to introspect token: %+v", err))
			return
		}

		logrus.WithFields(logrus.Fields{"introspection_response": introspectionResponse.Payload}).Infof("token introspected")

		grantedPermissions := introspectionResponse.Payload.Permissions

		scopes := strings.Split(introspectionResponse.Payload.Scope, " ")
		if !has(scopes, "accounts") {
			c.String(http.StatusForbidden, "token has no accounts scope granted")
			return
		}

		if !has(grantedPermissions, "ReadAccountsBasic") {
			c.String(http.StatusForbidden, "ReadAccountsBasic permission has not been granted")
			return
		}

		if userAccounts, err = s.Storage.GetAccount(introspectionResponse.Payload.Subject); err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		accounts := []*models.OBAccount6{}

		for _, a := range userAccounts {
			if has(introspectionResponse.Payload.AccountIDs, string(a.AccountID)) {
				if !has(grantedPermissions, "ReadAccountsDetail") {
					a.Account = []*models.OBAccount6AccountItems0{}
				}

				accounts = append(accounts, &a)
			}
		}

		response := models.OBReadAccount6{Data: &models.OBReadAccount6Data{Account: accounts}}
		//response.Meta.TotalPages = len(response.Data.Account)
		//response.Links.Self = fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port))

		logrus.WithFields(logrus.Fields{"response": response}).Infof("accounts response")

		c.PureJSON(http.StatusOK, response)
	}
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID   models.AccountID `json:"id"`
	Name models.Nickname  `json:"name"`
}

// this API is bank specific. It should return all users's account.
func (s *Server) InternalGetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			accounts []models.OBAccount6
			sub      = c.Param("sub")
			err      error
		)

		if accounts, err = s.Storage.GetAccount(sub); err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		ia := make([]InternalAccount, len(accounts))

		for i, a := range accounts {
			ia[i] = InternalAccount{
				ID:   a.AccountID,
				Name: a.Nickname,
			}
		}

		c.PureJSON(http.StatusOK, InternalAccounts{Accounts: ia})
	}
}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.File("./data/balances.json")
	}
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.File("./data/transactions.json")
	}
}

func has(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
