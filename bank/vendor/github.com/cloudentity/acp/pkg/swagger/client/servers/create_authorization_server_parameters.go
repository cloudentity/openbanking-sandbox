// Code generated by go-swagger; DO NOT EDIT.

package servers

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
	"github.com/go-openapi/swag"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// NewCreateAuthorizationServerParams creates a new CreateAuthorizationServerParams object
// with the default values initialized.
func NewCreateAuthorizationServerParams() *CreateAuthorizationServerParams {
	var (
		tidDefault = string("default")
	)
	return &CreateAuthorizationServerParams{
		Tid: tidDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthorizationServerParamsWithTimeout creates a new CreateAuthorizationServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAuthorizationServerParamsWithTimeout(timeout time.Duration) *CreateAuthorizationServerParams {
	var (
		tidDefault = string("default")
	)
	return &CreateAuthorizationServerParams{
		Tid: tidDefault,

		timeout: timeout,
	}
}

// NewCreateAuthorizationServerParamsWithContext creates a new CreateAuthorizationServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAuthorizationServerParamsWithContext(ctx context.Context) *CreateAuthorizationServerParams {
	var (
		tidDefault = string("default")
	)
	return &CreateAuthorizationServerParams{
		Tid: tidDefault,

		Context: ctx,
	}
}

// NewCreateAuthorizationServerParamsWithHTTPClient creates a new CreateAuthorizationServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAuthorizationServerParamsWithHTTPClient(client *http.Client) *CreateAuthorizationServerParams {
	var (
		tidDefault = string("default")
	)
	return &CreateAuthorizationServerParams{
		Tid:        tidDefault,
		HTTPClient: client,
	}
}

/*CreateAuthorizationServerParams contains all the parameters to send to the API endpoint
for the create authorization server operation typically these are written to a http.Request
*/
type CreateAuthorizationServerParams struct {

	/*Server*/
	Server *models.Server
	/*InitialProfile
	  Initial profile

	*/
	InitialProfile *string
	/*Tid
	  Tenant id

	*/
	Tid string
	/*WithDemoClient
	  With demo client

	*/
	WithDemoClient *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create authorization server params
func (o *CreateAuthorizationServerParams) WithTimeout(timeout time.Duration) *CreateAuthorizationServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create authorization server params
func (o *CreateAuthorizationServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create authorization server params
func (o *CreateAuthorizationServerParams) WithContext(ctx context.Context) *CreateAuthorizationServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create authorization server params
func (o *CreateAuthorizationServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create authorization server params
func (o *CreateAuthorizationServerParams) WithHTTPClient(client *http.Client) *CreateAuthorizationServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create authorization server params
func (o *CreateAuthorizationServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServer adds the server to the create authorization server params
func (o *CreateAuthorizationServerParams) WithServer(server *models.Server) *CreateAuthorizationServerParams {
	o.SetServer(server)
	return o
}

// SetServer adds the server to the create authorization server params
func (o *CreateAuthorizationServerParams) SetServer(server *models.Server) {
	o.Server = server
}

// WithInitialProfile adds the initialProfile to the create authorization server params
func (o *CreateAuthorizationServerParams) WithInitialProfile(initialProfile *string) *CreateAuthorizationServerParams {
	o.SetInitialProfile(initialProfile)
	return o
}

// SetInitialProfile adds the initialProfile to the create authorization server params
func (o *CreateAuthorizationServerParams) SetInitialProfile(initialProfile *string) {
	o.InitialProfile = initialProfile
}

// WithTid adds the tid to the create authorization server params
func (o *CreateAuthorizationServerParams) WithTid(tid string) *CreateAuthorizationServerParams {
	o.SetTid(tid)
	return o
}

// SetTid adds the tid to the create authorization server params
func (o *CreateAuthorizationServerParams) SetTid(tid string) {
	o.Tid = tid
}

// WithWithDemoClient adds the withDemoClient to the create authorization server params
func (o *CreateAuthorizationServerParams) WithWithDemoClient(withDemoClient *bool) *CreateAuthorizationServerParams {
	o.SetWithDemoClient(withDemoClient)
	return o
}

// SetWithDemoClient adds the withDemoClient to the create authorization server params
func (o *CreateAuthorizationServerParams) SetWithDemoClient(withDemoClient *bool) {
	o.WithDemoClient = withDemoClient
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthorizationServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Server != nil {
		if err := r.SetBodyParam(o.Server); err != nil {
			return err
		}
	}

	if o.InitialProfile != nil {

		// query param initial_profile
		var qrInitialProfile string
		if o.InitialProfile != nil {
			qrInitialProfile = *o.InitialProfile
		}
		qInitialProfile := qrInitialProfile
		if qInitialProfile != "" {
			if err := r.SetQueryParam("initial_profile", qInitialProfile); err != nil {
				return err
			}
		}

	}

	// path param tid
	if err := r.SetPathParam("tid", o.Tid); err != nil {
		return err
	}

	if o.WithDemoClient != nil {

		// query param with_demo_client
		var qrWithDemoClient bool
		if o.WithDemoClient != nil {
			qrWithDemoClient = *o.WithDemoClient
		}
		qWithDemoClient := swag.FormatBool(qrWithDemoClient)
		if qWithDemoClient != "" {
			if err := r.SetQueryParam("with_demo_client", qWithDemoClient); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}