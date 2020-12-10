// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateTenant(params *CreateTenantParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTenantCreated, error)

	DeleteTenant(params *DeleteTenantParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTenantNoContent, error)

	ExportConfiguration(params *ExportConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ExportConfigurationOK, error)

	GetAdminTenant(params *GetAdminTenantParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdminTenantOK, error)

	GetTenant(params *GetTenantParams, authInfo runtime.ClientAuthInfoWriter) (*GetTenantOK, error)

	ImportConfiguration(params *ImportConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ImportConfigurationNoContent, error)

	UpdateAdminTenant(params *UpdateAdminTenantParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAdminTenantOK, error)

	UpdateTenant(params *UpdateTenantParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTenantOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateTenant creates tenant

  Tenant must have unique id, if it's not provided will be generated.

Tenant name must be provided.

Tenant url is optional and if omitted, it is set to ACP deployment url with
appened id of a tenant, example: https://example.com/default.

When tenant is created, preconfigured authorization servers with default scopes,
services and oauth clients are automatically created underneath. See API response for more details.
*/
func (a *Client) CreateTenant(params *CreateTenantParams, authInfo runtime.ClientAuthInfoWriter) (*CreateTenantCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createTenant",
		Method:             "POST",
		PathPattern:        "/api/system/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateTenantCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteTenant deletes tenant

  Deletes index in elasticsearch if exists as well.
*/
func (a *Client) DeleteTenant(params *DeleteTenantParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteTenantNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTenant",
		Method:             "DELETE",
		PathPattern:        "/api/system/tenants/{tid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTenantNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ExportConfiguration exports configuration

  Export entire tenant configuration as json.
*/
func (a *Client) ExportConfiguration(params *ExportConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ExportConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExportConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "exportConfiguration",
		Method:             "GET",
		PathPattern:        "/api/system/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ExportConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExportConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for exportConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAdminTenant gets tenant

  Get tenant.
*/
func (a *Client) GetAdminTenant(params *GetAdminTenantParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdminTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdminTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAdminTenant",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/tenant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAdminTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAdminTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAdminTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTenant gets tenant

  Get tenant system api.
*/
func (a *Client) GetTenant(params *GetTenantParams, authInfo runtime.ClientAuthInfoWriter) (*GetTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTenant",
		Method:             "GET",
		PathPattern:        "/api/system/tenants/{tid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportConfiguration imports configuration

  Allows to quickly import tenant basic configuration.
*/
func (a *Client) ImportConfiguration(params *ImportConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ImportConfigurationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importConfiguration",
		Method:             "PUT",
		PathPattern:        "/api/system/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportConfigurationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAdminTenant updates tenant

  Update tenant.
*/
func (a *Client) UpdateAdminTenant(params *UpdateAdminTenantParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAdminTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAdminTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAdminTenant",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/tenant",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAdminTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAdminTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAdminTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateTenant updates tenant

  Update tenant system api.
*/
func (a *Client) UpdateTenant(params *UpdateTenantParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTenant",
		Method:             "PUT",
		PathPattern:        "/api/system/tenants/{tid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTenantOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
