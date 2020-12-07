package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/cloudentity/acp/pkg/openbanking/models"
)

const (
	challengeLength = 43
	nonceLength     = 20
)

type AppStorage struct {
	Verifier    string
	Nonce       string
	AccessToken string
}

func (s *Server) Index() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}

type PermissionsRes struct {
	Permissions []string `json:"permissions" binding:"required"`
}

func (s *Server) Connect() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			registerResponse   *models.OBReadConsentResponse1
			encodedVerifier    string
			encodedNonce       string
			encodedCookieValue string
			challenge          string
			loginURL           string
			data               = gin.H{}
			permissionsRes     = PermissionsRes{}
			err                error
		)

		if err = c.BindJSON(&permissionsRes); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to parse request body: %+v", err))
			return
		}

		if registerResponse, err = s.AccountAccessClient.RegisterAccountAccessConsent(permissionsRes.Permissions); err != nil {
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

		if loginURL, err = s.WebClient.LoginURL(*registerResponse.Data.ConsentID, challenge, []string{"openid", "accounts"}, encodedNonce); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
		}

		data["login_url"] = loginURL

		c.JSON(http.StatusOK, data)
	}
}

func (s *Server) Callback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app        string
			appStorage = AppStorage{}
			code       = c.Query("code")
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

		if token, err = s.WebClient.Exchange(code, appStorage.Verifier); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		c.SetCookie("token", token.AccessToken, 0, "/", "", false, true)

		c.Redirect(http.StatusFound, os.Getenv("FINANCROO_UI_URL"))
	}
}
