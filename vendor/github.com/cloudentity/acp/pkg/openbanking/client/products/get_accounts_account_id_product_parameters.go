// Code generated by go-swagger; DO NOT EDIT.

package products

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

// NewGetAccountsAccountIDProductParams creates a new GetAccountsAccountIDProductParams object
// with the default values initialized.
func NewGetAccountsAccountIDProductParams() *GetAccountsAccountIDProductParams {
	var ()
	return &GetAccountsAccountIDProductParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDProductParamsWithTimeout creates a new GetAccountsAccountIDProductParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsAccountIDProductParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDProductParams {
	var ()
	return &GetAccountsAccountIDProductParams{

		timeout: timeout,
	}
}

// NewGetAccountsAccountIDProductParamsWithContext creates a new GetAccountsAccountIDProductParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsAccountIDProductParamsWithContext(ctx context.Context) *GetAccountsAccountIDProductParams {
	var ()
	return &GetAccountsAccountIDProductParams{

		Context: ctx,
	}
}

// NewGetAccountsAccountIDProductParamsWithHTTPClient creates a new GetAccountsAccountIDProductParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountsAccountIDProductParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDProductParams {
	var ()
	return &GetAccountsAccountIDProductParams{
		HTTPClient: client,
	}
}

/*GetAccountsAccountIDProductParams contains all the parameters to send to the API endpoint
for the get accounts account Id product operation typically these are written to a http.Request
*/
type GetAccountsAccountIDProductParams struct {

	/*AccountID
	  AccountId

	*/
	AccountID string
	/*Authorization
	  An Authorisation Token as per https://tools.ietf.org/html/rfc6750

	*/
	Authorization string
	/*XCustomerUserAgent
	  Indicates the user-agent that the PSU is using.

	*/
	XCustomerUserAgent *string
	/*XFapiAuthDate
	  The time when the PSU last logged in with the TPP.
	All dates in the HTTP headers are represented as RFC 7231 Full Dates. An example is below:
	Sun, 10 Sep 2017 19:43:31 UTC

	*/
	XFapiAuthDate *string
	/*XFapiCustomerIPAddress
	  The PSU's IP address if the PSU is currently logged in with the TPP.

	*/
	XFapiCustomerIPAddress *string
	/*XFapiInteractionID
	  An RFC4122 UID used as a correlation id.

	*/
	XFapiInteractionID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDProductParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithContext(ctx context.Context) *GetAccountsAccountIDProductParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDProductParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithAccountID(accountID string) *GetAccountsAccountIDProductParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithAuthorization adds the authorization to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithAuthorization(authorization string) *GetAccountsAccountIDProductParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountsAccountIDProductParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountsAccountIDProductParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountsAccountIDProductParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountsAccountIDProductParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get accounts account Id product params
func (o *GetAccountsAccountIDProductParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDProductParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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