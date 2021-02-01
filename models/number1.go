// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Number1 fee/charges are captured dependent on the number of occurrences rather than capped at a particular amount
//
// swagger:model Number_1
type Number1 int64

// Validate validates this number 1
func (m Number1) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this number 1 based on context it is used
func (m Number1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
