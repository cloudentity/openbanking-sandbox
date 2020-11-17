package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"time"
)

type BalancesRes struct {
	Data struct {
		Balance []struct {
			AccountID string `json:"AccountId"`
			Amount    struct {
				Amount   string `json:"Amount"`
				Currency string `json:"Currency"`
			} `json:"Amount"`
			CreditDebitIndicator string    `json:"CreditDebitIndicator"`
			Type                 string    `json:"Type"`
			DateTime             time.Time `json:"DateTime"`
			CreditLine           []struct {
				Included bool `json:"Included"`
				Amount   struct {
					Amount   string `json:"Amount"`
					Currency string `json:"Currency"`
				} `json:"Amount"`
				Type string `json:"Type"`
			} `json:"CreditLine,omitempty"`
		} `json:"Balance"`
	} `json:"Data"`
	Links struct {
		Self string `json:"Self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"TotalPages"`
	} `json:"Meta"`
}

// balances

var balances = BalancesRes{}

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jsonFile, err := os.Open("balances.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var balances BalancesRes
		json.Unmarshal(byteValue, &balances)
		c.JSON(200, balances)
	}
}

func (s *Server) UpdateBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		json.NewDecoder(c.Request.Body).Decode(&balances)
		c.JSON(200, balances)
	}
}
