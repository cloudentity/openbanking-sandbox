package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/cloudentity/openbanking-sandbox/models"
	paymentModels "github.com/cloudentity/openbanking-sandbox/openbanking/paymentinitiation/models"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/client/openbanking"
	acpClient "github.com/cloudentity/acp-client-go/models"
)

var (
	StatusPayment     string = "AcceptedSettlementInProcess"
	DomesticPaymentID string = "58923-001"
)

func (s *Server) CreateDomesticPayment() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			obDomesticRequest     *paymentModels.OBWriteDomestic2
			consent               paymentModels.OBWriteDomesticConsentResponse5
			err                   error
		)

		if introspectionResponse, err = s.IntrospectToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "payments") {
			msg := "token has no payments scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if err = json.NewDecoder(c.Request.Body).Decode(&obDomesticRequest); err != nil {
			msg := "unable to decode domestic payments request object"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		if consent, err = s.Storage.GetConsent(introspectionResponse.Subject, *obDomesticRequest.Data.ConsentID); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBError1{
				Message: &msg,
			})
			return
		}

		if *consent.Data.Status != "Authorised" { //nolint
			msg := "domestic payment consent does not have status authorised" //nolint
			c.JSON(http.StatusUnprocessableEntity, models.OBError1{
				Message: &msg,
			})
			return
		}

		// request initiation and consent initiation must match
		if !initiationsAreEqual(*obDomesticRequest.Data.Initiation, *consent.Data.Initiation) {
			msg := "request initiation does not match consent initiation"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		// request risk and consent risk must match
		if !risksAreEqual(*obDomesticRequest.Risk, *consent.Risk) {
			msg := "request risk does not match consent risk"
			c.JSON(http.StatusBadRequest, models.OBError1{
				Message: &msg,
			})
			return
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/domestic-payments", strconv.Itoa(s.Config.Port)))
		response := paymentModels.OBWriteDomesticResponse5{
			Data: &paymentModels.OBWriteDomesticResponse5Data{
				DomesticPaymentID:    &DomesticPaymentID,
				ConsentID:            consent.Data.ConsentID,
				Status:               &StatusPayment,
				Charges:              []*paymentModels.OBWriteDomesticResponse5DataChargesItems0{},
				CreationDateTime:     newDateTimePtr(time.Now()),
				StatusUpdateDateTime: newDateTimePtr(time.Now()),
				Initiation:           toDomesticResponse5DataInitiation(*consent.Data.Initiation),
			},
			Links: &paymentModels.Links{
				Self: &self,
			},
		}

		c.PureJSON(http.StatusOK, response)
	}
}

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			userAccounts          []models.OBAccount6
			err                   error
		)

		if introspectionResponse, err = s.IntrospectToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadAccountsBasic") {
			msg := "ReadAccountsBasic permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if userAccounts, err = s.Storage.GetAccounts(introspectionResponse.Subject); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		accounts := []*models.OBAccount6{}

		for _, a := range userAccounts {
			if has(introspectionResponse.AccountIDs, string(*a.AccountID)) {
				account := a
				if !has(grantedPermissions, "ReadAccountsDetail") {
					account.Account = []*models.OBAccount6AccountItems0{}
				}

				accounts = append(accounts, &account)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port)))
		response := models.OBReadAccount6{
			Data: &models.OBReadAccount6Data{
				Account: accounts,
			},
			Meta: &models.Meta{
				TotalPages: int32(len(accounts)),
			},
			Links: &models.Links{
				Self: &self,
			},
		}

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

		if accounts, err = s.Storage.GetAccounts(sub); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		ia := make([]InternalAccount, len(accounts))

		for i, a := range accounts {
			ia[i] = InternalAccount{
				ID:   *a.AccountID,
				Name: a.Nickname,
			}
		}

		c.PureJSON(http.StatusOK, InternalAccounts{Accounts: ia})
	}
}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			userBalances          []models.OBReadBalance1DataBalanceItems0
			err                   error
		)

		if introspectionResponse, err = s.IntrospectToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadBalances") {
			msg := "ReadBalances permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if userBalances, err = s.Storage.GetBalances(introspectionResponse.Subject); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		balances := []*models.OBReadBalance1DataBalanceItems0{}

		for _, balance := range userBalances {
			b := balance
			if has(introspectionResponse.AccountIDs, string(*b.AccountID)) {
				balances = append(balances, &b)
			}
		}

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/balances", strconv.Itoa(s.Config.Port)))
		response := models.OBReadBalance1{
			Data: &models.OBReadBalance1Data{
				Balance: balances,
			},
			Meta: &models.Meta{
				TotalPages: int32(len(balances)),
			},
			Links: &models.Links{
				Self: &self,
			},
		}

		c.PureJSON(http.StatusOK, response)
	}
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			introspectionResponse *acpClient.IntrospectOpenbankingAccountAccessConsentResponse
			userTransactions      []models.OBTransaction6
			err                   error
		)

		if introspectionResponse, err = s.IntrospectToken(c); err != nil {
			msg := fmt.Sprintf("failed to introspect token: %+v", err)
			c.JSON(http.StatusBadRequest, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		grantedPermissions := introspectionResponse.Permissions

		scopes := strings.Split(introspectionResponse.Scope, " ")
		if !has(scopes, "accounts") {
			msg := "token has no accounts scope granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if !has(grantedPermissions, "ReadTransactionsBasic") {
			msg := "ReadTransactionsBasic permission has not been granted"
			c.JSON(http.StatusForbidden, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		if userTransactions, err = s.Storage.GetTransactions(introspectionResponse.Subject); err != nil {
			msg := err.Error()
			c.JSON(http.StatusNotFound, models.OBErrorResponse1{
				Message: &msg,
			})
			return
		}

		transactions := []*models.OBTransaction6{}

		for _, transaction := range userTransactions {
			t := transaction
			if has(introspectionResponse.AccountIDs, string(*t.AccountID)) {
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

		self := strfmt.URI(fmt.Sprintf("http://localhost:%s/transactions", strconv.Itoa(s.Config.Port)))

		response := models.OBReadTransaction6{
			Data: &models.OBReadDataTransaction6{
				Transaction: transactions,
			},
			Meta: &models.Meta{
				TotalPages: int32(len(transactions)),
			},
			Links: &models.Links{
				Self: &self,
			},
		}

		c.PureJSON(http.StatusOK, response)
	}
}

func (s *Server) IntrospectToken(c *gin.Context) (*acpClient.IntrospectOpenbankingAccountAccessConsentResponse, error) {
	var (
		introspectionResponse *openbanking.OpenbankingAccountAccessConsentIntrospectOK
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
		return nil, err
	}

	return introspectionResponse.Payload, nil
}

func has(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func newDateTimePtr(t time.Time) *strfmt.DateTime {
	str := strfmt.DateTime(t)
	return &str
}

func initiationsAreEqual(initiation1, initiation2 interface{}) bool {
	var (
		initiation1Bytes []byte
		initiation2Bytes []byte
		err              error
	)
	if initiation1Bytes, err = json.Marshal(initiation1); err != nil {
		return false
	}
	if initiation2Bytes, err = json.Marshal(initiation2); err != nil {
		return false
	}
	return bytes.Equal(initiation1Bytes, initiation2Bytes)
}

func risksAreEqual(risk1, risk2 paymentModels.OBRisk1) bool {
	return reflect.DeepEqual(risk1, risk2)
}

func toDomesticResponse5DataInitiation(initiation paymentModels.OBWriteDomesticConsentResponse5DataInitiation) *paymentModels.OBWriteDomesticResponse5DataInitiation {
	return &paymentModels.OBWriteDomesticResponse5DataInitiation{
		CreditorAccount:           (*paymentModels.OBWriteDomesticResponse5DataInitiationCreditorAccount)(initiation.CreditorAccount),
		CreditorPostalAddress:     initiation.CreditorPostalAddress,
		DebtorAccount:             (*paymentModels.OBWriteDomesticResponse5DataInitiationDebtorAccount)(initiation.DebtorAccount),
		EndToEndIdentification:    initiation.EndToEndIdentification,
		InstructedAmount:          (*paymentModels.OBWriteDomesticResponse5DataInitiationInstructedAmount)(initiation.InstructedAmount),
		InstructionIdentification: initiation.InstructionIdentification,
		LocalInstrument:           initiation.LocalInstrument,
		RemittanceInformation:     (*paymentModels.OBWriteDomesticResponse5DataInitiationRemittanceInformation)(initiation.RemittanceInformation),
		SupplementaryData:         initiation.SupplementaryData,
	}
}
