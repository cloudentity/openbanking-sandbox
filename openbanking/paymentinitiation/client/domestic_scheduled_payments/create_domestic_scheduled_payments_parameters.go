// Code generated by go-swagger; DO NOT EDIT.

package domestic_scheduled_payments

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

// NewCreateDomesticScheduledPaymentsParams creates a new CreateDomesticScheduledPaymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDomesticScheduledPaymentsParams() *CreateDomesticScheduledPaymentsParams {
	return &CreateDomesticScheduledPaymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDomesticScheduledPaymentsParamsWithTimeout creates a new CreateDomesticScheduledPaymentsParams object
// with the ability to set a timeout on a request.
func NewCreateDomesticScheduledPaymentsParamsWithTimeout(timeout time.Duration) *CreateDomesticScheduledPaymentsParams {
	return &CreateDomesticScheduledPaymentsParams{
		timeout: timeout,
	}
}

// NewCreateDomesticScheduledPaymentsParamsWithContext creates a new CreateDomesticScheduledPaymentsParams object
// with the ability to set a context for a request.
func NewCreateDomesticScheduledPaymentsParamsWithContext(ctx context.Context) *CreateDomesticScheduledPaymentsParams {
	return &CreateDomesticScheduledPaymentsParams{
		Context: ctx,
	}
}

// NewCreateDomesticScheduledPaymentsParamsWithHTTPClient creates a new CreateDomesticScheduledPaymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDomesticScheduledPaymentsParamsWithHTTPClient(client *http.Client) *CreateDomesticScheduledPaymentsParams {
	return &CreateDomesticScheduledPaymentsParams{
		HTTPClient: client,
	}
}

/* CreateDomesticScheduledPaymentsParams contains all the parameters to send to the API endpoint
   for the create domestic scheduled payments operation.

   Typically these are written to a http.Request.
*/
type CreateDomesticScheduledPaymentsParams struct {

	/* Authorization.

	   An Authorisation Token as per https://tools.ietf.org/html/rfc6750
	*/
	Authorization string

	/* OBWriteDomesticScheduled2Param.

	   Default
	*/
	OBWriteDomesticScheduled2Param *models.OBWriteDomesticScheduled2

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

// WithDefaults hydrates default values in the create domestic scheduled payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomesticScheduledPaymentsParams) WithDefaults() *CreateDomesticScheduledPaymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create domestic scheduled payments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDomesticScheduledPaymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithTimeout(timeout time.Duration) *CreateDomesticScheduledPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithContext(ctx context.Context) *CreateDomesticScheduledPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithHTTPClient(client *http.Client) *CreateDomesticScheduledPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithAuthorization(authorization string) *CreateDomesticScheduledPaymentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithOBWriteDomesticScheduled2Param adds the oBWriteDomesticScheduled2Param to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithOBWriteDomesticScheduled2Param(oBWriteDomesticScheduled2Param *models.OBWriteDomesticScheduled2) *CreateDomesticScheduledPaymentsParams {
	o.SetOBWriteDomesticScheduled2Param(oBWriteDomesticScheduled2Param)
	return o
}

// SetOBWriteDomesticScheduled2Param adds the oBWriteDomesticScheduled2Param to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetOBWriteDomesticScheduled2Param(oBWriteDomesticScheduled2Param *models.OBWriteDomesticScheduled2) {
	o.OBWriteDomesticScheduled2Param = oBWriteDomesticScheduled2Param
}

// WithXCustomerUserAgent adds the xCustomerUserAgent to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithXCustomerUserAgent(xCustomerUserAgent *string) *CreateDomesticScheduledPaymentsParams {
	o.SetXCustomerUserAgent(xCustomerUserAgent)
	return o
}

// SetXCustomerUserAgent adds the xCustomerUserAgent to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetXCustomerUserAgent(xCustomerUserAgent *string) {
	o.XCustomerUserAgent = xCustomerUserAgent
}

// WithXFapiAuthDate adds the xFapiAuthDate to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithXFapiAuthDate(xFapiAuthDate *string) *CreateDomesticScheduledPaymentsParams {
	o.SetXFapiAuthDate(xFapiAuthDate)
	return o
}

// SetXFapiAuthDate adds the xFapiAuthDate to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetXFapiAuthDate(xFapiAuthDate *string) {
	o.XFapiAuthDate = xFapiAuthDate
}

// WithXFapiCustomerIPAddress adds the xFapiCustomerIPAddress to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) *CreateDomesticScheduledPaymentsParams {
	o.SetXFapiCustomerIPAddress(xFapiCustomerIPAddress)
	return o
}

// SetXFapiCustomerIPAddress adds the xFapiCustomerIpAddress to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetXFapiCustomerIPAddress(xFapiCustomerIPAddress *string) {
	o.XFapiCustomerIPAddress = xFapiCustomerIPAddress
}

// WithXFapiInteractionID adds the xFapiInteractionID to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithXFapiInteractionID(xFapiInteractionID *string) *CreateDomesticScheduledPaymentsParams {
	o.SetXFapiInteractionID(xFapiInteractionID)
	return o
}

// SetXFapiInteractionID adds the xFapiInteractionId to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetXFapiInteractionID(xFapiInteractionID *string) {
	o.XFapiInteractionID = xFapiInteractionID
}

// WithXIdempotencyKey adds the xIdempotencyKey to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithXIdempotencyKey(xIdempotencyKey string) *CreateDomesticScheduledPaymentsParams {
	o.SetXIdempotencyKey(xIdempotencyKey)
	return o
}

// SetXIdempotencyKey adds the xIdempotencyKey to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetXIdempotencyKey(xIdempotencyKey string) {
	o.XIdempotencyKey = xIdempotencyKey
}

// WithXJwsSignature adds the xJwsSignature to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) WithXJwsSignature(xJwsSignature string) *CreateDomesticScheduledPaymentsParams {
	o.SetXJwsSignature(xJwsSignature)
	return o
}

// SetXJwsSignature adds the xJwsSignature to the create domestic scheduled payments params
func (o *CreateDomesticScheduledPaymentsParams) SetXJwsSignature(xJwsSignature string) {
	o.XJwsSignature = xJwsSignature
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDomesticScheduledPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}
	if o.OBWriteDomesticScheduled2Param != nil {
		if err := r.SetBodyParam(o.OBWriteDomesticScheduled2Param); err != nil {
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
