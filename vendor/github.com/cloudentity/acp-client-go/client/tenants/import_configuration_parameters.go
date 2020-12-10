// Code generated by go-swagger; DO NOT EDIT.

package tenants

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

// NewImportConfigurationParams creates a new ImportConfigurationParams object
// with the default values initialized.
func NewImportConfigurationParams() *ImportConfigurationParams {
	var ()
	return &ImportConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewImportConfigurationParamsWithTimeout creates a new ImportConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewImportConfigurationParamsWithTimeout(timeout time.Duration) *ImportConfigurationParams {
	var ()
	return &ImportConfigurationParams{

		timeout: timeout,
	}
}

// NewImportConfigurationParamsWithContext creates a new ImportConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewImportConfigurationParamsWithContext(ctx context.Context) *ImportConfigurationParams {
	var ()
	return &ImportConfigurationParams{

		Context: ctx,
	}
}

// NewImportConfigurationParamsWithHTTPClient creates a new ImportConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewImportConfigurationParamsWithHTTPClient(client *http.Client) *ImportConfigurationParams {
	var ()
	return &ImportConfigurationParams{
		HTTPClient: client,
	}
}

/*ImportConfigurationParams contains all the parameters to send to the API endpoint
for the import configuration operation typically these are written to a http.Request
*/
type ImportConfigurationParams struct {

	/*Configuration*/
	Configuration *models.Dump
	/*Mode
	  Insert mode

	*/
	Mode *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the import configuration params
func (o *ImportConfigurationParams) WithTimeout(timeout time.Duration) *ImportConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import configuration params
func (o *ImportConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import configuration params
func (o *ImportConfigurationParams) WithContext(ctx context.Context) *ImportConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import configuration params
func (o *ImportConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import configuration params
func (o *ImportConfigurationParams) WithHTTPClient(client *http.Client) *ImportConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import configuration params
func (o *ImportConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfiguration adds the configuration to the import configuration params
func (o *ImportConfigurationParams) WithConfiguration(configuration *models.Dump) *ImportConfigurationParams {
	o.SetConfiguration(configuration)
	return o
}

// SetConfiguration adds the configuration to the import configuration params
func (o *ImportConfigurationParams) SetConfiguration(configuration *models.Dump) {
	o.Configuration = configuration
}

// WithMode adds the mode to the import configuration params
func (o *ImportConfigurationParams) WithMode(mode *string) *ImportConfigurationParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the import configuration params
func (o *ImportConfigurationParams) SetMode(mode *string) {
	o.Mode = mode
}

// WriteToRequest writes these params to a swagger request
func (o *ImportConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Configuration != nil {
		if err := r.SetBodyParam(o.Configuration); err != nil {
			return err
		}
	}

	if o.Mode != nil {

		// query param mode
		var qrMode string
		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {
			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
