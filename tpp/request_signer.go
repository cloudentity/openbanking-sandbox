package main

import "github.com/dgrijalva/jwt-go"

type Request struct {
	ClientID     string `json:"client_id"`
	Scope        string `json:"scope"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Claims       Claims `json:"claims"`
	Nonce        string `json:"nonce"`
}

func (r Request) Valid() error {
	return nil
}

type Claim struct {
	Essential bool   `json:"essential"`
	Value     string `json:"value"`
}

type Claims struct {
	IdToken  map[string]Claim `json:"id_token"`
	Userinfo map[string]Claim `json:"userinfo"`
}

func SignRequest(request Request, signingKey interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, request)
	return token.SignedString(signingKey)
}
