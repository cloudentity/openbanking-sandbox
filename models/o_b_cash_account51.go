// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBCashAccount51 Provides the details to identify the beneficiary account.
//
// swagger:model OBCashAccount5_1
type OBCashAccount51 struct {

	// Beneficiary account identification.
	// Required: true
	// Max Length: 256
	// Min Length: 1
	Identification *string `json:"Identification"`

	// name
	Name Name0 `json:"Name,omitempty"`

	// scheme name
	// Required: true
	SchemeName *OBExternalAccountIdentification4Code `json:"SchemeName"`

	// secondary identification
	SecondaryIdentification SecondaryIdentification `json:"SecondaryIdentification,omitempty"`
}

// Validate validates this o b cash account5 1
func (m *OBCashAccount51) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBCashAccount51) validateIdentification(formats strfmt.Registry) error {

	if err := validate.Required("Identification", "body", m.Identification); err != nil {
		return err
	}

	if err := validate.MinLength("Identification", "body", *m.Identification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Identification", "body", *m.Identification, 256); err != nil {
		return err
	}

	return nil
}

func (m *OBCashAccount51) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Name")
		}
		return err
	}

	return nil
}

func (m *OBCashAccount51) validateSchemeName(formats strfmt.Registry) error {

	if err := validate.Required("SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if err := validate.Required("SchemeName", "body", m.SchemeName); err != nil {
		return err
	}

	if m.SchemeName != nil {
		if err := m.SchemeName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBCashAccount51) validateSecondaryIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryIdentification) { // not required
		return nil
	}

	if err := m.SecondaryIdentification.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecondaryIdentification")
		}
		return err
	}

	return nil
}

// ContextValidate validate this o b cash account5 1 based on the context it is used
func (m *OBCashAccount51) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemeName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryIdentification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBCashAccount51) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Name.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Name")
		}
		return err
	}

	return nil
}

func (m *OBCashAccount51) contextValidateSchemeName(ctx context.Context, formats strfmt.Registry) error {

	if m.SchemeName != nil {
		if err := m.SchemeName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("SchemeName")
			}
			return err
		}
	}

	return nil
}

func (m *OBCashAccount51) contextValidateSecondaryIdentification(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SecondaryIdentification.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SecondaryIdentification")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBCashAccount51) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBCashAccount51) UnmarshalBinary(b []byte) error {
	var res OBCashAccount51
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
