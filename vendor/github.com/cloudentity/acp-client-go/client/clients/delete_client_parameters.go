// Code generated by go-swagger; DO NOT EDIT.

package clients

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

// NewDeleteClientParams creates a new DeleteClientParams object
// with the default values initialized.
func NewDeleteClientParams() *DeleteClientParams {
	var (
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DeleteClientParams{
		Cid: cidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClientParamsWithTimeout creates a new DeleteClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClientParamsWithTimeout(timeout time.Duration) *DeleteClientParams {
	var (
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DeleteClientParams{
		Cid: cidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewDeleteClientParamsWithContext creates a new DeleteClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClientParamsWithContext(ctx context.Context) *DeleteClientParams {
	var (
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DeleteClientParams{
		Cid: cidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewDeleteClientParamsWithHTTPClient creates a new DeleteClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClientParamsWithHTTPClient(client *http.Client) *DeleteClientParams {
	var (
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DeleteClientParams{
		Cid:        cidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*DeleteClientParams contains all the parameters to send to the API endpoint
for the delete client operation typically these are written to a http.Request
*/
type DeleteClientParams struct {

	/*Cid
	  Client id

	*/
	Cid string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete client params
func (o *DeleteClientParams) WithTimeout(timeout time.Duration) *DeleteClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete client params
func (o *DeleteClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete client params
func (o *DeleteClientParams) WithContext(ctx context.Context) *DeleteClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete client params
func (o *DeleteClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete client params
func (o *DeleteClientParams) WithHTTPClient(client *http.Client) *DeleteClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete client params
func (o *DeleteClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the delete client params
func (o *DeleteClientParams) WithCid(cid string) *DeleteClientParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the delete client params
func (o *DeleteClientParams) SetCid(cid string) {
	o.Cid = cid
}

// WithTid adds the tid to the delete client params
func (o *DeleteClientParams) WithTid(tid string) *DeleteClientParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the delete client params
func (o *DeleteClientParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cid
	if err := r.SetPathParam("cid", o.Cid); err != nil {
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
