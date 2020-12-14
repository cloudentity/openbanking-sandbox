// Code generated by go-swagger; DO NOT EDIT.

package scheduled_payments

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

// NewGetAccountsAccountIDScheduledPaymentsParams creates a new GetAccountsAccountIDScheduledPaymentsParams object
// with the default values initialized.
func NewGetAccountsAccountIDScheduledPaymentsParams() *GetAccountsAccountIDScheduledPaymentsParams {
	var ()
	return &GetAccountsAccountIDScheduledPaymentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDScheduledPaymentsParamsWithTimeout creates a new GetAccountsAccountIDScheduledPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsAccountIDScheduledPaymentsParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDScheduledPaymentsParams {
	var ()
	return &GetAccountsAccountIDScheduledPaymentsParams{

		timeout: timeout,
	}
}

// NewGetAccountsAccountIDScheduledPaymentsParamsWithContext creates a new GetAccountsAccountIDScheduledPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsAccountIDScheduledPaymentsParamsWithContext(ctx context.Context) *GetAccountsAccountIDScheduledPaymentsParams {
	var ()
	return &GetAccountsAccountIDScheduledPaymentsParams{

		Context: ctx,
	}
}

// NewGetAccountsAccountIDScheduledPaymentsParamsWithHTTPClient creates a new GetAccountsAccountIDScheduledPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountsAccountIDScheduledPaymentsParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDScheduledPaymentsParams {
	var ()
	return &GetAccountsAccountIDScheduledPaymentsParams{
		HTTPClient: client,
	}
}

/*GetAccountsAccountIDScheduledPaymentsParams contains all the parameters to send to the API endpoint
for the get accounts account Id scheduled payments operation typically these are written to a http.Request
*/
type GetAccountsAccountIDScheduledPaymentsParams struct {

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

// WithTimeout adds the timeout to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithContext(ctx context.Context) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithAccountID(accountID string) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithAuthorization adds the authorization to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithAuthorization(authorization string) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountsAccountIDScheduledPaymentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get accounts account Id scheduled payments params
func (o *GetAccountsAccountIDScheduledPaymentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDScheduledPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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