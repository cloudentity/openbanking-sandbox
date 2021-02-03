// Code generated by go-swagger; DO NOT EDIT.

package domestic_scheduled_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new domestic scheduled payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for domestic scheduled payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDomesticScheduledPaymentConsents(params *CreateDomesticScheduledPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticScheduledPaymentConsentsCreated, error)

	CreateDomesticScheduledPayments(params *CreateDomesticScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticScheduledPaymentsCreated, error)

	GetDomesticScheduledPaymentConsentsConsentID(params *GetDomesticScheduledPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticScheduledPaymentConsentsConsentIDOK, error)

	GetDomesticScheduledPaymentsDomesticScheduledPaymentID(params *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticScheduledPaymentsDomesticScheduledPaymentIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDomesticScheduledPaymentConsents creates domestic scheduled payment consents
*/
func (a *Client) CreateDomesticScheduledPaymentConsents(params *CreateDomesticScheduledPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticScheduledPaymentConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomesticScheduledPaymentConsentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDomesticScheduledPaymentConsents",
		Method:             "POST",
		PathPattern:        "/domestic-scheduled-payment-consents",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomesticScheduledPaymentConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDomesticScheduledPaymentConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDomesticScheduledPaymentConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateDomesticScheduledPayments creates domestic scheduled payments
*/
func (a *Client) CreateDomesticScheduledPayments(params *CreateDomesticScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDomesticScheduledPaymentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDomesticScheduledPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDomesticScheduledPayments",
		Method:             "POST",
		PathPattern:        "/domestic-scheduled-payments",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDomesticScheduledPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDomesticScheduledPaymentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateDomesticScheduledPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticScheduledPaymentConsentsConsentID gets domestic scheduled payment consents
*/
func (a *Client) GetDomesticScheduledPaymentConsentsConsentID(params *GetDomesticScheduledPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticScheduledPaymentConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticScheduledPaymentConsentsConsentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomesticScheduledPaymentConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/domestic-scheduled-payment-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticScheduledPaymentConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticScheduledPaymentConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticScheduledPaymentConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDomesticScheduledPaymentsDomesticScheduledPaymentID gets domestic scheduled payments
*/
func (a *Client) GetDomesticScheduledPaymentsDomesticScheduledPaymentID(params *GetDomesticScheduledPaymentsDomesticScheduledPaymentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetDomesticScheduledPaymentsDomesticScheduledPaymentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDomesticScheduledPaymentsDomesticScheduledPaymentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDomesticScheduledPaymentsDomesticScheduledPaymentId",
		Method:             "GET",
		PathPattern:        "/domestic-scheduled-payments/{DomesticScheduledPaymentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDomesticScheduledPaymentsDomesticScheduledPaymentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDomesticScheduledPaymentsDomesticScheduledPaymentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDomesticScheduledPaymentsDomesticScheduledPaymentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
