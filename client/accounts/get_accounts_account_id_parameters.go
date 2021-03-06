// Code generated by go-swagger; DO NOT EDIT.

package accounts

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

// NewGetAccountsAccountIDParams creates a new GetAccountsAccountIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountsAccountIDParams() *GetAccountsAccountIDParams {
	return &GetAccountsAccountIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDParamsWithTimeout creates a new GetAccountsAccountIDParams object
// with the ability to set a timeout on a request.
func NewGetAccountsAccountIDParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDParams {
	return &GetAccountsAccountIDParams{
		timeout: timeout,
	}
}

// NewGetAccountsAccountIDParamsWithContext creates a new GetAccountsAccountIDParams object
// with the ability to set a context for a request.
func NewGetAccountsAccountIDParamsWithContext(ctx context.Context) *GetAccountsAccountIDParams {
	return &GetAccountsAccountIDParams{
		Context: ctx,
	}
}

// NewGetAccountsAccountIDParamsWithHTTPClient creates a new GetAccountsAccountIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountsAccountIDParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDParams {
	return &GetAccountsAccountIDParams{
		HTTPClient: client,
	}
}

/* GetAccountsAccountIDParams contains all the parameters to send to the API endpoint
   for the get accounts account Id operation.

   Typically these are written to a http.Request.
*/
type GetAccountsAccountIDParams struct {

	/* AccountID.

	   AccountId
	*/
	AccountID string

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

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

// WithDefaults hydrates default values in the get accounts account Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsAccountIDParams) WithDefaults() *GetAccountsAccountIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accounts account Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsAccountIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithContext(ctx context.Context) *GetAccountsAccountIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithAccountID(accountID string) *GetAccountsAccountIDParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithAuthorization adds the authorization to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithAuthorization(authorization string) *GetAccountsAccountIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountsAccountIDParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountsAccountIDParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountsAccountIDParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get accounts account Id params
func (o *GetAccountsAccountIDParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountsAccountIDParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get accounts account Id params
func (o *GetAccountsAccountIDParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param AccountId
	if err := r.SetPathParam("AccountId", o.AccountID); err != nil {
		return err
	}

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
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
