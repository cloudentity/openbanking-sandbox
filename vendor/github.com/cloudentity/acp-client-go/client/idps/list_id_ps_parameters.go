// Code generated by go-swagger; DO NOT EDIT.

package idps

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

// NewListIDPsParams creates a new ListIDPsParams object
// with the default values initialized.
func NewListIDPsParams() *ListIDPsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListIDPsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListIDPsParamsWithTimeout creates a new ListIDPsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIDPsParamsWithTimeout(timeout time.Duration) *ListIDPsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListIDPsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewListIDPsParamsWithContext creates a new ListIDPsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIDPsParamsWithContext(ctx context.Context) *ListIDPsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListIDPsParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewListIDPsParamsWithHTTPClient creates a new ListIDPsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIDPsParamsWithHTTPClient(client *http.Client) *ListIDPsParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &ListIDPsParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*ListIDPsParams contains all the parameters to send to the API endpoint
for the list ID ps operation typically these are written to a http.Request
*/
type ListIDPsParams struct {

	/*Aid
	  Authorization server id

	*/
	Aid string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list ID ps params
func (o *ListIDPsParams) WithTimeout(timeout time.Duration) *ListIDPsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ID ps params
func (o *ListIDPsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ID ps params
func (o *ListIDPsParams) WithContext(ctx context.Context) *ListIDPsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ID ps params
func (o *ListIDPsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ID ps params
func (o *ListIDPsParams) WithHTTPClient(client *http.Client) *ListIDPsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ID ps params
func (o *ListIDPsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the list ID ps params
func (o *ListIDPsParams) WithAid(aid string) *ListIDPsParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the list ID ps params
func (o *ListIDPsParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the list ID ps params
func (o *ListIDPsParams) WithTid(tid string) *ListIDPsParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the list ID ps params
func (o *ListIDPsParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *ListIDPsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
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
