// Code generated by go-swagger; DO NOT EDIT.

package openbanking

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

// NewRevokeOpenbankingConsentParams creates a new RevokeOpenbankingConsentParams object
// with the default values initialized.
func NewRevokeOpenbankingConsentParams() *RevokeOpenbankingConsentParams {
	var (
		tidDefault = string("default")
	)
	return &RevokeOpenbankingConsentParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeOpenbankingConsentParamsWithTimeout creates a new RevokeOpenbankingConsentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevokeOpenbankingConsentParamsWithTimeout(timeout time.Duration) *RevokeOpenbankingConsentParams {
	var (
		tidDefault = string("default")
	)
	return &RevokeOpenbankingConsentParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewRevokeOpenbankingConsentParamsWithContext creates a new RevokeOpenbankingConsentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevokeOpenbankingConsentParamsWithContext(ctx context.Context) *RevokeOpenbankingConsentParams {
	var (
		tidDefault = string("default")
	)
	return &RevokeOpenbankingConsentParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewRevokeOpenbankingConsentParamsWithHTTPClient creates a new RevokeOpenbankingConsentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevokeOpenbankingConsentParamsWithHTTPClient(client *http.Client) *RevokeOpenbankingConsentParams {
	var (
		tidDefault = string("default")
	)
	return &RevokeOpenbankingConsentParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*RevokeOpenbankingConsentParams contains all the parameters to send to the API endpoint
for the revoke openbanking consent operation typically these are written to a http.Request
*/
type RevokeOpenbankingConsentParams struct {

	/*ConsentID*/
	ConsentID string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) WithTimeout(timeout time.Duration) *RevokeOpenbankingConsentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) WithContext(ctx context.Context) *RevokeOpenbankingConsentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) WithHTTPClient(client *http.Client) *RevokeOpenbankingConsentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsentID adds the consentID to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) WithConsentID(consentID string) *RevokeOpenbankingConsentParams {
	o.SetConsentID(consentID)
	return o
}

// SetConsentID adds the consentId to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) SetConsentID(consentID string) {
	o.ConsentID = consentID
}

// WithTid adds the tid to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) WithTid(tid string) *RevokeOpenbankingConsentParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the revoke openbanking consent params
func (o *RevokeOpenbankingConsentParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeOpenbankingConsentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param consentID
	if err := r.SetPathParam("consentID", o.ConsentID); err != nil {
		return err
	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}