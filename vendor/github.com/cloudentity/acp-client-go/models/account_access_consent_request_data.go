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

// AccountAccessConsentRequestData account access consent request data
//
// swagger:model AccountAccessConsentRequestData
type AccountAccessConsentRequestData struct {

	// Specified date and time the permissions will expire. If this is not populated,
	// the permissions will be open ended.
	// Format: date-time
	ExpirationDateTime strfmt.DateTime `json:"ExpirationDateTime,omitempty"`

	// Specifies the Open Banking account access data types. This is a list of the data clusters
	// being consented by the PSU, and requested for authorisation with the ASPSP.
	Permissions []string `json:"Permissions"`

	// Specified start date and time for the transaction query period. If this is not populated,
	// the start date will be open ended, and data will be returned from the earliest available transaction.
	// Format: date-time
	TransactionFromDateTime strfmt.DateTime `json:"TransactionFromDateTime,omitempty"`

	// Specified end date and time for the transaction query period. If this is not populated,
	// the end date will be open ended, and data will be returned to the latest available transaction.
	// Format: date-time
	TransactionToDateTime strfmt.DateTime `json:"TransactionToDateTime,omitempty"`
}

// Validate validates this account access consent request data
func (m *AccountAccessConsentRequestData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpirationDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionFromDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionToDateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAccessConsentRequestData) validateExpirationDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("ExpirationDateTime", "body", "date-time", m.ExpirationDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAccessConsentRequestData) validateTransactionFromDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionFromDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("TransactionFromDateTime", "body", "date-time", m.TransactionFromDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAccessConsentRequestData) validateTransactionToDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionToDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("TransactionToDateTime", "body", "date-time", m.TransactionToDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAccessConsentRequestData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAccessConsentRequestData) UnmarshalBinary(b []byte) error {
	var res AccountAccessConsentRequestData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
