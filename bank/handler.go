package main

import (
	"fmt"
	"github.com/cloudentity/openbanking-sandbox/models"
	"github.com/go-openapi/strfmt"
	"net/http"
	"strconv"
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
				account := a
				if !has(grantedPermissions, "ReadAccountsDetail") {
					account.Account = []*models.OBAccount6AccountItems0{}
				}

				accounts = append(accounts, &account)
			}
		}

		response := models.OBReadAccount6{Data: &models.OBReadAccount6Data{Account: accounts}, Meta: &models.Meta{}, Links: &models.Links{}}
		response.Meta.TotalPages = int32(len(accounts))

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port)))
		response.Links.Self = &self

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
		var (
			introspectionResponse *openbanking.OpenbankingAccountAccessConsentIntrospectOK
			userBalances          []models.OBReadBalance1DataBalanceItems0
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

		if !has(grantedPermissions, "ReadBalances") {
			c.String(http.StatusForbidden, "ReadBalances permission has not been granted")
			return
		}

		if userBalances, err = s.Storage.GetBalance(introspectionResponse.Payload.Subject); err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		balances := []*models.OBReadBalance1DataBalanceItems0{}

		for _, balance := range userBalances {
			b := balance
			if has(introspectionResponse.Payload.AccountIDs, string(b.AccountID)) {
				balances = append(balances, &b)
			}
		}

		response := models.OBReadBalance1{Data: &models.OBReadBalance1Data{Balance: balances}, Meta: &models.Meta{}, Links: &models.Links{}}

		response.Meta.TotalPages = int32(len(balances))

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(s.Config.Port)))
		response.Links.Self = &self

		logrus.WithFields(logrus.Fields{"response": response}).Infof("balances response")

		c.PureJSON(http.StatusOK, response)
	}
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *openbanking.OpenbankingAccountAccessConsentIntrospectOK
			userTransactions      []models.OBTransaction6
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

		if !has(grantedPermissions, "ReadTransactionsBasic") {
			c.String(http.StatusForbidden, "ReadTransactionsBasic permission has not been granted")
			return
		}

		if userTransactions, err = s.Storage.GetTransaction(introspectionResponse.Payload.Subject); err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}

		transactions := []*models.OBTransaction6{}

		for _, transaction := range userTransactions {
			t := transaction
			if has(introspectionResponse.Payload.AccountIDs, string(t.AccountID)) {
				if !has(grantedPermissions, "ReadTransactionsDetail") {
					t.TransactionInformation = ""
					t.Balance = &models.OBTransactionCashBalance{}
					t.MerchantDetails = &models.OBMerchantDetails1{}
					t.CreditorAgent = &models.OBBranchAndFinancialInstitutionIdentification61{}
					t.CreditorAccount = &models.OBCashAccount60{}
					t.DebtorAgent = &models.OBBranchAndFinancialInstitutionIdentification62{}
					t.DebtorAccount = &models.OBCashAccount61{}
				}

				transactions = append(transactions, &t)
			}
		}

		response := models.OBReadTransaction6{Data: &models.OBReadDataTransaction6{Transaction: transactions}, Meta: &models.Meta{}, Links: &models.Links{}}

		response.Meta.TotalPages = int32(len(transactions))

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/transactions", strconv.Itoa(s.Config.Port)))
		response.Links.Self = &self

		logrus.WithFields(logrus.Fields{"response": response}).Infof("transactions response")

		c.PureJSON(http.StatusOK, response)
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
