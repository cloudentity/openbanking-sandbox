// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OIDCCredentials o ID c credentials
//
// swagger:model OIDCCredentials
type OIDCCredentials struct {

	// client secret
	ClientSecret string `json:"client_secret,omitempty"`
}

// Validate validates this o ID c credentials
func (m *OIDCCredentials) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OIDCCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OIDCCredentials) UnmarshalBinary(b []byte) error {
	var res OIDCCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
