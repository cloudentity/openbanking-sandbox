package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
)

func TestSignRequest(t *testing.T) {
	bs, err := ioutil.ReadFile("../data/tpp_key.pem")
	require.NoError(t, err)

	block, _ := pem.Decode(bs)
	signingKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	require.NoError(t, err)

	request := Request{
		ClientID: "123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "issuer",
		},
	}

	_, err = SignRequest(request, signingKey)
	require.NoError(t, err)
}
