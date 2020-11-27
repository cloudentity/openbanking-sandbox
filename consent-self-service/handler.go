package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "app.tmpl", gin.H{})
	}
}
