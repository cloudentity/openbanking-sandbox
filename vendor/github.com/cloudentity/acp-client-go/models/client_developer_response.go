// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClientDeveloperResponse client developer response
//
// swagger:model ClientDeveloperResponse
type ClientDeveloperResponse struct {

	// Kind of the application. The default, if omitted, is web. The defined values are native or web.
	ApplicationType string `json:"application_type,omitempty"`

	// oauth client allowed audience
	Audience []string `json:"audience"`

	// Boolean value indicating server support for mutual TLS client certificate-bound access tokens. If omitted, the default value is "false".
	CertificateBoundAccessToken bool `json:"tls_client_certificate_bound_access_tokens,omitempty"`

	// Time at which the client identifier was issued. The time is represented as the number of seconds from
	// 1970-01-01T00:00:00Z as measured in UTC until the date/time of issuance
	ClientIDIssuedAt int64 `json:"client_id_issued_at,omitempty"`

	// SecretExpiresAt is an integer holding the time at which the client secret will expire or 0 if it will not expire.
	ClientSecretExpiresAt int64 `json:"client_secret_expires_at,omitempty"`

	// client URI
	ClientURI string `json:"client_uri,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// oauth client grant types, allowed values: password, refresh_token, client_credentials, implicit, authorization_code
	GrantTypes []string `json:"grant_types"`

	// oauth client id
	ID string `json:"client_id,omitempty"`

	// Algorithm for signing the ID Token issued to this Client.
	// The default value depends on authorization server configuration.
	IDTokenSignedResponseAlg string `json:"id_token_signed_response_alg,omitempty"`

	// URL of JSON Web Key Set containing the public keys used by the client to authenticate
	JSONWebKeysURI string `json:"jwks_uri,omitempty"`

	// logo URI
	LogoURI string `json:"logo_uri,omitempty"`

	// human redable name
	Name string `json:"client_name,omitempty"`

	// policy url to read about how the profile data will be used
	PolicyURI string `json:"policy_uri,omitempty"`

	// oauth allowed redirect URIs
	RedirectURIs []string `json:"redirect_uris"`

	// Signing algorithm for a request object
	RequestObjectSigningAlg string `json:"request_object_signing_alg,omitempty"`

	// Array of absolute URIs that points to the Request Object that holds authorization request parameters
	RequestURIs []string `json:"request_uris"`

	// oauth client response types, allowed values: token, id_token, code
	ResponseTypes []string `json:"response_types"`

	// A string containing the value of an expected dNSName SAN entry in the certificate
	SanDNS string `json:"tls_client_auth_san_dns,omitempty"`

	// A string containing the value of an expected rfc822Name SAN entry in the certificate
	SanEmail string `json:"tls_client_auth_san_email,omitempty"`

	// A string representation of an IP address in either dotted decimal notation (for IPv4) or colon-delimited hexadecimal (for IPv6, as defined in [RFC5952]) that is expected to be present as an iPAddress SAN entry in the certificate
	SanIP string `json:"tls_client_auth_san_ip,omitempty"`

	// A string containing the value of an expected uniformResourceIdentifier SAN entry in the certificate
	SanURI string `json:"tls_client_auth_san_uri,omitempty"`

	// Optional comma separated scopes for compatibility with spec
	Scope string `json:"scope,omitempty"`

	// oauth client scopes
	Scopes []string `json:"scopes"`

	// oauth client secret
	Secret string `json:"client_secret,omitempty"`

	// URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a
	// file with a single JSON array of redirect_uri values.
	SectorIdentifierURI string `json:"sector_identifier_uri,omitempty"`

	// Authorization server id
	ServerID string `json:"authorization_server_id,omitempty"`

	// An [RFC4514] string representation of the expected subject distinguished name of the certificate
	SubjectDN string `json:"tls_client_auth_subject_dn,omitempty"`

	// Subject identifier type
	SubjectType string `json:"subject_type,omitempty"`

	// tenant ID
	TenantID string `json:"tenant_id,omitempty"`

	// Signing algorithm for a token endpoint
	TokenEndpointAuthSigningAlg string `json:"token_endpoint_auth_signing_alg,omitempty"`

	// Token endpoint authentication method
	TokenEndpointAuthnMethod string `json:"token_endpoint_auth_method,omitempty"`

	// terms of service url
	TosURI string `json:"tos_uri,omitempty"`

	// JWS alg algorithm REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT
	// [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims
	// as a UTF-8 encoded JSON object using the application/json content-type.
	UserinfoSignedResponseAlg string `json:"userinfo_signed_response_alg,omitempty"`

	// jwks
	Jwks *JWKs `json:"jwks,omitempty"`

	// privacy
	Privacy *ClientPrivacy `json:"privacy,omitempty"`
}

// Validate validates this client developer response
func (m *ClientDeveloperResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJwks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientDeveloperResponse) validateJwks(formats strfmt.Registry) error {

	if swag.IsZero(m.Jwks) { // not required
		return nil
	}

	if m.Jwks != nil {
		if err := m.Jwks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jwks")
			}
			return err
		}
	}

	return nil
}

func (m *ClientDeveloperResponse) validatePrivacy(formats strfmt.Registry) error {

	if swag.IsZero(m.Privacy) { // not required
		return nil
	}

	if m.Privacy != nil {
		if err := m.Privacy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("privacy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientDeveloperResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientDeveloperResponse) UnmarshalBinary(b []byte) error {
	var res ClientDeveloperResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
