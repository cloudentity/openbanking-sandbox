// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OIDCAuthentication OIDC authentication settings.
//
// Provide OAuth client details here.
//
// swagger:model OIDCAuthentication
type OIDCAuthentication struct {

	// OAuth client identifier
	ClientID string `json:"client_id,omitempty"`

	// OAuth client secret
	ClientSecret string `json:"client_secret,omitempty"`

	// flag to fetch additional user attributes from userinfo endpoint
	GetUserInfo bool `json:"get_user_info,omitempty"`

	// Authorization server issuer URL
	IssuerURL string `json:"issuer_url,omitempty"`

	// OAuth redirect URL
	RedirectURL string `json:"redirect_url,omitempty"`

	// OAuth scopes which client will be requesting
	Scopes []string `json:"scopes"`
}

// Validate validates this o ID c authentication
func (m *OIDCAuthentication) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OIDCAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCAuthentication) UnmarshalBinary(b []byte) error {
	var res OIDCAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
