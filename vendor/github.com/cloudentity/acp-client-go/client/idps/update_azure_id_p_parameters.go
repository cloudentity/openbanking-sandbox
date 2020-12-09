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

// NewUpdateAzureIDPParams creates a new UpdateAzureIDPParams object
// with the default values initialized.
func NewUpdateAzureIDPParams() *UpdateAzureIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateAzureIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAzureIDPParamsWithTimeout creates a new UpdateAzureIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAzureIDPParamsWithTimeout(timeout time.Duration) *UpdateAzureIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateAzureIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewUpdateAzureIDPParamsWithContext creates a new UpdateAzureIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAzureIDPParamsWithContext(ctx context.Context) *UpdateAzureIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateAzureIDPParams{
		Aid: aidDefault,
		Iid: iidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewUpdateAzureIDPParamsWithHTTPClient creates a new UpdateAzureIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAzureIDPParamsWithHTTPClient(client *http.Client) *UpdateAzureIDPParams {
	var (
		aidDefault = string("default")
		iidDefault = string("default")
		tidDefault = string("default")
	)
	return &UpdateAzureIDPParams{
		Aid:        aidDefault,
		Iid:        iidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*UpdateAzureIDPParams contains all the parameters to send to the API endpoint
for the update azure ID p operation typically these are written to a http.Request
*/
type UpdateAzureIDPParams struct {

	/*AzureIDP
	  AzureIDP

	*/
	AzureIDP *models.AzureIDP
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

// WithTimeout adds the timeout to the update azure ID p params
func (o *UpdateAzureIDPParams) WithTimeout(timeout time.Duration) *UpdateAzureIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update azure ID p params
func (o *UpdateAzureIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update azure ID p params
func (o *UpdateAzureIDPParams) WithContext(ctx context.Context) *UpdateAzureIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update azure ID p params
func (o *UpdateAzureIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update azure ID p params
func (o *UpdateAzureIDPParams) WithHTTPClient(client *http.Client) *UpdateAzureIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update azure ID p params
func (o *UpdateAzureIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzureIDP adds the azureIDP to the update azure ID p params
func (o *UpdateAzureIDPParams) WithAzureIDP(azureIDP *models.AzureIDP) *UpdateAzureIDPParams {
	o.SetAzureIDP(azureIDP)
	return o
}

// SetAzureIDP adds the azureIdP to the update azure ID p params
func (o *UpdateAzureIDPParams) SetAzureIDP(azureIDP *models.AzureIDP) {
	o.AzureIDP = azureIDP
}

// WithAid adds the aid to the update azure ID p params
func (o *UpdateAzureIDPParams) WithAid(aid string) *UpdateAzureIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the update azure ID p params
func (o *UpdateAzureIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithIid adds the iid to the update azure ID p params
func (o *UpdateAzureIDPParams) WithIid(iid string) *UpdateAzureIDPParams {
	o.SetIid(iid)
	return o
}

// SetIid adds the iid to the update azure ID p params
func (o *UpdateAzureIDPParams) SetIid(iid string) {
	o.Iid = iid
}

// WithTid adds the tid to the update azure ID p params
func (o *UpdateAzureIDPParams) WithTid(tid string) *UpdateAzureIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the update azure ID p params
func (o *UpdateAzureIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAzureIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AzureIDP != nil {
		if err := r.SetBodyParam(o.AzureIDP); err != nil {
			return err
		}
	}

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