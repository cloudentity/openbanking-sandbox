package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"time"
)

type TransactionsRes struct {
	Data struct {
		Transaction []struct {
			BankID               string `json:"BankId"`
			AccountID            string `json:"AccountId"`
			TransactionID        string `json:"TransactionId"`
			TransactionReference string `json:"TransactionReference"`
			Amount               struct {
				Amount   string `json:"Amount"`
				Currency string `json:"Currency"`
			} `json:"Amount"`
			CreditDebitIndicator   string    `json:"CreditDebitIndicator"`
			Status                 string    `json:"Status"`
			BookingDateTime        time.Time `json:"BookingDateTime"`
			ValueDateTime          time.Time `json:"ValueDateTime"`
			TransactionInformation string    `json:"TransactionInformation"`
			BankTransactionCode    struct {
				Code    string `json:"Code"`
				SubCode string `json:"SubCode"`
			} `json:"BankTransactionCode"`
			ProprietaryBankTransactionCode struct {
				Code   string `json:"Code"`
				Issuer string `json:"Issuer"`
			} `json:"ProprietaryBankTransactionCode"`
			Balance struct {
				Amount struct {
					Amount   string `json:"Amount"`
					Currency string `json:"Currency"`
				} `json:"Amount"`
				CreditDebitIndicator string `json:"CreditDebitIndicator"`
				Type                 string `json:"Type"`
			} `json:"Balance"`
			AddressLine string `json:"AddressLine,omitempty"`
		} `json:"Transaction"`
	} `json:"Data"`
	Links struct {
		Self string `json:"Self"`
	} `json:"Links"`
	Meta struct {
		TotalPages             int       `json:"TotalPages"`
		FirstAvailableDateTime time.Time `json:"FirstAvailableDateTime"`
		LastAvailableDateTime  time.Time `json:"LastAvailableDateTime"`
	} `json:"Meta"`
}

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jsonFile, err := os.Open("transactions.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var transactions TransactionsRes
		json.Unmarshal(byteValue, &transactions)
		c.JSON(200, transactions)
	}
}
