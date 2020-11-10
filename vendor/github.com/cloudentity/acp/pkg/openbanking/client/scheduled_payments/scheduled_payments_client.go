// Code generated by go-swagger; DO NOT EDIT.

package scheduled_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scheduled payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scheduled payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAccountsAccountIDScheduledPayments(params *GetAccountsAccountIDScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDScheduledPaymentsOK, error)

	GetScheduledPayments(params *GetScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetScheduledPaymentsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAccountsAccountIDScheduledPayments gets scheduled payments
*/
func (a *Client) GetAccountsAccountIDScheduledPayments(params *GetAccountsAccountIDScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsAccountIDScheduledPaymentsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetAccountsAccountIDScheduledPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAccountsAccountIdScheduledPayments",
		Method:             "GET",
		PathPattern:        "/accounts/{AccountId}/scheduled-payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsAccountIDScheduledPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountsAccountIDScheduledPaymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAccountsAccountIdScheduledPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetScheduledPayments gets scheduled payments
*/
func (a *Client) GetScheduledPayments(params *GetScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*GetScheduledPaymentsOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetScheduledPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetScheduledPayments",
		Method:             "GET",
		PathPattern:        "/scheduled-payments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScheduledPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScheduledPaymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScheduledPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
