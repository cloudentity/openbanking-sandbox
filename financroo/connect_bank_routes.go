package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/cloudentity/openbanking-sandbox/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	challengeLength = 43
	nonceLength     = 20
)

type AppStorage struct {
	Verifier string
	Nonce    string
	Sub      string
	IntentID string
	BankID   BankID
}

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

func (s *Server) WebConfig() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authorizationServerURL": s.Config.ACPURL,
			"clientId":               s.Config.Login.ClientID,
			"authorizationServerId":  s.Config.Login.ServerID,
			"tenantId":               s.Config.Login.TenantID,
		})
	}
}

type ConnectBankRequest struct {
	Permissions []string `json:"permissions" binding:"required"`
}

func (s *Server) ConnectBank() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID             = BankID(c.Param("bankId"))
			clients            Clients
			ok                 bool
			registerResponse   *models.OBReadConsentResponse1
			encodedVerifier    string
			encodedNonce       string
			encodedCookieValue string
			challenge          string
			loginURL           string
			data               = gin.H{}
			connectRequest     = ConnectBankRequest{}
			user               User
			err                error
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if err = c.BindJSON(&connectRequest); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if clients, ok = s.Clients[bankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", bankID))
		}

		if registerResponse, err = clients.AcpSystemClient.RegisterAccountAccessConsent(connectRequest.Permissions); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register account access consent: %+v", err))
			return
		}

		// generate verifier
		verifier := make([]byte, challengeLength)
		if _, err = io.ReadFull(rand.Reader, verifier); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to generate challenge: %+v", err))
			return
		}
		encodedVerifier = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(verifier)

		// generate nonce
		nonce := make([]byte, nonceLength)
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to generate nonce: %+v", err))
			return
		}
		encodedNonce = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(nonce)

		app := AppStorage{
			Verifier: encodedVerifier,
			Nonce:    encodedNonce,
			BankID:   bankID,
			IntentID: *registerResponse.Data.ConsentID,
			Sub:      user.Sub,
		}

		// persist verifier and nonce in a secure encrypted cookie
		if encodedCookieValue, err = s.SecureCookie.Encode("app", app); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error while encoding cookie: %+v", err))
			return
		}

		c.SetCookie("app", encodedCookieValue, 0, "/", "", false, true)

		hash := sha256.New()
		if _, err = hash.Write([]byte(encodedVerifier)); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error while creating challenge: %+v", err))
			return
		}
		challenge = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash.Sum([]byte{}))

		if loginURL, err = clients.AcpWebClient.LoginURL(*registerResponse.Data.ConsentID, challenge, []string{"openid", "accounts", "offline_access"}, encodedNonce); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
		}

		data["login_url"] = loginURL

		c.JSON(http.StatusOK, data)
	}
}

func (s *Server) ConnectBankCallback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app        string
			appStorage = AppStorage{}
			code       = c.Query("code")
			ok         bool
			clients    Clients
			token      Token
			err        error
		)

		if c.Query("error") != "" {
			c.String(http.StatusBadRequest, fmt.Sprintf("acp returned an error: %s: %s", c.Query("error"), c.Query("error_description")))
			return
		}

		if app, err = c.Cookie("app"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get app cookie: %+v", err))
			return
		}

		if err = s.SecureCookie.Decode("app", app, &appStorage); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode app storage: %+v", err))
			return
		}

		if clients, ok = s.Clients[appStorage.BankID]; !ok {
			c.String(http.StatusBadRequest, fmt.Sprintf("client not configured for bank: %s", appStorage.BankID))
		}

		if token, err = clients.AcpWebClient.Exchange(code, appStorage.Verifier); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		if err = s.ConnectBankForUser(appStorage, token); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))

			return
		}

		c.SetCookie("app", "", -1, "/", "", false, true)

		c.Redirect(http.StatusFound, s.Config.Login.UIURL)
	}
}

func (s *Server) ConnectedBanks() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			user           User
			err            error
			connectedBanks = []string{}
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		for _, b := range user.Banks {
			connectedBanks = append(connectedBanks, b.BankID)
		}

		c.JSON(200, gin.H{
			"connected_banks": connectedBanks,
		})
	}
}

func (s *Server) DisconnectBank() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			bankID = c.Param("bankId")
			user   User
			err    error
		)

		if user, err = s.WithUser(c); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to get user: %+v", err))
			return
		}

		cb := []ConnectedBank{}
		for _, b := range user.Banks {
			if b.BankID != bankID {
				cb = append(cb, b)
			}
		}
		user.Banks = cb

		if err = s.UserRepo.Set(user); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %+v", err))
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *Server) ConnectBankForUser(appStorage AppStorage, token Token) error {
	var (
		user User
		err  error
		cb   = ConnectedBank{
			BankID:       string(appStorage.BankID),
			IntentID:     appStorage.IntentID,
			RefreshToken: token.RefreshToken,
		}
		found = false
	)

	if user, err = s.UserRepo.Get(appStorage.Sub); err != nil {
		return errors.Wrapf(err, "failed to get user")
	}

	for i, b := range user.Banks {
		if b.BankID == string(appStorage.BankID) {
			user.Banks[i] = cb
			found = true
			break
		}
	}

	if !found {
		user.Banks = append(user.Banks, cb)
	}

	if err = s.UserRepo.Set(user); err != nil {
		return errors.Wrapf(err, "failed to update user")
	}

	return nil
}
