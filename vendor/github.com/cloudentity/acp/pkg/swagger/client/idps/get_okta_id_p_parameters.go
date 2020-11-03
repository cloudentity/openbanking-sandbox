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

// NewGetOktaIDPParams creates a new GetOktaIDPParams object
// with the default values initialized.
func NewGetOktaIDPParams() *GetOktaIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOktaIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOktaIDPParamsWithTimeout creates a new GetOktaIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOktaIDPParamsWithTimeout(timeout time.Duration) *GetOktaIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOktaIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewGetOktaIDPParamsWithContext creates a new GetOktaIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOktaIDPParamsWithContext(ctx context.Context) *GetOktaIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOktaIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewGetOktaIDPParamsWithHTTPClient creates a new GetOktaIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOktaIDPParamsWithHTTPClient(client *http.Client) *GetOktaIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &GetOktaIDPParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*GetOktaIDPParams contains all the parameters to send to the API endpoint
for the get okta ID p operation typically these are written to a http.Request
*/
type GetOktaIDPParams struct {

	/*Aid
	  Authorization server id

	*/
	Aid string
	/*Iid
	  IDP id

	*/
	Iid string
	/*Tid
	  Tenant id

	*/
	Tid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get okta ID p params
func (o *GetOktaIDPParams) WithTimeout(timeout time.Duration) *GetOktaIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get okta ID p params
func (o *GetOktaIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get okta ID p params
func (o *GetOktaIDPParams) WithContext(ctx context.Context) *GetOktaIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get okta ID p params
func (o *GetOktaIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get okta ID p params
func (o *GetOktaIDPParams) WithHTTPClient(client *http.Client) *GetOktaIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get okta ID p params
func (o *GetOktaIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAid adds the aid to the get okta ID p params
func (o *GetOktaIDPParams) WithAid(aid string) *GetOktaIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the get okta ID p params
func (o *GetOktaIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the get okta ID p params
func (o *GetOktaIDPParams) WithIid(iid string) *GetOktaIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the get okta ID p params
func (o *GetOktaIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the get okta ID p params
func (o *GetOktaIDPParams) WithTid(tid string) *GetOktaIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the get okta ID p params
func (o *GetOktaIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *GetOktaIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aid
	if err := r.SetPathParam("aid", o.Aid); err != nil {
		return err
	}

	// path param iid
	if err := r.SetPathParam("iid", o.Iid); err != nil {
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