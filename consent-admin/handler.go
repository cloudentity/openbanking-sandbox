package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cloudentity/acp/pkg/swagger/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type Client struct {
	ID       string                                    `json:"client_id"`
	Name     string                                    `json:"client_name,omitempty"`
	Consents []*models.OpenbankingAccountAccessConsent `json:"consents"`
}

type ListClientsResponse struct {
	Clients []Client `json:"clients"`
}

func (s *Server) ListClients() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			clients             *models.Clients
			consents            *models.ListAccountAccessConsents
			clientsWithConsents []Client
			err                 error
		)

		if err = s.IntrospectToken(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		if clients, err = s.Client.ListClients(); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to list clients from acp: %+v", err))
			return
		}

		for _, oc := range clients.Clients {
			// todo parallel
			if consents, err = s.Client.ListConsentsForClient(oc.ID); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("failed to list consents for client: %s, err: %+v", oc.ID, err))
				return
			}

			if !oc.System {
				clientsWithConsents = append(clientsWithConsents, Client{
					ID:       oc.ID,
					Name:     oc.Name,
					Consents: consents.Consents,
				})
			}
		}

		resp := ListClientsResponse{Clients: clientsWithConsents}

		c.JSON(http.StatusOK, &resp)
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
