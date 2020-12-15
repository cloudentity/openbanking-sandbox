// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScopeWithServiceDeveloperResponse scope with service developer response
//
// swagger:model ScopeWithServiceDeveloperResponse
type ScopeWithServiceDeveloperResponse struct {

	// scope description
	ScopeDescription string `json:"scope_description,omitempty"`

	// scope display name
	ScopeDisplayName string `json:"scope_display_name,omitempty"`

	// scope id
	ScopeID string `json:"scope_id,omitempty"`

	// scope name
	ScopeName string `json:"scope_name,omitempty"`

	// service description
	ServiceDescription string `json:"service_description,omitempty"`

	// service id
	ServiceID string `json:"service_id,omitempty"`

	// service name
	ServiceName string `json:"service_name,omitempty"`

	// is scope assigned to a service
	WithService bool `json:"with_service,omitempty"`
}

// Validate validates this scope with service developer response
func (m *ScopeWithServiceDeveloperResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScopeWithServiceDeveloperResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScopeWithServiceDeveloperResponse) UnmarshalBinary(b []byte) error {
	var res ScopeWithServiceDeveloperResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}