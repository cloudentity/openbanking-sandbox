package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"key": "value"})
	}
}
