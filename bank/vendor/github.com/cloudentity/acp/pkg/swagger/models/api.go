// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// API API
//
// swagger:model API
type API struct {

	// if false it is not possible to assign a policy
	CanHavePolicy bool `json:"can_have_policy,omitempty"`

	// Data classifications
	DataClassifications []string `json:"data_classifications"`

	// scope id
	ID string `json:"id,omitempty"`

	// HTTP request method
	Method string `json:"method,omitempty"`

	// HTTP request path
	Path string `json:"path,omitempty"`

	// optional id of a policy
	PolicyID string `json:"policy_id,omitempty"`

	// position of the api in the apis list
	Position int64 `json:"position,omitempty"`

	// server id
	ServerID string `json:"server_id,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`
}

// Validate validates this API
func (m *API) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *API) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *API) UnmarshalBinary(b []byte) error {
	var res API
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}