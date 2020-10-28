// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureIDP azure ID p
//
// swagger:model AzureIDP
type AzureIDP struct {

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// ID
	ID string `json:"id,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// server ID
	ServerID string `json:"authorization_server_id,omitempty"`

	// static a m r
	StaticAMR []string `json:"static_amr"`

	// tenant ID
	TenantID string `json:"tenant_id,omitempty"`

	// attributes
	Attributes Attributes `json:"attributes,omitempty"`

	// credentials
	Credentials *AzureCredentials `json:"credentials,omitempty"`

	// mappings
	Mappings Mappings `json:"mappings,omitempty"`

	// settings
	Settings *AzureSettings `json:"settings,omitempty"`
}

// Validate validates this azure ID p
func (m *AzureIDP) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AzureIDP) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if err := m.Attributes.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes")
		}
		return err
	}

	return nil
}

func (m *AzureIDP) validateCredentials(formats strfmt.Registry) error {

	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

func (m *AzureIDP) validateMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.Mappings) { // not required
		return nil
	}

	if err := m.Mappings.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mappings")
		}
		return err
	}

	return nil
}

func (m *AzureIDP) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AzureIDP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureIDP) UnmarshalBinary(b []byte) error {
	var res AzureIDP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}