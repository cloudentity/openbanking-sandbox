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

// Mapping Attribute mapping.
//
// Source and target must be provided.
//
// swagger:model Mapping
type Mapping struct {

	// Weak decoding
	//
	// If AllowWeakDecoding is true, the decoder will make the following
	// "weak" conversions:
	//
	// bools to string (true = "1", false = "0")
	// numbers to string (base 10)
	// bools to int/uint (true = 1, false = 0)
	// strings to int/uint (base implied by prefix)
	// int to bool (true if value != 0)
	// string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
	// FALSE, false, False. Anything else is an error)
	// empty array = empty map and vice versa
	// negative numbers to overflowed uint values (base 10)
	// slice of maps to a merged map
	// single values are converted to slices if required. Each
	// element is weakly decoded.
	AllowWeakDecoding bool `json:"allow_weak_decoding,omitempty"`

	// Source attribute.
	//
	// Source path to attribute(s) which should be copied to the identity context.
	// Use '.' to copy everything.
	// Required: true
	Source *string `json:"source"`

	// Target attribute.
	//
	// Target path in the identity context where source attribute(s) should be pasted.
	// Use '.' to paste to context top level object.
	// Required: true
	Target *string `json:"target"`

	// Target type
	//
	// number, string, bool, number_array, string_array, bool_array or any
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this mapping
func (m *Mapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Mapping) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

func (m *Mapping) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Mapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Mapping) UnmarshalBinary(b []byte) error {
	var res Mapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}