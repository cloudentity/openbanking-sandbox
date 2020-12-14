// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NumberOfPayments Number of the payments that will be made in completing this frequency sequence including any executed since the sequence start date.
//
// swagger:model NumberOfPayments
type NumberOfPayments string

// Validate validates this number of payments
func (m NumberOfPayments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinLength("", "body", string(m), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("", "body", string(m), 35); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}