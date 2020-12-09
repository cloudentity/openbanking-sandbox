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

// NewDynamicClientRegistrationUpdateClientParams creates a new DynamicClientRegistrationUpdateClientParams object
// with the default values initialized.
func NewDynamicClientRegistrationUpdateClientParams() *DynamicClientRegistrationUpdateClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationUpdateClientParams{
		Aid: aidDefault,
		Cid: cidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDynamicClientRegistrationUpdateClientParamsWithTimeout creates a new DynamicClientRegistrationUpdateClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDynamicClientRegistrationUpdateClientParamsWithTimeout(timeout time.Duration) *DynamicClientRegistrationUpdateClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationUpdateClientParams{
		Aid: aidDefault,
		Cid: cidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewDynamicClientRegistrationUpdateClientParamsWithContext creates a new DynamicClientRegistrationUpdateClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewDynamicClientRegistrationUpdateClientParamsWithContext(ctx context.Context) *DynamicClientRegistrationUpdateClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationUpdateClientParams{
		Aid: aidDefault,
		Cid: cidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewDynamicClientRegistrationUpdateClientParamsWithHTTPClient creates a new DynamicClientRegistrationUpdateClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDynamicClientRegistrationUpdateClientParamsWithHTTPClient(client *http.Client) *DynamicClientRegistrationUpdateClientParams {
	var (
		aidDefault = string("default")
		cidDefault = string("default")
		tidDefault = string("default")
	)
	return &DynamicClientRegistrationUpdateClientParams{
		Aid:        aidDefault,
		Cid:        cidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*DynamicClientRegistrationUpdateClientParams contains all the parameters to send to the API endpoint
for the dynamic client registration update client operation typically these are written to a http.Request
*/
type DynamicClientRegistrationUpdateClientParams struct {

	/*Client*/
	Client *models.DynamicClientRegistrationRequest
	/*Aid
	  Authorization server id

	*/
	Aid string
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

// WithTimeout adds the timeout to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithTimeout(timeout time.Duration) *DynamicClientRegistrationUpdateClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithContext(ctx context.Context) *DynamicClientRegistrationUpdateClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithHTTPClient(client *http.Client) *DynamicClientRegistrationUpdateClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClient adds the client to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithClient(client *models.DynamicClientRegistrationRequest) *DynamicClientRegistrationUpdateClientParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetClient(client *models.DynamicClientRegistrationRequest) {
	o.Client = client
}

// WithAid adds the aid to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithAid(aid string) *DynamicClientRegistrationUpdateClientParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetAid(aid string) {
	o.Aid = aid
}

// WithCid adds the cid to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithCid(cid string) *DynamicClientRegistrationUpdateClientParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetCid(cid string) {
	o.Cid = cid
}

// WithTid adds the tid to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) WithTid(tid string) *DynamicClientRegistrationUpdateClientParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the dynamic client registration update client params
func (o *DynamicClientRegistrationUpdateClientParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *DynamicClientRegistrationUpdateClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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