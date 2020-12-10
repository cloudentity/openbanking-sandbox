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

	"github.com/cloudentity/acp-client-go/models"
)

// NewCreateOIDCIDPParams creates a new CreateOIDCIDPParams object
// with the default values initialized.
func NewCreateOIDCIDPParams() *CreateOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateOIDCIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOIDCIDPParamsWithTimeout creates a new CreateOIDCIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOIDCIDPParamsWithTimeout(timeout time.Duration) *CreateOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateOIDCIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewCreateOIDCIDPParamsWithContext creates a new CreateOIDCIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOIDCIDPParamsWithContext(ctx context.Context) *CreateOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateOIDCIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewCreateOIDCIDPParamsWithHTTPClient creates a new CreateOIDCIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOIDCIDPParamsWithHTTPClient(client *http.Client) *CreateOIDCIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateOIDCIDPParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*CreateOIDCIDPParams contains all the parameters to send to the API endpoint
for the create o ID c ID p operation typically these are written to a http.Request
*/
type CreateOIDCIDPParams struct {

	/*OIDCIDP
	  OIDCIDP

	*/
	OIDCIDP *models.OIDCIDP
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

// WithTimeout adds the timeout to the create o ID c ID p params
func (o *CreateOIDCIDPParams) WithTimeout(timeout time.Duration) *CreateOIDCIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create o ID c ID p params
func (o *CreateOIDCIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create o ID c ID p params
func (o *CreateOIDCIDPParams) WithContext(ctx context.Context) *CreateOIDCIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create o ID c ID p params
func (o *CreateOIDCIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create o ID c ID p params
func (o *CreateOIDCIDPParams) WithHTTPClient(client *http.Client) *CreateOIDCIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create o ID c ID p params
func (o *CreateOIDCIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOIDCIDP adds the oIDCIDP to the create o ID c ID p params
func (o *CreateOIDCIDPParams) WithOIDCIDP(oIDCIDP *models.OIDCIDP) *CreateOIDCIDPParams {
	o.SetOIDCIDP(oIDCIDP)
	return o
}

// SetOIDCIDP adds the oIdCIdP to the create o ID c ID p params
func (o *CreateOIDCIDPParams) SetOIDCIDP(oIDCIDP *models.OIDCIDP) {
	o.OIDCIDP = oIDCIDP
}

// WithAid adds the aid to the create o ID c ID p params
func (o *CreateOIDCIDPParams) WithAid(aid string) *CreateOIDCIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the create o ID c ID p params
func (o *CreateOIDCIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the create o ID c ID p params
func (o *CreateOIDCIDPParams) WithTid(tid string) *CreateOIDCIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the create o ID c ID p params
func (o *CreateOIDCIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOIDCIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OIDCIDP != nil {
		if err := r.SetBodyParam(o.OIDCIDP); err != nil {
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
