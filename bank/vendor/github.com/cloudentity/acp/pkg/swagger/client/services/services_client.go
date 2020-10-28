// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateService(params *CreateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceCreated, error)

	DeleteService(params *DeleteServiceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceNoContent, error)

	GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceOK, error)

	ImportSpecificationFromFile(params *ImportSpecificationFromFileParams, authInfo runtime.ClientAuthInfoWriter) (*ImportSpecificationFromFileOK, error)

	ImportSpecificationFromText(params *ImportSpecificationFromTextParams, authInfo runtime.ClientAuthInfoWriter) (*ImportSpecificationFromTextOK, error)

	ImportSpecificationFromURL(params *ImportSpecificationFromURLParams, authInfo runtime.ClientAuthInfoWriter) (*ImportSpecificationFromURLOK, error)

	ListServices(params *ListServicesParams, authInfo runtime.ClientAuthInfoWriter) (*ListServicesOK, error)

	RemoveSpecification(params *RemoveSpecificationParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveSpecificationOK, error)

	UpdateService(params *UpdateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateService creates service

  Service name is required. Service id will be generated if not provided.
*/
func (a *Client) CreateService(params *CreateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateServiceCreated, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewCreateServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createService",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateServiceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteService deletes service

  Delete service.
*/
func (a *Client) DeleteService(params *DeleteServiceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteServiceNoContent, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewDeleteServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteService",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/services/{sid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteServiceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetService gets service

  Get service.
*/
func (a *Client) GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter) (*GetServiceOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getService",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/services/{sid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportSpecificationFromFile imports open API specification from file

  It removes all apis before import.
*/
func (a *Client) ImportSpecificationFromFile(params *ImportSpecificationFromFileParams, authInfo runtime.ClientAuthInfoWriter) (*ImportSpecificationFromFileOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewImportSpecificationFromFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importSpecificationFromFile",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/services/{sid}/apis/import/file",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportSpecificationFromFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportSpecificationFromFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importSpecificationFromFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportSpecificationFromText imports open API specification from text

  It removes all apis and policies (created by previous import) before import.
*/
func (a *Client) ImportSpecificationFromText(params *ImportSpecificationFromTextParams, authInfo runtime.ClientAuthInfoWriter) (*ImportSpecificationFromTextOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewImportSpecificationFromTextParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importSpecificationFromText",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/services/{sid}/apis/import/text",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"text/plain"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportSpecificationFromTextReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportSpecificationFromTextOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importSpecificationFromText: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportSpecificationFromURL imports open API specification from url

  It removes all apis before import.
*/
func (a *Client) ImportSpecificationFromURL(params *ImportSpecificationFromURLParams, authInfo runtime.ClientAuthInfoWriter) (*ImportSpecificationFromURLOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewImportSpecificationFromURLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importSpecificationFromURL",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/services/{sid}/apis/import/url",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportSpecificationFromURLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportSpecificationFromURLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importSpecificationFromURL: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListServices lists services

  List services.
*/
func (a *Client) ListServices(params *ListServicesParams, authInfo runtime.ClientAuthInfoWriter) (*ListServicesOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServices",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemoveSpecification removes imported apis and policies

  Remove all apis and policies created during import for a service.

If policy is used by another service it will not be removed.

If service is connected to a gateway, it will be disconnected.

If a gateway api group is connected to this service, it will be disconnected.
*/
func (a *Client) RemoveSpecification(params *RemoveSpecificationParams, authInfo runtime.ClientAuthInfoWriter) (*RemoveSpecificationOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewRemoveSpecificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removeSpecification",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/services/{sid}/apis",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveSpecificationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveSpecificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for removeSpecification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateService updates service

  Update service.
*/
func (a *Client) UpdateService(params *UpdateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateServiceOK, error) {
	// : Validate the params before sending
	if params == nil {
		params = NewUpdateServiceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateService",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/services/{sid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateService: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}