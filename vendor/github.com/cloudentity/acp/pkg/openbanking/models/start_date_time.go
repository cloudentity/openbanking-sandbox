// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StartDateTime Date and time at which the statement period starts.All dates in the JSON payloads are represented in ISO 8601 date-time format.
// All date-time fields in responses must include the timezone. An example is below:
// 2017-04-05T10:43:07+00:00
//
// swagger:model StartDateTime
type StartDateTime strfmt.DateTime

// UnmarshalJSON sets a StartDateTime value from JSON input
func (m *StartDateTime) UnmarshalJSON(b []byte) error {
	return ((*strfmt.DateTime)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a StartDateTime value as JSON output
func (m StartDateTime) MarshalJSON() ([]byte, error) {
	return (strfmt.DateTime(m)).MarshalJSON()
}

// Validate validates this start date time
func (m StartDateTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.FormatOf("", "body", "date-time", strfmt.DateTime(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *StartDateTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartDateTime) UnmarshalBinary(b []byte) error {
	var res StartDateTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}