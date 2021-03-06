// Code generated by go-swagger; DO NOT EDIT.

package payment_details

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

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams creates a new GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParamsWithTimeout creates a new GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams object
// with the ability to set a timeout on a request.
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParamsWithTimeout(timeout time.Duration) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams{
		timeout: timeout,
	}
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParamsWithContext creates a new GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams object
// with the ability to set a context for a request.
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParamsWithContext(ctx context.Context) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams{
		Context: ctx,
	}
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParamsWithHTTPClient creates a new GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParamsWithHTTPClient(client *http.Client) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams{
		HTTPClient: client,
	}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams contains all the parameters to send to the API endpoint
   for the get international standing orders international standing order payment Id payment details operation.

   Typically these are written to a http.Request.
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* InternationalStandingOrderPaymentID.

	   InternationalStandingOrderPaymentId
	*/
	InternationalStandingOrderPaymentID string

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

// WithDefaults hydrates default values in the get international standing orders international standing order payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithDefaults() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get international standing orders international standing order payment Id payment details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithTimeout(timeout time.Duration) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithContext(ctx context.Context) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithHTTPClient(client *http.Client) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithAuthorization(authorization string) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithInternationalStandingOrderPaymentID adds the internationalStandingOrderPaymentID to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithInternationalStandingOrderPaymentID(internationalStandingOrderPaymentID string) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetInternationalStandingOrderPaymentID(internationalStandingOrderPaymentID)
	return o
}

// SetInternationalStandingOrderPaymentID adds the internationalStandingOrderPaymentId to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetInternationalStandingOrderPaymentID(internationalStandingOrderPaymentID string) {
	o.InternationalStandingOrderPaymentID = internationalStandingOrderPaymentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get international standing orders international standing order payment Id payment details params
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDPaymentDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param InternationalStandingOrderPaymentId
	if err := r.SetPathParam("InternationalStandingOrderPaymentId", o.InternationalStandingOrderPaymentID); err != nil {
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
