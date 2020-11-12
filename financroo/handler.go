package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func (s *Server) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		// call ACP account-access-consent (cc with mtls) -> get intent_id
		// init authorization code grant flow with pkce, put intent_id in request param
		// c.Redirect(http.StatusFound, redirect)
	}
}

func (s *Server) Callback() func(*gin.Context) {
	return func(c *gin.Context) {
		// exchange code for a token
		c.HTML(http.StatusOK, "authenticated.tmpl", gin.H{})
	}
}

type Account struct {
	AccountId      string `json:"account_id"`
	Description    string `json:"description"`
	Identification string `json:"identification"`
	Name           string `json:"name"`
}

type AccountsRes struct {
	Accounts []Account `json:"accounts"`
}

// mock

var accounts = AccountsRes{
	Accounts: []Account{},
}

func (s *Server) GetAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, accounts)
	}
}

func (s *Server) UpdateAccounts() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		json.NewDecoder(c.Request.Body).Decode(&accounts)
		c.JSON(200, accounts)
	}
}
