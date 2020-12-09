// Code generated by go-swagger; DO NOT EDIT.

package oauth2

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

// NewDynamicClientRegistrationParams creates a new DynamicClientRegistrationParams object
// with the default values initialized.
func NewDynamicClientRegistrationParams() *DynamicClientRegistrationParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDynamicClientRegistrationParamsWithTimeout creates a new DynamicClientRegistrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDynamicClientRegistrationParamsWithTimeout(timeout time.Duration) *DynamicClientRegistrationParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewDynamicClientRegistrationParamsWithContext creates a new DynamicClientRegistrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDynamicClientRegistrationParamsWithContext(ctx context.Context) *DynamicClientRegistrationParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewDynamicClientRegistrationParamsWithHTTPClient creates a new DynamicClientRegistrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDynamicClientRegistrationParamsWithHTTPClient(client *http.Client) *DynamicClientRegistrationParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*DynamicClientRegistrationParams contains all the parameters to send to the API endpoint
for the dynamic client registration operation typically these are written to a http.Request
*/
type DynamicClientRegistrationParams struct {

	/*Client*/
	Client *models.DynamicClientRegistrationRequest
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

// WithTimeout adds the timeout to the dynamic client registration params
func (o *DynamicClientRegistrationParams) WithTimeout(timeout time.Duration) *DynamicClientRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dynamic client registration params
func (o *DynamicClientRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dynamic client registration params
func (o *DynamicClientRegistrationParams) WithContext(ctx context.Context) *DynamicClientRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dynamic client registration params
func (o *DynamicClientRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dynamic client registration params
func (o *DynamicClientRegistrationParams) WithHTTPClient(client *http.Client) *DynamicClientRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dynamic client registration params
func (o *DynamicClientRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the dynamic client registration params
func (o *DynamicClientRegistrationParams) WithClient(client *models.DynamicClientRegistrationRequest) *DynamicClientRegistrationParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the dynamic client registration params
func (o *DynamicClientRegistrationParams) SetClient(client *models.DynamicClientRegistrationRequest) {
	o.Client = client
}

// WithAid adds the aid to the dynamic client registration params
func (o *DynamicClientRegistrationParams) WithAid(aid string) *DynamicClientRegistrationParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the dynamic client registration params
func (o *DynamicClientRegistrationParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the dynamic client registration params
func (o *DynamicClientRegistrationParams) WithTid(tid string) *DynamicClientRegistrationParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the dynamic client registration params
func (o *DynamicClientRegistrationParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *DynamicClientRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Client != nil {
		if err := r.SetBodyParam(o.Client); err != nil {
			return err
		}
	}

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
