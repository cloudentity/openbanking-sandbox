// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewGetAccountsAccountIDStatementsStatementIDTransactionsParams creates a new GetAccountsAccountIDStatementsStatementIDTransactionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountsAccountIDStatementsStatementIDTransactionsParams() *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsParamsWithTimeout creates a new GetAccountsAccountIDStatementsStatementIDTransactionsParams object
// with the ability to set a timeout on a request.
func NewGetAccountsAccountIDStatementsStatementIDTransactionsParamsWithTimeout(timeout time.Duration) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsParams{
		timeout: timeout,
	}
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsParamsWithContext creates a new GetAccountsAccountIDStatementsStatementIDTransactionsParams object
// with the ability to set a context for a request.
func NewGetAccountsAccountIDStatementsStatementIDTransactionsParamsWithContext(ctx context.Context) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsParams{
		Context: ctx,
	}
}

// NewGetAccountsAccountIDStatementsStatementIDTransactionsParamsWithHTTPClient creates a new GetAccountsAccountIDStatementsStatementIDTransactionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountsAccountIDStatementsStatementIDTransactionsParamsWithHTTPClient(client *http.Client) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	return &GetAccountsAccountIDStatementsStatementIDTransactionsParams{
		HTTPClient: client,
	}
}

/* GetAccountsAccountIDStatementsStatementIDTransactionsParams contains all the parameters to send to the API endpoint
   for the get accounts account Id statements statement Id transactions operation.

   Typically these are written to a http.Request.
*/
type GetAccountsAccountIDStatementsStatementIDTransactionsParams struct {

	/* AccountID.

	   AccountId
	*/
	AccountID string

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* StatementID.

	   StatementId
	*/
	StatementID string

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

// WithDefaults hydrates default values in the get accounts account Id statements statement Id transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithDefaults() *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get accounts account Id statements statement Id transactions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithTimeout(timeout time.Duration) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithContext(ctx context.Context) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithHTTPClient(client *http.Client) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithAccountID(accountID string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithAuthorization adds the authorization to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithAuthorization(authorization string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithStatementID adds the statementID to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithStatementID(statementID string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetStatementID(statementID)
	return o
}

// SetStatementID adds the statementId to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetStatementID(statementID string) {
	o.StatementID = statementID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetAccountsAccountIDStatementsStatementIDTransactionsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get accounts account Id statements statement Id transactions params
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsAccountIDStatementsStatementIDTransactionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param StatementId
	if err := r.SetPathParam("StatementId", o.StatementID); err != nil {
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
