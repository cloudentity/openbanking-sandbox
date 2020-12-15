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

// ConsentGrant consent grant
//
// swagger:model ConsentGrant
type ConsentGrant struct {

	// consent grant id
	ConsentGrantActID string `json:"consent_grant_act_id,omitempty"`

	// consent id
	ConsentID string `json:"consent_id,omitempty"`

	// given at timestamp
	// Format: date-time
	GivenAt strfmt.DateTime `json:"given_at,omitempty"`

	// grant type, one of: implicit, explicit
	GrantType string `json:"grant_type,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// tenant id
	TenantID string `json:"tenant_id,omitempty"`

	// version
	Version int64 `json:"version,omitempty"`
}

// Validate validates this consent grant
func (m *ConsentGrant) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGivenAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsentGrant) validateGivenAt(formats strfmt.Registry) error {

	if swag.IsZero(m.GivenAt) { // not required
		return nil
	}

	if err := validate.FormatOf("given_at", "body", "date-time", m.GivenAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsentGrant) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentGrant) UnmarshalBinary(b []byte) error {
	var res ConsentGrant
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}