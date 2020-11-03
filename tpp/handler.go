package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const challengeLength = 43

func (s *Server) Get() func(*gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
	}
}

func (s *Server) Login() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			intentID           string
			encodedVerifier    string
			encodedCookieValue string
			challenge          string
			err                error
		)

		requestPermissions := strings.Split(c.PostForm("permissions"), ",")

		if intentID, err = s.Client.RegisterAccountAccessConsent(requestPermissions); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to register account access consent: %+v", err))
			return
		}

		verifier := make([]byte, challengeLength)
		if _, err = io.ReadFull(rand.Reader, verifier); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to generate challenge: %+v", err))
			return
		}

		encodedVerifier = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(verifier)
		if encodedCookieValue, err = s.SecureCookie.Encode("verifier", encodedVerifier); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error while encoding cookie: %+v", err))
			return
		}

		c.SetCookie("verifier", encodedCookieValue, 0, "/", "", false, true)

		hash := sha256.New()
		hash.Write([]byte(encodedVerifier))
		challenge = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash.Sum([]byte{}))

		loginURL := s.WebClient.AuthorizeURL(intentID, challenge, []string{"accounts"})

		c.HTML(http.StatusOK, "intent_registered.tmpl", gin.H{
			"intent_id": intentID,
			"login_url": loginURL,
		})
	}
}

func (s *Server) Callback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			verfier       string
			verifierValue string
			code          = c.Query("code")
			body          []byte
			err           error
		)

		if verfier, err = c.Cookie("verifier"); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to get verifier from cookie: %+v", err))
			return
		}

		if err = s.SecureCookie.Decode("verifier", verfier, &verifierValue); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("failed to decode verifier: %+v", err))
			return
		}

		if body, err = s.WebClient.Exchange(code, verifierValue); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to exchange code: %+v", err))
			return
		}

		c.HTML(http.StatusOK, "authenticated.tmpl", gin.H{
			"token_response": string(body),
		})
	}
}
