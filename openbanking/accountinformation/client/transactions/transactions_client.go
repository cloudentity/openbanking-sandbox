// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountsAccountIDStatementsStatementIDTransactions(params *GetAccountsAccountIDStatementsStatementIDTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDStatementsStatementIDTransactionsOK, error)

	GetAccountsAccountIDTransactions(params *GetAccountsAccountIDTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDTransactionsOK, error)

	GetTransactions(params *GetTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransactionsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountsAccountIDStatementsStatementIDTransactions gets transactions
*/
func (a *Client) GetAccountsAccountIDStatementsStatementIDTransactions(params *GetAccountsAccountIDStatementsStatementIDTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDStatementsStatementIDTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsAccountIDStatementsStatementIDTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountsAccountIdStatementsStatementIdTransactions",
		Method:             "GET",
		PathPattern:        "/accounts/{AccountId}/statements/{StatementId}/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsAccountIDStatementsStatementIDTransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsAccountIDStatementsStatementIDTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountsAccountIdStatementsStatementIdTransactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAccountsAccountIDTransactions gets transactions
*/
func (a *Client) GetAccountsAccountIDTransactions(params *GetAccountsAccountIDTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsAccountIDTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountsAccountIdTransactions",
		Method:             "GET",
		PathPattern:        "/accounts/{AccountId}/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsAccountIDTransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsAccountIDTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountsAccountIdTransactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTransactions gets transactions
*/
func (a *Client) GetTransactions(params *GetTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactions",
		Method:             "GET",
		PathPattern:        "/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTransactionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTransactions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
