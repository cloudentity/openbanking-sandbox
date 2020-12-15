package main

import (
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"

	obc "github.com/cloudentity/openbanking-sandbox/client"
)

type OpenbankingClient struct {
	*obc.Openbanking
}

func NewOpenbankingClient(config Config) OpenbankingClient {
	var (
		c  = OpenbankingClient{}
		hc = &http.Client{}
	)

	c.Openbanking = obc.New(httptransport.NewWithClient(
		config.BankURL.Host,
		"/",
		[]string{config.BankURL.Scheme},
		hc,
	), nil)

	return c
}
