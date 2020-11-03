package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
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
