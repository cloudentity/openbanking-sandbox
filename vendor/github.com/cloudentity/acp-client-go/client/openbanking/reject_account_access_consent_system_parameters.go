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

	"github.com/cloudentity/acp-client-go/models"
)

// NewRejectAccountAccessConsentSystemParams creates a new RejectAccountAccessConsentSystemParams object
// with the default values initialized.
func NewRejectAccountAccessConsentSystemParams() *RejectAccountAccessConsentSystemParams {
	var (
		tidDefault = string("default")
	)
	return &RejectAccountAccessConsentSystemParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewRejectAccountAccessConsentSystemParamsWithTimeout creates a new RejectAccountAccessConsentSystemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRejectAccountAccessConsentSystemParamsWithTimeout(timeout time.Duration) *RejectAccountAccessConsentSystemParams {
	var (
		tidDefault = string("default")
	)
	return &RejectAccountAccessConsentSystemParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewRejectAccountAccessConsentSystemParamsWithContext creates a new RejectAccountAccessConsentSystemParams object
// with the default values initialized, and the ability to set a context for a request
func NewRejectAccountAccessConsentSystemParamsWithContext(ctx context.Context) *RejectAccountAccessConsentSystemParams {
	var (
		tidDefault = string("default")
	)
	return &RejectAccountAccessConsentSystemParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewRejectAccountAccessConsentSystemParamsWithHTTPClient creates a new RejectAccountAccessConsentSystemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRejectAccountAccessConsentSystemParamsWithHTTPClient(client *http.Client) *RejectAccountAccessConsentSystemParams {
	var (
		tidDefault = string("default")
	)
	return &RejectAccountAccessConsentSystemParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*RejectAccountAccessConsentSystemParams contains all the parameters to send to the API endpoint
for the reject account access consent system operation typically these are written to a http.Request
*/
type RejectAccountAccessConsentSystemParams struct {

	/*RejectAccountAccessConsent*/
	RejectAccountAccessConsent *models.RejectAccountAccessConsentRequest
	/*Login*/
	LoginID string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) WithTimeout(timeout time.Duration) *RejectAccountAccessConsentSystemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) WithContext(ctx context.Context) *RejectAccountAccessConsentSystemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) WithHTTPClient(client *http.Client) *RejectAccountAccessConsentSystemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRejectAccountAccessConsent adds the rejectAccountAccessConsent to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) WithRejectAccountAccessConsent(rejectAccountAccessConsent *models.RejectAccountAccessConsentRequest) *RejectAccountAccessConsentSystemParams {
	o.SetRejectAccountAccessConsent(rejectAccountAccessConsent)
	return o
}

// SetRejectAccountAccessConsent adds the rejectAccountAccessConsent to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) SetRejectAccountAccessConsent(rejectAccountAccessConsent *models.RejectAccountAccessConsentRequest) {
	o.RejectAccountAccessConsent = rejectAccountAccessConsent
}

// WithLoginID adds the login to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) WithLoginID(login string) *RejectAccountAccessConsentSystemParams {
	o.SetLoginID(login)
	return o
}

// SetLoginID adds the login to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) SetLoginID(login string) {
	o.LoginID = login
}

// WithTid adds the tid to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) WithTid(tid string) *RejectAccountAccessConsentSystemParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the reject account access consent system params
func (o *RejectAccountAccessConsentSystemParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *RejectAccountAccessConsentSystemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RejectAccountAccessConsent != nil {
		if err := r.SetBodyParam(o.RejectAccountAccessConsent); err != nil {
			return err
		}
	}

	// path param login
	if err := r.SetPathParam("login", o.LoginID); err != nil {
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
