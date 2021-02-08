// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OBEntryStatus1Code Status of a transaction entry on the books of the account servicer.
//
// swagger:model OBEntryStatus1Code
type OBEntryStatus1Code string

const (

	// OBEntryStatus1CodeBooked captures enum value "Booked"
	OBEntryStatus1CodeBooked OBEntryStatus1Code = "Booked"

	// OBEntryStatus1CodePending captures enum value "Pending"
	OBEntryStatus1CodePending OBEntryStatus1Code = "Pending"
)

// for schema
var oBEntryStatus1CodeEnum []interface{}

func init() {
	var res []OBEntryStatus1Code
	if err := json.Unmarshal([]byte(`["Booked","Pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBEntryStatus1CodeEnum = append(oBEntryStatus1CodeEnum, v)
	}
}

func (m OBEntryStatus1Code) validateOBEntryStatus1CodeEnum(path, location string, value OBEntryStatus1Code) error {
	if err := validate.EnumCase(path, location, value, oBEntryStatus1CodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b entry status1 code
func (m OBEntryStatus1Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBEntryStatus1CodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this o b entry status1 code based on context it is used
func (m OBEntryStatus1Code) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}