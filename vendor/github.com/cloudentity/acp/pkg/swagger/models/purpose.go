// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Purpose purpose
//
// swagger:model Purpose
type Purpose struct {

	// name
	Name string `json:"name,omitempty"`

	// is the purpose primary. If there are multiple purposes defined, one of them must be marked as primary.
	Primary bool `json:"primary,omitempty"`
}

// Validate validates this purpose
func (m *Purpose) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Purpose) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Purpose) UnmarshalBinary(b []byte) error {
	var res Purpose
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}