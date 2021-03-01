// Code generated by go-swagger; DO NOT EDIT.

package international_scheduled_payments

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

	"github.com/cloudentity/openbanking-sandbox/openbanking/paymentinitiation/models"
)

// NewCreateInternationalScheduledPaymentConsentsParams creates a new CreateInternationalScheduledPaymentConsentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateInternationalScheduledPaymentConsentsParams() *CreateInternationalScheduledPaymentConsentsParams {
	return &CreateInternationalScheduledPaymentConsentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInternationalScheduledPaymentConsentsParamsWithTimeout creates a new CreateInternationalScheduledPaymentConsentsParams object
// with the ability to set a timeout on a request.
func NewCreateInternationalScheduledPaymentConsentsParamsWithTimeout(timeout time.Duration) *CreateInternationalScheduledPaymentConsentsParams {
	return &CreateInternationalScheduledPaymentConsentsParams{
		timeout: timeout,
	}
}

// NewCreateInternationalScheduledPaymentConsentsParamsWithContext creates a new CreateInternationalScheduledPaymentConsentsParams object
// with the ability to set a context for a request.
func NewCreateInternationalScheduledPaymentConsentsParamsWithContext(ctx context.Context) *CreateInternationalScheduledPaymentConsentsParams {
	return &CreateInternationalScheduledPaymentConsentsParams{
		Context: ctx,
	}
}

// NewCreateInternationalScheduledPaymentConsentsParamsWithHTTPClient creates a new CreateInternationalScheduledPaymentConsentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateInternationalScheduledPaymentConsentsParamsWithHTTPClient(client *http.Client) *CreateInternationalScheduledPaymentConsentsParams {
	return &CreateInternationalScheduledPaymentConsentsParams{
		HTTPClient: client,
	}
}

/* CreateInternationalScheduledPaymentConsentsParams contains all the parameters to send to the API endpoint
   for the create international scheduled payment consents operation.

   Typically these are written to a http.Request.
*/
type CreateInternationalScheduledPaymentConsentsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* OBWriteInternationalScheduledConsent5Param.

	   Default
	*/
	OBWriteInternationalScheduledConsent5Param *models.OBWriteInternationalScheduledConsent5

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

	/* XIdempotencyKey.

	     Every request will be processed only once per x-idempotency-key.  The
	Idempotency Key will be valid for 24 hours.

	*/
	XIdempotencyKey string

	/* XJwsSignature.

	   A detached JWS signature of the body of the payload.
	*/
	XJwsSignature string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create international scheduled payment consents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInternationalScheduledPaymentConsentsParams) WithDefaults() *CreateInternationalScheduledPaymentConsentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create international scheduled payment consents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateInternationalScheduledPaymentConsentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithTimeout(timeout time.Duration) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithContext(ctx context.Context) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithHTTPClient(client *http.Client) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithAuthorization(authorization string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithOBWriteInternationalScheduledConsent5Param adds the oBWriteInternationalScheduledConsent5Param to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithOBWriteInternationalScheduledConsent5Param(oBWriteInternationalScheduledConsent5Param *models.OBWriteInternationalScheduledConsent5) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetOBWriteInternationalScheduledConsent5Param(oBWriteInternationalScheduledConsent5Param)
	return o
}

// SetOBWriteInternationalScheduledConsent5Param adds the oBWriteInternationalScheduledConsent5Param to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetOBWriteInternationalScheduledConsent5Param(oBWriteInternationalScheduledConsent5Param *models.OBWriteInternationalScheduledConsent5) {
	o.OBWriteInternationalScheduledConsent5Param = oBWriteInternationalScheduledConsent5Param
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXIdempotencyKey adds the xIdempotencyKey to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithXIdempotencyKey(xIdempotencyKey string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetXIdempotencyKey(xIdempotencyKey)
	return o
}

// SetXIdempotencyKey adds the xIdempotencyKey to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetXIdempotencyKey(xIdempotencyKey string) {
	o.XIdempotencyKey = xIdempotencyKey
}

// WithXJwsSignature adds the xJwsSignature to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) WithXJwsSignature(xJwsSignature string) *CreateInternationalScheduledPaymentConsentsParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the create international scheduled payment consents params
func (o *CreateInternationalScheduledPaymentConsentsParams) SetXJwsSignature(xJwsSignature string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInternationalScheduledPaymentConsentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}
	if o.OBWriteInternationalScheduledConsent5Param != nil {
		if err := r.SetBodyParam(o.OBWriteInternationalScheduledConsent5Param); err != nil {
			return err
		}
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

	// header param x-idempotency-key
	if err := r.SetHeaderParam("x-idempotency-key", o.XIdempotencyKey); err != nil {
		return err
	}

	// header param x-jws-signature
	if err := r.SetHeaderParam("x-jws-signature", o.XJwsSignature); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
