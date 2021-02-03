// Code generated by go-swagger; DO NOT EDIT.

package domestic_standing_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new domestic standing orders API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for domestic standing orders API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDomesticStandingOrderConsents(params *CreateDomesticStandingOrderConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticStandingOrderConsentsCreated, error)

	CreateDomesticStandingOrders(params *CreateDomesticStandingOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticStandingOrdersCreated, error)

	GetDomesticStandingOrderConsentsConsentID(params *GetDomesticStandingOrderConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticStandingOrderConsentsConsentIDOK, error)

	GetDomesticStandingOrdersDomesticStandingOrderID(params *GetDomesticStandingOrdersDomesticStandingOrderIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticStandingOrdersDomesticStandingOrderIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDomesticStandingOrderConsents creates domestic standing order consents
*/
func (a *Client) CreateDomesticStandingOrderConsents(params *CreateDomesticStandingOrderConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticStandingOrderConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomesticStandingOrderConsentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDomesticStandingOrderConsents",
		Method:             "POST",
		PathPattern:        "/domestic-standing-order-consents",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomesticStandingOrderConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDomesticStandingOrderConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDomesticStandingOrderConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDomesticStandingOrders creates domestic standing orders
*/
func (a *Client) CreateDomesticStandingOrders(params *CreateDomesticStandingOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticStandingOrdersCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomesticStandingOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDomesticStandingOrders",
		Method:             "POST",
		PathPattern:        "/domestic-standing-orders",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomesticStandingOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDomesticStandingOrdersCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDomesticStandingOrders: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticStandingOrderConsentsConsentID gets domestic standing order consents
*/
func (a *Client) GetDomesticStandingOrderConsentsConsentID(params *GetDomesticStandingOrderConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticStandingOrderConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticStandingOrderConsentsConsentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomesticStandingOrderConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/domestic-standing-order-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticStandingOrderConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticStandingOrderConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticStandingOrderConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticStandingOrdersDomesticStandingOrderID gets domestic standing orders
*/
func (a *Client) GetDomesticStandingOrdersDomesticStandingOrderID(params *GetDomesticStandingOrdersDomesticStandingOrderIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticStandingOrdersDomesticStandingOrderIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticStandingOrdersDomesticStandingOrderIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomesticStandingOrdersDomesticStandingOrderId",
		Method:             "GET",
		PathPattern:        "/domestic-standing-orders/{DomesticStandingOrderId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticStandingOrdersDomesticStandingOrderIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticStandingOrdersDomesticStandingOrderIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticStandingOrdersDomesticStandingOrderId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
