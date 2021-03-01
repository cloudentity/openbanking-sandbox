// Code generated by go-swagger; DO NOT EDIT.

package international_scheduled_payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new international scheduled payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for international scheduled payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInternationalScheduledPaymentConsents(params *CreateInternationalScheduledPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInternationalScheduledPaymentConsentsCreated, error)

	CreateInternationalScheduledPayments(params *CreateInternationalScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInternationalScheduledPaymentsCreated, error)

	GetInternationalScheduledPaymentConsentsConsentID(params *GetInternationalScheduledPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternationalScheduledPaymentConsentsConsentIDOK, error)

	GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmation(params *GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationOK, error)

	GetInternationalScheduledPaymentsInternationalScheduledPaymentID(params *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternationalScheduledPaymentsInternationalScheduledPaymentIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateInternationalScheduledPaymentConsents creates international scheduled payment consents
*/
func (a *Client) CreateInternationalScheduledPaymentConsents(params *CreateInternationalScheduledPaymentConsentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInternationalScheduledPaymentConsentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternationalScheduledPaymentConsentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateInternationalScheduledPaymentConsents",
		Method:             "POST",
		PathPattern:        "/international-scheduled-payment-consents",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternationalScheduledPaymentConsentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateInternationalScheduledPaymentConsentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateInternationalScheduledPaymentConsents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateInternationalScheduledPayments creates international scheduled payments
*/
func (a *Client) CreateInternationalScheduledPayments(params *CreateInternationalScheduledPaymentsParams, authInfo runtime.ClientAuthInfoWriter) (*CreateInternationalScheduledPaymentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInternationalScheduledPaymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateInternationalScheduledPayments",
		Method:             "POST",
		PathPattern:        "/international-scheduled-payments",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateInternationalScheduledPaymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateInternationalScheduledPaymentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateInternationalScheduledPayments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalScheduledPaymentConsentsConsentID gets international scheduled payment consents
*/
func (a *Client) GetInternationalScheduledPaymentConsentsConsentID(params *GetInternationalScheduledPaymentConsentsConsentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternationalScheduledPaymentConsentsConsentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalScheduledPaymentConsentsConsentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInternationalScheduledPaymentConsentsConsentId",
		Method:             "GET",
		PathPattern:        "/international-scheduled-payment-consents/{ConsentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalScheduledPaymentConsentsConsentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternationalScheduledPaymentConsentsConsentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalScheduledPaymentConsentsConsentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmation gets international scheduled payment consents funds confirmation
*/
func (a *Client) GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmation(params *GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInternationalScheduledPaymentConsentsConsentIdFundsConfirmation",
		Method:             "GET",
		PathPattern:        "/international-scheduled-payment-consents/{ConsentId}/funds-confirmation",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternationalScheduledPaymentConsentsConsentIDFundsConfirmationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalScheduledPaymentConsentsConsentIdFundsConfirmation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInternationalScheduledPaymentsInternationalScheduledPaymentID gets international scheduled payments
*/
func (a *Client) GetInternationalScheduledPaymentsInternationalScheduledPaymentID(params *GetInternationalScheduledPaymentsInternationalScheduledPaymentIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetInternationalScheduledPaymentsInternationalScheduledPaymentIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInternationalScheduledPaymentsInternationalScheduledPaymentIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInternationalScheduledPaymentsInternationalScheduledPaymentId",
		Method:             "GET",
		PathPattern:        "/international-scheduled-payments/{InternationalScheduledPaymentId}",
		ProducesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		ConsumesMediaTypes: []string{"application/jose+jwe", "application/json; charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInternationalScheduledPaymentsInternationalScheduledPaymentIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInternationalScheduledPaymentsInternationalScheduledPaymentIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInternationalScheduledPaymentsInternationalScheduledPaymentId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
