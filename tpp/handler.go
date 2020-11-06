package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	challengeLength = 43
	nonceLength     = 20
)

type AppStorage struct {
	Verifier string
	Nonce    string
}

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
			encodedNonce       string
			encodedCookieValue string
			challenge          string
			loginURL           string
			u                  *url.URL
			data               = gin.H{}
			err                error
		)

		requestPermissions := strings.Split(c.PostForm("permissions"), ",")

		if intentID, err = s.Client.RegisterAccountAccessConsent(requestPermissions); err != nil {
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
		hash.Write([]byte(encodedVerifier))
		challenge = base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash.Sum([]byte{}))

		if loginURL, err = s.WebClient.AuthorizeURL(intentID, challenge, []string{"openid", "accounts"}, encodedNonce); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to build authorize url: %+v", err))
		}

		// debug only
		if u, err = url.Parse(loginURL); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("failed to parse login url: %+v", err))
		}
		rp := u.Query()["request"]
		if len(rp) > 0 {
			parser := jwt.Parser{}
			claims := jwt.MapClaims{}
			idToken, _, _ := parser.ParseUnverified(rp[0], &claims)
			header, _ := json.MarshalIndent(idToken.Header, "", "  ")
			payload, _ := json.MarshalIndent(claims, "", "  ")

			data["request_raw"] = rp[0]
			data["request_header"] = string(header)
			data["request_payload"] = string(payload)
		}

		data["intent_id"] = intentID
		data["login_url"] = loginURL

		c.HTML(http.StatusOK, "intent_registered.tmpl", data)
	}
}

func (s *Server) Callback() func(*gin.Context) {
	return func(c *gin.Context) {
		var (
			app              string
			appStorage       = AppStorage{}
			userinfoResponse map[string]interface{}
			code             = c.Query("code")
			token            Token
			data             = gin.H{}
			err              error
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

		// todo validate id_token and compare nonce val

		if userinfoResponse, err = s.WebClient.Userinfo(token.AccessToken); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("failed to introspect access token: %+v", err))
			return
		}

		data["access_token"] = token.AccessToken
		userinfoResp, _ := json.MarshalIndent(userinfoResponse, "", "  ")
		data["userinfo"] = string(userinfoResp)

		if token.IDToken != nil {
			parser := jwt.Parser{}
			claims := jwt.MapClaims{}
			idToken, _, _ := parser.ParseUnverified(*token.IDToken, &claims)
			header, _ := json.MarshalIndent(idToken.Header, "", "  ")
			payload, _ := json.MarshalIndent(claims, "", "  ")

			data["id_token_raw"] = idToken
			data["id_token_header"] = string(header)
			data["id_token_payload"] = string(payload)
		}

		c.HTML(http.StatusOK, "authenticated.tmpl", data)
	}
}
