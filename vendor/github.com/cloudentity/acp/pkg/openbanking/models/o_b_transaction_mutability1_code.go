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

// OBTransactionMutability1Code Specifies the Mutability of the Transaction record.
//
// swagger:model OBTransactionMutability1Code
type OBTransactionMutability1Code string

const (

	// OBTransactionMutability1CodeMutable captures enum value "Mutable"
	OBTransactionMutability1CodeMutable OBTransactionMutability1Code = "Mutable"

	// OBTransactionMutability1CodeImmutable captures enum value "Immutable"
	OBTransactionMutability1CodeImmutable OBTransactionMutability1Code = "Immutable"
)

// for schema
var oBTransactionMutability1CodeEnum []interface{}

func init() {
	var res []OBTransactionMutability1Code
	if err := json.Unmarshal([]byte(`["Mutable","Immutable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBTransactionMutability1CodeEnum = append(oBTransactionMutability1CodeEnum, v)
	}
}

func (m OBTransactionMutability1Code) validateOBTransactionMutability1CodeEnum(path, location string, value OBTransactionMutability1Code) error {
	if err := validate.Enum(path, location, value, oBTransactionMutability1CodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b transaction mutability1 code
func (m OBTransactionMutability1Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBTransactionMutability1CodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
