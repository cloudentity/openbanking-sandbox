// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBReadStandingOrder6 o b read standing order6
//
// swagger:model OBReadStandingOrder6
type OBReadStandingOrder6 struct {

	// data
	// Required: true
	Data *OBReadStandingOrder6Data `json:"Data"`

	// links
	Links *Links `json:"Links,omitempty"`

	// meta
	Meta *Meta `json:"Meta,omitempty"`
}

// Validate validates this o b read standing order6
func (m *OBReadStandingOrder6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBReadStandingOrder6) validateData(formats strfmt.Registry) error {

	if err := validate.Required("Data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Data")
			}
			return err
		}
	}

	return nil
}

func (m *OBReadStandingOrder6) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Links")
			}
			return err
		}
	}

	return nil
}

func (m *OBReadStandingOrder6) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBReadStandingOrder6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBReadStandingOrder6) UnmarshalBinary(b []byte) error {
	var res OBReadStandingOrder6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBReadStandingOrder6Data o b read standing order6 data
//
// swagger:model OBReadStandingOrder6Data
type OBReadStandingOrder6Data struct {

	// standing order
	StandingOrder []*OBStandingOrder6 `json:"StandingOrder"`
}

// Validate validates this o b read standing order6 data
func (m *OBReadStandingOrder6Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandingOrder(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBReadStandingOrder6Data) validateStandingOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.StandingOrder) { // not required
		return nil
	}

	for i := 0; i < len(m.StandingOrder); i++ {
		if swag.IsZero(m.StandingOrder[i]) { // not required
			continue
		}

		if m.StandingOrder[i] != nil {
			if err := m.StandingOrder[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Data" + "." + "StandingOrder" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBReadStandingOrder6Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBReadStandingOrder6Data) UnmarshalBinary(b []byte) error {
	var res OBReadStandingOrder6Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
