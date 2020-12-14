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

// OBBeneficiaryType1Code Specifies the Beneficiary Type.
//
// swagger:model OBBeneficiaryType1Code
type OBBeneficiaryType1Code string

const (

	// OBBeneficiaryType1CodeTrusted captures enum value "Trusted"
	OBBeneficiaryType1CodeTrusted OBBeneficiaryType1Code = "Trusted"

	// OBBeneficiaryType1CodeOrdinary captures enum value "Ordinary"
	OBBeneficiaryType1CodeOrdinary OBBeneficiaryType1Code = "Ordinary"
)

// for schema
var oBBeneficiaryType1CodeEnum []interface{}

func init() {
	var res []OBBeneficiaryType1Code
	if err := json.Unmarshal([]byte(`["Trusted","Ordinary"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBBeneficiaryType1CodeEnum = append(oBBeneficiaryType1CodeEnum, v)
	}
}

func (m OBBeneficiaryType1Code) validateOBBeneficiaryType1CodeEnum(path, location string, value OBBeneficiaryType1Code) error {
	if err := validate.EnumCase(path, location, value, oBBeneficiaryType1CodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b beneficiary type1 code
func (m OBBeneficiaryType1Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBBeneficiaryType1CodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}