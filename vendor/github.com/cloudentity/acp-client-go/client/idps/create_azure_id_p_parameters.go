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

// NewCreateAzureIDPParams creates a new CreateAzureIDPParams object
// with the default values initialized.
func NewCreateAzureIDPParams() *CreateAzureIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateAzureIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAzureIDPParamsWithTimeout creates a new CreateAzureIDPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAzureIDPParamsWithTimeout(timeout time.Duration) *CreateAzureIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateAzureIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewCreateAzureIDPParamsWithContext creates a new CreateAzureIDPParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAzureIDPParamsWithContext(ctx context.Context) *CreateAzureIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateAzureIDPParams{
		Aid: aidDefault,
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewCreateAzureIDPParamsWithHTTPClient creates a new CreateAzureIDPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAzureIDPParamsWithHTTPClient(client *http.Client) *CreateAzureIDPParams {
	var (
		aidDefault = string("default")
		tidDefault = string("default")
	)
	return &CreateAzureIDPParams{
		Aid:        aidDefault,
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*CreateAzureIDPParams contains all the parameters to send to the API endpoint
for the create azure ID p operation typically these are written to a http.Request
*/
type CreateAzureIDPParams struct {

	/*AzureIDP
	  AzureIDP

	*/
	AzureIDP *models.AzureIDP
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

// WithTimeout adds the timeout to the create azure ID p params
func (o *CreateAzureIDPParams) WithTimeout(timeout time.Duration) *CreateAzureIDPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create azure ID p params
func (o *CreateAzureIDPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create azure ID p params
func (o *CreateAzureIDPParams) WithContext(ctx context.Context) *CreateAzureIDPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create azure ID p params
func (o *CreateAzureIDPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create azure ID p params
func (o *CreateAzureIDPParams) WithHTTPClient(client *http.Client) *CreateAzureIDPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create azure ID p params
func (o *CreateAzureIDPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAzureIDP adds the azureIDP to the create azure ID p params
func (o *CreateAzureIDPParams) WithAzureIDP(azureIDP *models.AzureIDP) *CreateAzureIDPParams {
	o.SetAzureIDP(azureIDP)
	return o
}

// SetAzureIDP adds the azureIdP to the create azure ID p params
func (o *CreateAzureIDPParams) SetAzureIDP(azureIDP *models.AzureIDP) {
	o.AzureIDP = azureIDP
}

// WithAid adds the aid to the create azure ID p params
func (o *CreateAzureIDPParams) WithAid(aid string) *CreateAzureIDPParams {
	o.SetAid(aid)
	return o
}

// SetAid adds the aid to the create azure ID p params
func (o *CreateAzureIDPParams) SetAid(aid string) {
	o.Aid = aid
}

// WithTid adds the tid to the create azure ID p params
func (o *CreateAzureIDPParams) WithTid(tid string) *CreateAzureIDPParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the create azure ID p params
func (o *CreateAzureIDPParams) SetTid(tid string) {
	o.Tid = tid
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAzureIDPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}