package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp/pkg/swagger/models"
	"strings"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func (s *Server) ListClients() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			clients *models.Clients
			err     error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if clients, err = s.Client.ListClients(); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to list clients from acp: %+v", err))
			return
		}

		c.JSON(http.StatusOK, &clients)
	}
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id  = c.Param("id")
			err error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if err = s.Client.RevokeConsent(id); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) RevokeConsentsForClient() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id  = c.Param("id")
			err error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if err = s.Client.RevokeConsentsForClient(id); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consents for client: %s, err: %+v", id, err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) IntrospectToken(c *gin.Context) error {
	var err error

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if _, err = s.IntrospectClient.IntrospectToken(token); err != nil {
		return fmt.Errorf("failed to introspect client: %w", err)
	}

	return nil
}
