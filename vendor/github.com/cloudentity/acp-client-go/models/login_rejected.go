// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LoginRejected login rejected
//
// swagger:model LoginRejected
type LoginRejected struct {

	// url where user should be redirected
	RedirectTo string `json:"redirect_to,omitempty"`
}

// Validate validates this login rejected
func (m *LoginRejected) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LoginRejected) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginRejected) UnmarshalBinary(b []byte) error {
	var res LoginRejected
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
