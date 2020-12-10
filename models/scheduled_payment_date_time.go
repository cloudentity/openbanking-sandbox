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

// ScheduledPaymentDateTime The date on which the scheduled payment will be made.All dates in the JSON payloads are represented in ISO 8601 date-time format.
// All date-time fields in responses must include the timezone. An example is below:
// 2017-04-05T10:43:07+00:00
//
// swagger:model ScheduledPaymentDateTime
type ScheduledPaymentDateTime strfmt.DateTime

// UnmarshalJSON sets a ScheduledPaymentDateTime value from JSON input
func (m *ScheduledPaymentDateTime) UnmarshalJSON(b []byte) error {
	return ((*strfmt.DateTime)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a ScheduledPaymentDateTime value as JSON output
func (m ScheduledPaymentDateTime) MarshalJSON() ([]byte, error) {
	return (strfmt.DateTime(m)).MarshalJSON()
}

// Validate validates this scheduled payment date time
func (m ScheduledPaymentDateTime) Validate(formats strfmt.Registry) error {
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
func (m *ScheduledPaymentDateTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduledPaymentDateTime) UnmarshalBinary(b []byte) error {
	var res ScheduledPaymentDateTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
