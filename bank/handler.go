package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *models.IntrospectOpenbankingAccountAccessConsentResponse
			userAccounts          []Account
			err                   error
		)

		token := c.GetHeader("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")

		if introspectionResponse, err = s.AcpClient.Introspect(token); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to introspect token: %+v", err))
			return
		}

		logrus.WithFields(logrus.Fields{"introspection_response": introspectionResponse}).Infof("token introspected")

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			c.String(http.StatusForbidden, "token has no accounts scope granted")
			return
		}

		if !has(grantedPermissions, "ReadAccountsBasic") {
			c.String(http.StatusForbidden, "ReadAccountsBasic permission has not been granted")
			return
		}

		if userAccounts, err = s.AccountsStorage.GetAccount(introspectionResponse.Subject); err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		accounts := []Account{}

		for _, a := range userAccounts {
			if has(introspectionResponse.AccountIDs, a.AccountID) {
				if !has(grantedPermissions, "ReadAccountsDetail") {
					a.Account = []AccountDetails{}
				}

				accounts = append(accounts, a)
			}
		}

		response := GetAccounts{Data: Data{Account: accounts}}
		response.Meta.TotalPages = len(response.Data.Account)
		response.Links.Self = fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port))

		logrus.WithFields(logrus.Fields{"response": response}).Infof("accounts response")

		c.PureJSON(http.StatusOK, response)
	}
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// this API is bank specific. It should return all users's account.
func (s *Server) InternalGetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			accounts []Account
			sub      = c.Param("sub")
			err      error
		)

		if accounts, err = s.AccountsStorage.GetAccount(sub); err != nil {
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

func has(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
