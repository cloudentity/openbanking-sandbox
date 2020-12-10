// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// OBCreditDebitCode1 Indicates whether the transaction is a credit or a debit entry.
//
// swagger:model OBCreditDebitCode_1
type OBCreditDebitCode1 string

const (

	// OBCreditDebitCode1Credit captures enum value "Credit"
	OBCreditDebitCode1Credit OBCreditDebitCode1 = "Credit"

	// OBCreditDebitCode1Debit captures enum value "Debit"
	OBCreditDebitCode1Debit OBCreditDebitCode1 = "Debit"
)

// for schema
var oBCreditDebitCode1Enum []interface{}

func init() {
	var res []OBCreditDebitCode1
	if err := json.Unmarshal([]byte(`["Credit","Debit"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBCreditDebitCode1Enum = append(oBCreditDebitCode1Enum, v)
	}
}

func (m OBCreditDebitCode1) validateOBCreditDebitCode1Enum(path, location string, value OBCreditDebitCode1) error {
	if err := validate.EnumCase(path, location, value, oBCreditDebitCode1Enum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b credit debit code 1
func (m OBCreditDebitCode1) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBCreditDebitCode1Enum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
