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

// OBBalanceType1Code Balance type, in a coded form.
//
// swagger:model OBBalanceType1Code
type OBBalanceType1Code string

const (

	// OBBalanceType1CodeClosingAvailable captures enum value "ClosingAvailable"
	OBBalanceType1CodeClosingAvailable OBBalanceType1Code = "ClosingAvailable"

	// OBBalanceType1CodeClosingBooked captures enum value "ClosingBooked"
	OBBalanceType1CodeClosingBooked OBBalanceType1Code = "ClosingBooked"

	// OBBalanceType1CodeClosingCleared captures enum value "ClosingCleared"
	OBBalanceType1CodeClosingCleared OBBalanceType1Code = "ClosingCleared"

	// OBBalanceType1CodeExpected captures enum value "Expected"
	OBBalanceType1CodeExpected OBBalanceType1Code = "Expected"

	// OBBalanceType1CodeForwardAvailable captures enum value "ForwardAvailable"
	OBBalanceType1CodeForwardAvailable OBBalanceType1Code = "ForwardAvailable"

	// OBBalanceType1CodeInformation captures enum value "Information"
	OBBalanceType1CodeInformation OBBalanceType1Code = "Information"

	// OBBalanceType1CodeInterimAvailable captures enum value "InterimAvailable"
	OBBalanceType1CodeInterimAvailable OBBalanceType1Code = "InterimAvailable"

	// OBBalanceType1CodeInterimBooked captures enum value "InterimBooked"
	OBBalanceType1CodeInterimBooked OBBalanceType1Code = "InterimBooked"

	// OBBalanceType1CodeInterimCleared captures enum value "InterimCleared"
	OBBalanceType1CodeInterimCleared OBBalanceType1Code = "InterimCleared"

	// OBBalanceType1CodeOpeningAvailable captures enum value "OpeningAvailable"
	OBBalanceType1CodeOpeningAvailable OBBalanceType1Code = "OpeningAvailable"

	// OBBalanceType1CodeOpeningBooked captures enum value "OpeningBooked"
	OBBalanceType1CodeOpeningBooked OBBalanceType1Code = "OpeningBooked"

	// OBBalanceType1CodeOpeningCleared captures enum value "OpeningCleared"
	OBBalanceType1CodeOpeningCleared OBBalanceType1Code = "OpeningCleared"

	// OBBalanceType1CodePreviouslyClosedBooked captures enum value "PreviouslyClosedBooked"
	OBBalanceType1CodePreviouslyClosedBooked OBBalanceType1Code = "PreviouslyClosedBooked"
)

// for schema
var oBBalanceType1CodeEnum []interface{}

func init() {
	var res []OBBalanceType1Code
	if err := json.Unmarshal([]byte(`["ClosingAvailable","ClosingBooked","ClosingCleared","Expected","ForwardAvailable","Information","InterimAvailable","InterimBooked","InterimCleared","OpeningAvailable","OpeningBooked","OpeningCleared","PreviouslyClosedBooked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBBalanceType1CodeEnum = append(oBBalanceType1CodeEnum, v)
	}
}

func (m OBBalanceType1Code) validateOBBalanceType1CodeEnum(path, location string, value OBBalanceType1Code) error {
	if err := validate.EnumCase(path, location, value, oBBalanceType1CodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this o b balance type1 code
func (m OBBalanceType1Code) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOBBalanceType1CodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
