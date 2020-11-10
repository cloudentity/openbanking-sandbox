// Code generated by go-swagger; DO NOT EDIT.

package offers

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

// NewGetOffersParams creates a new GetOffersParams object
// with the default values initialized.
func NewGetOffersParams() *GetOffersParams {
	var ()
	return &GetOffersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOffersParamsWithTimeout creates a new GetOffersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOffersParamsWithTimeout(timeout time.Duration) *GetOffersParams {
	var ()
	return &GetOffersParams{

		timeout: timeout,
	}
}

// NewGetOffersParamsWithContext creates a new GetOffersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOffersParamsWithContext(ctx context.Context) *GetOffersParams {
	var ()
	return &GetOffersParams{

		Context: ctx,
	}
}

// NewGetOffersParamsWithHTTPClient creates a new GetOffersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOffersParamsWithHTTPClient(client *http.Client) *GetOffersParams {
	var ()
	return &GetOffersParams{
		HTTPClient: client,
	}
}

/*GetOffersParams contains all the parameters to send to the API endpoint
for the get offers operation typically these are written to a http.Request
*/
type GetOffersParams struct {

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

// WithTimeout adds the timeout to the get offers params
func (o *GetOffersParams) WithTimeout(timeout time.Duration) *GetOffersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get offers params
func (o *GetOffersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get offers params
func (o *GetOffersParams) WithContext(ctx context.Context) *GetOffersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get offers params
func (o *GetOffersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get offers params
func (o *GetOffersParams) WithHTTPClient(client *http.Client) *GetOffersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get offers params
func (o *GetOffersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get offers params
func (o *GetOffersParams) WithAuthorization(authorization string) *GetOffersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get offers params
func (o *GetOffersParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the get offers params
func (o *GetOffersParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *GetOffersParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the get offers params
func (o *GetOffersParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the get offers params
func (o *GetOffersParams) WithXFapiAuthDate(xFapiAuthDate *string) *GetOffersParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the get offers params
func (o *GetOffersParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the get offers params
func (o *GetOffersParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *GetOffersParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the get offers params
func (o *GetOffersParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the get offers params
func (o *GetOffersParams) WithXFapiInteractionID(xFapiInteractionID *string) *GetOffersParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the get offers params
func (o *GetOffersParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOffersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
