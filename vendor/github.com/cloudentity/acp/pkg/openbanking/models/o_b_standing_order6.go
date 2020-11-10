// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OBStandingOrder6 o b standing order6
//
// swagger:model OBStandingOrder6
type OBStandingOrder6 struct {

	// account Id
	// Required: true
	AccountID AccountID `json:"AccountId"`

	// creditor account
	CreditorAccount *OBCashAccount51 `json:"CreditorAccount,omitempty"`

	// creditor agent
	CreditorAgent *OBBranchAndFinancialInstitutionIdentification51 `json:"CreditorAgent,omitempty"`

	// final payment amount
	FinalPaymentAmount *OBActiveOrHistoricCurrencyAndAmount4 `json:"FinalPaymentAmount,omitempty"`

	// final payment date time
	// Format: date-time
	FinalPaymentDateTime FinalPaymentDateTime `json:"FinalPaymentDateTime,omitempty"`

	// first payment amount
	FirstPaymentAmount *OBActiveOrHistoricCurrencyAndAmount2 `json:"FirstPaymentAmount,omitempty"`

	// first payment date time
	// Format: date-time
	FirstPaymentDateTime FirstPaymentDateTime `json:"FirstPaymentDateTime,omitempty"`

	// frequency
	// Required: true
	Frequency Frequency1 `json:"Frequency"`

	// last payment amount
	LastPaymentAmount *OBActiveOrHistoricCurrencyAndAmount11 `json:"LastPaymentAmount,omitempty"`

	// last payment date time
	// Format: date-time
	LastPaymentDateTime LastPaymentDateTime `json:"LastPaymentDateTime,omitempty"`

	// next payment amount
	NextPaymentAmount *OBActiveOrHistoricCurrencyAndAmount3 `json:"NextPaymentAmount,omitempty"`

	// next payment date time
	// Format: date-time
	NextPaymentDateTime NextPaymentDateTime `json:"NextPaymentDateTime,omitempty"`

	// number of payments
	NumberOfPayments NumberOfPayments `json:"NumberOfPayments,omitempty"`

	// reference
	Reference Reference `json:"Reference,omitempty"`

	// standing order Id
	StandingOrderID StandingOrderID `json:"StandingOrderId,omitempty"`

	// standing order status code
	StandingOrderStatusCode OBExternalStandingOrderStatus1Code `json:"StandingOrderStatusCode,omitempty"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b standing order6
func (m *OBStandingOrder6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditorAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextPaymentAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextPaymentDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfPayments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStandingOrderStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStandingOrder6) validateAccountID(formats strfmt.Registry) error {

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateCreditorAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditorAccount) { // not required
		return nil
	}

	if m.CreditorAccount != nil {
		if err := m.CreditorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAccount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6) validateCreditorAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.CreditorAgent) { // not required
		return nil
	}

	if m.CreditorAgent != nil {
		if err := m.CreditorAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CreditorAgent")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6) validateFinalPaymentAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.FinalPaymentAmount) { // not required
		return nil
	}

	if m.FinalPaymentAmount != nil {
		if err := m.FinalPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FinalPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6) validateFinalPaymentDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FinalPaymentDateTime) { // not required
		return nil
	}

	if err := m.FinalPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FinalPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateFirstPaymentAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstPaymentAmount) { // not required
		return nil
	}

	if m.FirstPaymentAmount != nil {
		if err := m.FirstPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FirstPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6) validateFirstPaymentDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstPaymentDateTime) { // not required
		return nil
	}

	if err := m.FirstPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FirstPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateFrequency(formats strfmt.Registry) error {

	if err := m.Frequency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Frequency")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateLastPaymentAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPaymentAmount) { // not required
		return nil
	}

	if m.LastPaymentAmount != nil {
		if err := m.LastPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6) validateLastPaymentDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastPaymentDateTime) { // not required
		return nil
	}

	if err := m.LastPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LastPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateNextPaymentAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.NextPaymentAmount) { // not required
		return nil
	}

	if m.NextPaymentAmount != nil {
		if err := m.NextPaymentAmount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NextPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6) validateNextPaymentDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.NextPaymentDateTime) { // not required
		return nil
	}

	if err := m.NextPaymentDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NextPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateNumberOfPayments(formats strfmt.Registry) error {

	if swag.IsZero(m.NumberOfPayments) { // not required
		return nil
	}

	if err := m.NumberOfPayments.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NumberOfPayments")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateReference(formats strfmt.Registry) error {

	if swag.IsZero(m.Reference) { // not required
		return nil
	}

	if err := m.Reference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateStandingOrderID(formats strfmt.Registry) error {

	if swag.IsZero(m.StandingOrderID) { // not required
		return nil
	}

	if err := m.StandingOrderID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderId")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6) validateStandingOrderStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.StandingOrderStatusCode) { // not required
		return nil
	}

	if err := m.StandingOrderStatusCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderStatusCode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStandingOrder6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStandingOrder6) UnmarshalBinary(b []byte) error {
	var res OBStandingOrder6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
