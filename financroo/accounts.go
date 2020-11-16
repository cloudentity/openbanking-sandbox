package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"time"
)

type AccountsRes struct {
	Data struct {
		Account []struct {
			BankID               string    `json:"BankID"`
			AccountID            string    `json:"AccountId"`
			Status               string    `json:"Status"`
			StatusUpdateDateTime time.Time `json:"StatusUpdateDateTime"`
			Currency             string    `json:"Currency"`
			AccountType          string    `json:"AccountType"`
			AccountSubType       string    `json:"AccountSubType"`
			Nickname             string    `json:"Nickname"`
		} `json:"Account"`
	} `json:"Data"`
	Links struct {
		Self string `json:"Self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"TotalPages"`
	} `json:"Meta"`
}

// accounts

var accounts = AccountsRes{}

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jsonFile, err := os.Open("accounts.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var accounts AccountsRes
		json.Unmarshal(byteValue, &accounts)
		c.JSON(200, accounts)
	}
}

func (s *Server) UpdateAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		json.NewDecoder(c.Request.Body).Decode(&accounts)
		c.JSON(200, accounts)
	}
}
