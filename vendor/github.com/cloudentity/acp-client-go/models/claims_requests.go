// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClaimsRequests claims requests
//
// swagger:model ClaimsRequests
type ClaimsRequests struct {

	// ID token
	IDToken map[string]ClaimRequest `json:"id_token,omitempty"`

	// userinfo
	Userinfo map[string]ClaimRequest `json:"userinfo,omitempty"`
}

// Validate validates this claims requests
func (m *ClaimsRequests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIDToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserinfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimsRequests) validateIDToken(formats strfmt.Registry) error {

	if swag.IsZero(m.IDToken) { // not required
		return nil
	}

	for k := range m.IDToken {

		if err := validate.Required("id_token"+"."+k, "body", m.IDToken[k]); err != nil {
			return err
		}
		if val, ok := m.IDToken[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ClaimsRequests) validateUserinfo(formats strfmt.Registry) error {

	if swag.IsZero(m.Userinfo) { // not required
		return nil
	}

	for k := range m.Userinfo {

		if err := validate.Required("userinfo"+"."+k, "body", m.Userinfo[k]); err != nil {
			return err
		}
		if val, ok := m.Userinfo[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimsRequests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimsRequests) UnmarshalBinary(b []byte) error {
	var res ClaimsRequests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
