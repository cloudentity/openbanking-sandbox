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

// NewDeleteAccountAccessConsentsConsentIDParams creates a new DeleteAccountAccessConsentsConsentIDParams object
// with the default values initialized.
func NewDeleteAccountAccessConsentsConsentIDParams() *DeleteAccountAccessConsentsConsentIDParams {
	var ()
	return &DeleteAccountAccessConsentsConsentIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccountAccessConsentsConsentIDParamsWithTimeout creates a new DeleteAccountAccessConsentsConsentIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAccountAccessConsentsConsentIDParamsWithTimeout(timeout time.Duration) *DeleteAccountAccessConsentsConsentIDParams {
	var ()
	return &DeleteAccountAccessConsentsConsentIDParams{

		timeout: timeout,
	}
}

// NewDeleteAccountAccessConsentsConsentIDParamsWithContext creates a new DeleteAccountAccessConsentsConsentIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAccountAccessConsentsConsentIDParamsWithContext(ctx context.Context) *DeleteAccountAccessConsentsConsentIDParams {
	var ()
	return &DeleteAccountAccessConsentsConsentIDParams{

		Context: ctx,
	}
}

// NewDeleteAccountAccessConsentsConsentIDParamsWithHTTPClient creates a new DeleteAccountAccessConsentsConsentIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAccountAccessConsentsConsentIDParamsWithHTTPClient(client *http.Client) *DeleteAccountAccessConsentsConsentIDParams {
	var ()
	return &DeleteAccountAccessConsentsConsentIDParams{
		HTTPClient: client,
	}
}

/*DeleteAccountAccessConsentsConsentIDParams contains all the parameters to send to the API endpoint
for the delete account access consents consent Id operation typically these are written to a http.Request
*/
type DeleteAccountAccessConsentsConsentIDParams struct {

	/*Authorization
	  An Authorisation Token as per https://tools.ietf.org/html/rfc6750

	*/
	Authorization string
	/*ConsentID
	  ConsentId

	*/
	ConsentID string
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

// WithTimeout adds the timeout to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithTimeout(timeout time.Duration) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithContext(ctx context.Context) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithHTTPClient(client *http.Client) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithAuthorization(authorization string) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithConsentID adds the consentID to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithConsentID(consentID string) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithXFapiAuthDate(xFapiAuthDate *string) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) WithXFapiInteractionID(xFapiInteractionID *string) *DeleteAccountAccessConsentsConsentIDParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the delete account access consents consent Id params
func (o *DeleteAccountAccessConsentsConsentIDParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccountAccessConsentsConsentIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
