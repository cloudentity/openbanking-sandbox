// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/account_access"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/accounts"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/balances"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/beneficiaries"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/direct_debits"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/offers"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/parties"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/products"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/scheduled_payments"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/standing_orders"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/statements"
	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/client/transactions"
)

// Default openbanking accounts client HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/open-banking/v3.1/aisp"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new openbanking accounts client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *OpenbankingAccountsClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new openbanking accounts client HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *OpenbankingAccountsClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new openbanking accounts client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *OpenbankingAccountsClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(OpenbankingAccountsClient)
	cli.Transport = transport
	cli.AccountAccess = account_access.New(transport, formats)
	cli.Accounts = accounts.New(transport, formats)
	cli.Balances = balances.New(transport, formats)
	cli.Beneficiaries = beneficiaries.New(transport, formats)
	cli.DirectDebits = direct_debits.New(transport, formats)
	cli.Offers = offers.New(transport, formats)
	cli.Parties = parties.New(transport, formats)
	cli.Products = products.New(transport, formats)
	cli.ScheduledPayments = scheduled_payments.New(transport, formats)
	cli.StandingOrders = standing_orders.New(transport, formats)
	cli.Statements = statements.New(transport, formats)
	cli.Transactions = transactions.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// OpenbankingAccountsClient is a client for openbanking accounts client
type OpenbankingAccountsClient struct {
	AccountAccess account_access.ClientService

	Accounts accounts.ClientService

	Balances balances.ClientService

	Beneficiaries beneficiaries.ClientService

	DirectDebits direct_debits.ClientService

	Offers offers.ClientService

	Parties parties.ClientService

	Products products.ClientService

	ScheduledPayments scheduled_payments.ClientService

	StandingOrders standing_orders.ClientService

	Statements statements.ClientService

	Transactions transactions.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *OpenbankingAccountsClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccountAccess.SetTransport(transport)
	c.Accounts.SetTransport(transport)
	c.Balances.SetTransport(transport)
	c.Beneficiaries.SetTransport(transport)
	c.DirectDebits.SetTransport(transport)
	c.Offers.SetTransport(transport)
	c.Parties.SetTransport(transport)
	c.Products.SetTransport(transport)
	c.ScheduledPayments.SetTransport(transport)
	c.StandingOrders.SetTransport(transport)
	c.Statements.SetTransport(transport)
	c.Transactions.SetTransport(transport)
}