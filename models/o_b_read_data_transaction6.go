// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OBReadDataTransaction6 o b read data transaction6
//
// swagger:model OBReadDataTransaction6
type OBReadDataTransaction6 struct {

	// transaction
	Transaction []*OBTransaction6 `json:"Transaction"`
}

// Validate validates this o b read data transaction6
func (m *OBReadDataTransaction6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBReadDataTransaction6) validateTransaction(formats strfmt.Registry) error {

	if swag.IsZero(m.Transaction) { // not required
		return nil
	}

	for i := 0; i < len(m.Transaction); i++ {
		if swag.IsZero(m.Transaction[i]) { // not required
			continue
		}

		if m.Transaction[i] != nil {
			if err := m.Transaction[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Transaction" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBReadDataTransaction6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBReadDataTransaction6) UnmarshalBinary(b []byte) error {
	var res OBReadDataTransaction6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}