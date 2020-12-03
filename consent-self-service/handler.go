package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type Client struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	LogoURI   string `json:"logo_uri"`
	ClientURI string `json:"client_uri"`
}

type ClientConsents struct {
	Client
	Consents []models.OpenbankingAccountAccessConsentWithClient `json:"consents"`
}

type ConsentsResponse struct {
	ClientConsents []ClientConsents `json:"client_consents"`
}

func (s *Server) ListConsents() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			consentsByAccounts *models.ListAccountAccessConsentsWithClient
			clientToConsents   = map[string][]models.OpenbankingAccountAccessConsentWithClient{}
			clients            = map[string]Client{}
			res                = []ClientConsents{}
			err                error
		)

		if consentsByAccounts, err = s.FetchAccounts(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		for _, c := range consentsByAccounts.Consents {
			if _, ok := clients[c.ID]; !ok {
				clients[c.ID] = Client{
					ID:        c.ID,
					Name:      c.Name,
					LogoURI:   c.LogoURI,
					ClientURI: c.ClientURI,
				}
			}

			clientToConsents[c.ID] = append(clientToConsents[c.ID], *c)
		}

		for _, x := range clients {
			res = append(res, ClientConsents{
				Client:   x,
				Consents: clientToConsents[x.ID],
			})
		}

		c.JSON(http.StatusOK, &ConsentsResponse{ClientConsents: res})
	}
}

func (s *Server) FetchAccounts(c *gin.Context) (*models.ListAccountAccessConsentsWithClient, error) {
	var (
		accounts           InternalAccounts
		consentsByAccounts *models.ListAccountAccessConsentsWithClient
		at                 *models.IntrospectResponse
		err                error
	)

	token := c.GetHeader("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")

	if at, err = s.IntrospectClient.IntrospectToken(token); err != nil {
		return nil, fmt.Errorf("failed to introspect client: %w", err)
	}

	if accounts, err = s.BankClient.GetInternalAccounts(at.Subject); err != nil {
		return nil, fmt.Errorf("failed to get accounts from bank: %w", err)
	}

	accountIDs := make([]string, len(accounts.Accounts))
	for i, a := range accounts.Accounts {
		accountIDs[i] = a.ID
	}

	if consentsByAccounts, err = s.Client.GetAccountAccessConsent(accountIDs); err != nil {
		return nil, fmt.Errorf("failed to get account access consent from acp: %wv", err)
	}

	return consentsByAccounts, nil
}

func (s *Server) RevokeConsent() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			id                 = c.Param("id")
			consentsByAccounts *models.ListAccountAccessConsentsWithClient
			canBeRevoked       bool
			err                error
		)

		if consentsByAccounts, err = s.FetchAccounts(c); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		for _, c := range consentsByAccounts.Consents {
			if c.ConsentID == id {
				canBeRevoked = true
				break
			}
		}

		if !canBeRevoked {
			c.String(http.StatusBadRequest, "user is not authorized to revoke this consent")
		}

		if err = s.Client.RevokeAccountAccessConsent(id); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to revoke account access consent: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}
