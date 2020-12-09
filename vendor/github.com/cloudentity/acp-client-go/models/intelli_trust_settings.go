// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IntelliTrustSettings intelli trust settings
//
// swagger:model IntelliTrustSettings
type IntelliTrustSettings struct {

	// OAuth client identifier
	ClientID string `json:"client_id,omitempty"`

	// domain
	Domain string `json:"domain,omitempty"`

	// flag to fetch groups. IDP must release groups claim in userinfo endpoint
	FetchGroups bool `json:"fetch_groups,omitempty"`

	// flag to fetch additional user data from userinfo endpoint
	GetUserInfo bool `json:"get_user_info,omitempty"`

	// OAuth redirect URL
	RedirectURL string `json:"redirect_url,omitempty"`

	// OAuth scopes which client will be requesting
	Scopes []string `json:"scopes"`
}

// Validate validates this intelli trust settings
func (m *IntelliTrustSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntelliTrustSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntelliTrustSettings) UnmarshalBinary(b []byte) error {
	var res IntelliTrustSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
