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

// NewListConsentsByClientIDParams creates a new ListConsentsByClientIDParams object
// with the default values initialized.
func NewListConsentsByClientIDParams() *ListConsentsByClientIDParams {
	var (
		tidDefault = string("default")
	)
	return &ListConsentsByClientIDParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListConsentsByClientIDParamsWithTimeout creates a new ListConsentsByClientIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListConsentsByClientIDParamsWithTimeout(timeout time.Duration) *ListConsentsByClientIDParams {
	var (
		tidDefault = string("default")
	)
	return &ListConsentsByClientIDParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewListConsentsByClientIDParamsWithContext creates a new ListConsentsByClientIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewListConsentsByClientIDParamsWithContext(ctx context.Context) *ListConsentsByClientIDParams {
	var (
		tidDefault = string("default")
	)
	return &ListConsentsByClientIDParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewListConsentsByClientIDParamsWithHTTPClient creates a new ListConsentsByClientIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListConsentsByClientIDParamsWithHTTPClient(client *http.Client) *ListConsentsByClientIDParams {
	var (
		tidDefault = string("default")
	)
	return &ListConsentsByClientIDParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*ListConsentsByClientIDParams contains all the parameters to send to the API endpoint
for the list consents by client ID operation typically these are written to a http.Request
*/
type ListConsentsByClientIDParams struct {

	/*ClientID*/
	ClientID string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list consents by client ID params
func (o *ListConsentsByClientIDParams) WithTimeout(timeout time.Duration) *ListConsentsByClientIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list consents by client ID params
func (o *ListConsentsByClientIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list consents by client ID params
func (o *ListConsentsByClientIDParams) WithContext(ctx context.Context) *ListConsentsByClientIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list consents by client ID params
func (o *ListConsentsByClientIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list consents by client ID params
func (o *ListConsentsByClientIDParams) WithHTTPClient(client *http.Client) *ListConsentsByClientIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list consents by client ID params
func (o *ListConsentsByClientIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientID adds the clientID to the list consents by client ID params
func (o *ListConsentsByClientIDParams) WithClientID(clientID string) *ListConsentsByClientIDParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the list consents by client ID params
func (o *ListConsentsByClientIDParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WithTid adds the tid to the list consents by client ID params
func (o *ListConsentsByClientIDParams) WithTid(tid string) *ListConsentsByClientIDParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the list consents by client ID params
func (o *ListConsentsByClientIDParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *ListConsentsByClientIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientID
	if err := r.SetPathParam("clientID", o.ClientID); err != nil {
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
