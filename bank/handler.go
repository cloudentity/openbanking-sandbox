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

var mockAccounts = []Account{
	{
		AccountID:      "22289",
		Status:         "Enabled",
		Currency:       "GBP",
		Nickname:       "Bills",
		AccountType:    "Personal",
		AccountSubType: "CurrentAccount",
		Account: []AccountDetails{
			{
				SchemeName:              "UK.OBIE.SortCodeAccountNumber",
				Identification:          "80200110203345",
				Name:                    "Mr Kevin",
				SecondaryIdentification: "00021",
			},
		},
	},
	{
		AccountID:      "31820",
		Status:         "Enabled",
		Currency:       "GBP",
		Nickname:       "Household",
		AccountType:    "Personal",
		AccountSubType: "CurrentAccount",
		Account: []AccountDetails{
			{
				SchemeName:     "UK.OBIE.SortCodeAccountNumber",
				Identification: "80200110203348",
				Name:           "Mr Kevin",
			},
		},
	},
}

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *models.IntrospectOpenbankingAccountAccessConsentResponse
			err                   error
		)

		token := c.GetHeader("Authorization")
		token = strings.ReplaceAll(token, "Bearer ", "")

		if introspectionResponse, err = s.AcpClient.Introspect(token); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to introspect token: %+v", err))
			return
		}

		logrus.WithFields(logrus.Fields{"introspection_response": introspectionResponse}).Infof("token introspected")

		authorizedAccounts := introspectionResponse.AccountIDs
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

		accounts := []Account{}
		for _, a := range mockAccounts {
			if has(authorizedAccounts, a.AccountID) {
				if !has(grantedPermissions, "ReadAccountsDetail") {
					a.Account = []AccountDetails{}
				}

				accounts = append(accounts, a)
			}
		}

		response := GetAccounts{Data: Data{Account: accounts}}
		response.Meta.TotalPages = len(response.Data.Account)
		response.Links.Self = fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port))

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
		// todo return accounts based on sub from path param
		accounts := make([]InternalAccount, len(mockAccounts))

		for i, a := range mockAccounts {
			accounts[i] = InternalAccount{
				ID:   a.AccountID,
				Name: a.Nickname,
			}
		}

		c.PureJSON(http.StatusOK, InternalAccounts{Accounts: accounts})
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
