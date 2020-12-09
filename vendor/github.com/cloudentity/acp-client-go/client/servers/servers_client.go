// Code generated by go-swagger; DO NOT EDIT.

package servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new servers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for servers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	BindServer(params *BindServerParams, authInfo runtime.ClientAuthInfoWriter) (*BindServerOK, error)

	CreateAuthorizationServer(params *CreateAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAuthorizationServerCreated, error)

	DeleteAuthorizationServer(params *DeleteAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAuthorizationServerNoContent, error)

	GetAuthorizationServer(params *GetAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*GetAuthorizationServerOK, error)

	GetServerConsent(params *GetServerConsentParams, authInfo runtime.ClientAuthInfoWriter) (*GetServerConsentOK, error)

	GetServerForDeveloper(params *GetServerForDeveloperParams, authInfo runtime.ClientAuthInfoWriter) (*GetServerForDeveloperOK, error)

	ListAuthorizationServers(params *ListAuthorizationServersParams, authInfo runtime.ClientAuthInfoWriter) (*ListAuthorizationServersOK, error)

	ListDashboards(params *ListDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*ListDashboardsOK, error)

	ListServersBindings(params *ListServersBindingsParams, authInfo runtime.ClientAuthInfoWriter) (*ListServersBindingsOK, error)

	ListServersForDeveloper(params *ListServersForDeveloperParams, authInfo runtime.ClientAuthInfoWriter) (*ListServersForDeveloperOK, error)

	SetServerConsent(params *SetServerConsentParams, authInfo runtime.ClientAuthInfoWriter) (*SetServerConsentOK, error)

	UnbindServer(params *UnbindServerParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindServerOK, error)

	UpdateAuthorizationServer(params *UpdateAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAuthorizationServerOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BindServer binds server

  Bind server.
*/
func (a *Client) BindServer(params *BindServerParams, authInfo runtime.ClientAuthInfoWriter) (*BindServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBindServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "bindServer",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/bind/{rid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BindServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BindServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for bindServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateAuthorizationServer creates authorization server

  Multiple authorization servers with unique id can be created within a tenant.
If id and secret are not provided, will be generated.
Secret if provided must have at least 32 characters.

You can set what grant types will be supported by authorization server. The defaults are:
`{"grant_types": ["authorization_code", "implicit", "client_credentials", "refresh_token"]}`

If jwks keys are not provided explicitly, will be generated based on provided `key_type` algorithm (rsa by default).

TTLs for tokens and authorization code can be customized. The defaults are:

`authorization_code_ttl` - 10 minutes
`access_token_ttl` - 1 hour
`id_token_ttl` - 1 hour
`refresh_token_ttl` - 30 days

If you want to enable dynamic client registration set `{"enable_dynamic_client_registration": true}`.

If you want to create FAPI read write compliant server set: `{"profiles"": ["fapi_rw"]}`.

If you want to enforce PKCE set `{"enforce_pkce": true}`.
*/
func (a *Client) CreateAuthorizationServer(params *CreateAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAuthorizationServerCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAuthorizationServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createAuthorizationServer",
		Method:             "POST",
		PathPattern:        "/api/admin/{tid}/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAuthorizationServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAuthorizationServerCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createAuthorizationServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAuthorizationServer deletes authorization server

  Delete authorization server.
*/
func (a *Client) DeleteAuthorizationServer(params *DeleteAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAuthorizationServerNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAuthorizationServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAuthorizationServer",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/servers/{aid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAuthorizationServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAuthorizationServerNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteAuthorizationServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAuthorizationServer gets authorization server

  Get authorization server.
*/
func (a *Client) GetAuthorizationServer(params *GetAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*GetAuthorizationServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthorizationServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAuthorizationServer",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers/{aid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuthorizationServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthorizationServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAuthorizationServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServerConsent gets server consent

  Get server consent.
*/
func (a *Client) GetServerConsent(params *GetServerConsentParams, authInfo runtime.ClientAuthInfoWriter) (*GetServerConsentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServerConsent",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/server-consent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerConsentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServerConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetServerForDeveloper gets authorization server

  Returns authorization server details with list of scopes.
*/
func (a *Client) GetServerForDeveloper(params *GetServerForDeveloperParams, authInfo runtime.ClientAuthInfoWriter) (*GetServerForDeveloperOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServerForDeveloperParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getServerForDeveloper",
		Method:             "GET",
		PathPattern:        "/api/developer/{tid}/{aid}/servers/{rid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetServerForDeveloperReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServerForDeveloperOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getServerForDeveloper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListAuthorizationServers lists authorization servers

  List authorization servers.
*/
func (a *Client) ListAuthorizationServers(params *ListAuthorizationServersParams, authInfo runtime.ClientAuthInfoWriter) (*ListAuthorizationServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAuthorizationServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAuthorizationServers",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAuthorizationServersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAuthorizationServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAuthorizationServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListDashboards lists links to dashboards

  List links to dashboards.
*/
func (a *Client) ListDashboards(params *ListDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*ListDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDashboards",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/dashboards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDashboardsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listDashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListServersBindings lists servers bindings

  List servers bindings.
*/
func (a *Client) ListServersBindings(params *ListServersBindingsParams, authInfo runtime.ClientAuthInfoWriter) (*ListServersBindingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServersBindingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServersBindings",
		Method:             "GET",
		PathPattern:        "/api/admin/{tid}/servers-bindings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServersBindingsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServersBindingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServersBindings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListServersForDeveloper lists authorization servers

  Returns list of authorization severs.
*/
func (a *Client) ListServersForDeveloper(params *ListServersForDeveloperParams, authInfo runtime.ClientAuthInfoWriter) (*ListServersForDeveloperOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServersForDeveloperParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listServersForDeveloper",
		Method:             "GET",
		PathPattern:        "/api/developer/{tid}/{aid}/servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListServersForDeveloperReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServersForDeveloperOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listServersForDeveloper: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetServerConsent sets server consent

  Set server consent. For custom server consent a client in system server is created automatically.
*/
func (a *Client) SetServerConsent(params *SetServerConsentParams, authInfo runtime.ClientAuthInfoWriter) (*SetServerConsentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetServerConsentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setServerConsent",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/server-consent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetServerConsentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetServerConsentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setServerConsent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UnbindServer unbinds server

  Unbind server.
*/
func (a *Client) UnbindServer(params *UnbindServerParams, authInfo runtime.ClientAuthInfoWriter) (*UnbindServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnbindServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "unbindServer",
		Method:             "DELETE",
		PathPattern:        "/api/admin/{tid}/servers/{aid}/unbind/{rid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UnbindServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnbindServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for unbindServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateAuthorizationServer updates authorization server

  Update authorization server.
*/
func (a *Client) UpdateAuthorizationServer(params *UpdateAuthorizationServerParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateAuthorizationServerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAuthorizationServerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateAuthorizationServer",
		Method:             "PUT",
		PathPattern:        "/api/admin/{tid}/servers/{aid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateAuthorizationServerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateAuthorizationServerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateAuthorizationServer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}