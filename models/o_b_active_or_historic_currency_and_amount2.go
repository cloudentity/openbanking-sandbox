// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OBActiveOrHistoricCurrencyAndAmount2 The amount of the first Standing Order
//
// swagger:model OBActiveOrHistoricCurrencyAndAmount_2
type OBActiveOrHistoricCurrencyAndAmount2 struct {

	// amount
	// Required: true
	Amount OBActiveCurrencyAndAmountSimpleType `json:"Amount"`

	// currency
	// Required: true
	Currency ActiveOrHistoricCurrencyCode1 `json:"Currency"`
}

// Validate validates this o b active or historic currency and amount 2
func (m *OBActiveOrHistoricCurrencyAndAmount2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBActiveOrHistoricCurrencyAndAmount2) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Amount")
		}
		return err
	}

	return nil
}

func (m *OBActiveOrHistoricCurrencyAndAmount2) validateCurrency(formats strfmt.Registry) error {

	if err := m.Currency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Currency")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBActiveOrHistoricCurrencyAndAmount2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBActiveOrHistoricCurrencyAndAmount2) UnmarshalBinary(b []byte) error {
	var res OBActiveOrHistoricCurrencyAndAmount2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
