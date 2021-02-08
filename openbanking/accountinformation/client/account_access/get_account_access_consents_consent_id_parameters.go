// Code generated by go-swagger; DO NOT EDIT.

package account_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAccountAccessConsentsConsentIDParams creates a new GetAccountAccessConsentsConsentIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountAccessConsentsConsentIDParams() *GetAccountAccessConsentsConsentIDParams {
	return &GetAccountAccessConsentsConsentIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountAccessConsentsConsentIDParamsWithTimeout creates a new GetAccountAccessConsentsConsentIDParams object
// with the ability to set a timeout on a request.
func NewGetAccountAccessConsentsConsentIDParamsWithTimeout(timeout time.Duration) *GetAccountAccessConsentsConsentIDParams {
	return &GetAccountAccessConsentsConsentIDParams{
		timeout: timeout,
	}
}

// NewGetAccountAccessConsentsConsentIDParamsWithContext creates a new GetAccountAccessConsentsConsentIDParams object
// with the ability to set a context for a request.
func NewGetAccountAccessConsentsConsentIDParamsWithContext(ctx context.Context) *GetAccountAccessConsentsConsentIDParams {
	return &GetAccountAccessConsentsConsentIDParams{
		Context: ctx,
	}
}

// NewGetAccountAccessConsentsConsentIDParamsWithHTTPClient creates a new GetAccountAccessConsentsConsentIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountAccessConsentsConsentIDParamsWithHTTPClient(client *http.Client) *GetAccountAccessConsentsConsentIDParams {
	return &GetAccountAccessConsentsConsentIDParams{
		HTTPClient: client,
	}
}

/* GetAccountAccessConsentsConsentIDParams contains all the parameters to send to the API endpoint
   for the get account access consents consent Id operation.

   Typically these are written to a http.Request.
*/
type GetAccountAccessConsentsConsentIDParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* ConsentID.

	   ConsentId
	*/
	ConsentID string

	/* XCustomerUserAgent.

	   Indicates the user-agent that the PSU is using.
	*/
	XCustomerUserAgent *string

	/* XFapiAuthDate.

	     The time when the PSU last logged in with the TPP.
	All dates in the HTTP headers are represented as RFC 7231 Full Dates. An example is below:
	Sun, 10 Sep 2017 19:43:31 UTC
	*/
	XFapiAuthDate *string

	/* XFapiCustomerIPAddress.

	   The PSU's IP address if the PSU is currently logged in with the TPP.
	*/
	XFapiCustomerIPAddress *string

	/* XFapiInteractionID.

	   An RFC4122 UID used as a correlation id.
	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account access consents consent Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountAccessConsentsConsentIDParams) WithDefaults() *GetAccountAccessConsentsConsentIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account access consents consent Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountAccessConsentsConsentIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithTimeout(timeout time.Duration) *GetAccountAccessConsentsConsentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithContext(ctx context.Context) *GetAccountAccessConsentsConsentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithHTTPClient(client *http.Client) *GetAccountAccessConsentsConsentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithAuthorization(authorization string) *GetAccountAccessConsentsConsentIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithConsentID adds the consentID to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithConsentID(consentID string) *GetAccountAccessConsentsConsentIDParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountAccessConsentsConsentIDParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountAccessConsentsConsentIDParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountAccessConsentsConsentIDParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountAccessConsentsConsentIDParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get account access consents consent Id params
func (o *GetAccountAccessConsentsConsentIDParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountAccessConsentsConsentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param ConsentId
	if err := r.SetPathParam("ConsentId", o.ConsentID); err != nil {
		return err
	}

	if o.XCustomerUserAgent != nil {

		// header param x-customer-user-agent
		if err := r.SetHeaderParam("x-customer-user-agent", *o.XCustomerUserAgent); err != nil {
			return err
		}
	}

	if o.XFapiAuthDate != nil {

		// header param x-fapi-auth-date
		if err := r.SetHeaderParam("x-fapi-auth-date", *o.XFapiAuthDate); err != nil {
			return err
		}
	}

	if o.XFapiCustomerIPAddress != nil {

		// header param x-fapi-customer-ip-address
		if err := r.SetHeaderParam("x-fapi-customer-ip-address", *o.XFapiCustomerIPAddress); err != nil {
			return err
		}
	}

	if o.XFapiInteractionID != nil {

		// header param x-fapi-interaction-id
		if err := r.SetHeaderParam("x-fapi-interaction-id", *o.XFapiInteractionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}