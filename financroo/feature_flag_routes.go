package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnableFeature(name string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println(name)
		c.SetCookie(name, "yes", 60*60*24, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	}
}

func DisableFeature(name string) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.SetCookie(name, "yes", -1, "/", "", false, false)
		c.Redirect(http.StatusFound, "/")
	}
}

func (s *Server) EnableInvestmentsFeature() func(ctx *gin.Context) {
	return EnableFeature("feature-flag-investments")
}

func (s *Server) DisableInvestmentsFeature() func(ctx *gin.Context) {
	return DisableFeature("feature-flag-investments")
}
