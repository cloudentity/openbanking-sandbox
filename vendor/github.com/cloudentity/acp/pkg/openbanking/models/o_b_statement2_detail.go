// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBStatement2Detail Provides further details on a statement resource.
//
// swagger:model OBStatement2Detail
type OBStatement2Detail struct {

	// account Id
	// Required: true
	AccountID AccountID `json:"AccountId"`

	// creation date time
	// Required: true
	// Format: date-time
	CreationDateTime CreationDateTime `json:"CreationDateTime"`

	// end date time
	// Required: true
	// Format: date-time
	EndDateTime EndDateTime `json:"EndDateTime"`

	// start date time
	// Required: true
	// Format: date-time
	StartDateTime StartDateTime `json:"StartDateTime"`

	// statement amount
	StatementAmount []*OBStatement2DetailStatementAmountItems0 `json:"StatementAmount"`

	// statement benefit
	StatementBenefit []*OBStatement2DetailStatementBenefitItems0 `json:"StatementBenefit"`

	// statement date time
	StatementDateTime []*OBStatement2DetailStatementDateTimeItems0 `json:"StatementDateTime"`

	// statement description
	StatementDescription []string `json:"StatementDescription"`

	// statement fee
	StatementFee []*OBStatement2DetailStatementFeeItems0 `json:"StatementFee"`

	// statement Id
	StatementID StatementID `json:"StatementId,omitempty"`

	// statement interest
	StatementInterest []*OBStatement2DetailStatementInterestItems0 `json:"StatementInterest"`

	// statement rate
	StatementRate []*OBStatement2DetailStatementRateItems0 `json:"StatementRate"`

	// statement reference
	StatementReference StatementReference `json:"StatementReference,omitempty"`

	// statement value
	StatementValue []*OBStatement2DetailStatementValueItems0 `json:"StatementValue"`

	// type
	// Required: true
	Type OBExternalStatementType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail
func (m *OBStatement2Detail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementBenefit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementInterest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatementValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2Detail) validateAccountID(formats strfmt.Registry) error {

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBStatement2Detail) validateCreationDateTime(formats strfmt.Registry) error {

	if err := m.CreationDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CreationDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStatement2Detail) validateEndDateTime(formats strfmt.Registry) error {

	if err := m.EndDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("EndDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStatement2Detail) validateStartDateTime(formats strfmt.Registry) error {

	if err := m.StartDateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StartDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStatement2Detail) validateStatementAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementAmount) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementAmount); i++ {
		if swag.IsZero(m.StatementAmount[i]) { // not required
			continue
		}

		if m.StatementAmount[i] != nil {
			if err := m.StatementAmount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementAmount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementBenefit(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementBenefit) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementBenefit); i++ {
		if swag.IsZero(m.StatementBenefit[i]) { // not required
			continue
		}

		if m.StatementBenefit[i] != nil {
			if err := m.StatementBenefit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementBenefit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementDateTime) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementDateTime); i++ {
		if swag.IsZero(m.StatementDateTime[i]) { // not required
			continue
		}

		if m.StatementDateTime[i] != nil {
			if err := m.StatementDateTime[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementDateTime" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementDescription) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementDescription); i++ {

		if err := validate.MinLength("StatementDescription"+"."+strconv.Itoa(i), "body", string(m.StatementDescription[i]), 1); err != nil {
			return err
		}

		if err := validate.MaxLength("StatementDescription"+"."+strconv.Itoa(i), "body", string(m.StatementDescription[i]), 500); err != nil {
			return err
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementFee(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementFee) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementFee); i++ {
		if swag.IsZero(m.StatementFee[i]) { // not required
			continue
		}

		if m.StatementFee[i] != nil {
			if err := m.StatementFee[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementFee" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementID(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementID) { // not required
		return nil
	}

	if err := m.StatementID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StatementId")
		}
		return err
	}

	return nil
}

func (m *OBStatement2Detail) validateStatementInterest(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementInterest) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementInterest); i++ {
		if swag.IsZero(m.StatementInterest[i]) { // not required
			continue
		}

		if m.StatementInterest[i] != nil {
			if err := m.StatementInterest[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementInterest" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementRate(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementRate) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementRate); i++ {
		if swag.IsZero(m.StatementRate[i]) { // not required
			continue
		}

		if m.StatementRate[i] != nil {
			if err := m.StatementRate[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementRate" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateStatementReference(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementReference) { // not required
		return nil
	}

	if err := m.StatementReference.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StatementReference")
		}
		return err
	}

	return nil
}

func (m *OBStatement2Detail) validateStatementValue(formats strfmt.Registry) error {

	if swag.IsZero(m.StatementValue) { // not required
		return nil
	}

	for i := 0; i < len(m.StatementValue); i++ {
		if swag.IsZero(m.StatementValue[i]) { // not required
			continue
		}

		if m.StatementValue[i] != nil {
			if err := m.StatementValue[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatementValue" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *OBStatement2Detail) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2Detail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2Detail) UnmarshalBinary(b []byte) error {
	var res OBStatement2Detail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementAmountItems0 Set of elements used to provide details of a generic amount for the statement resource.
//
// swagger:model OBStatement2DetailStatementAmountItems0
type OBStatement2DetailStatementAmountItems0 struct {

	// amount
	// Required: true
	Amount *OBActiveOrHistoricCurrencyAndAmount8 `json:"Amount"`

	// credit debit indicator
	// Required: true
	CreditDebitIndicator OBCreditDebitCode0 `json:"CreditDebitIndicator"`

	// type
	// Required: true
	Type OBExternalStatementAmountType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail statement amount items0
func (m *OBStatement2DetailStatementAmountItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditDebitIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementAmountItems0) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStatement2DetailStatementAmountItems0) validateCreditDebitIndicator(formats strfmt.Registry) error {

	if err := m.CreditDebitIndicator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CreditDebitIndicator")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementAmountItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementAmountItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementAmountItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementBenefitItems0 Set of elements used to provide details of a benefit or reward amount for the statement resource.
//
// swagger:model OBStatement2DetailStatementBenefitItems0
type OBStatement2DetailStatementBenefitItems0 struct {

	// amount
	// Required: true
	Amount *OBActiveOrHistoricCurrencyAndAmount5 `json:"Amount"`

	// type
	// Required: true
	Type OBExternalStatementBenefitType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail statement benefit items0
func (m *OBStatement2DetailStatementBenefitItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementBenefitItems0) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStatement2DetailStatementBenefitItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementBenefitItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementBenefitItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementBenefitItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementDateTimeItems0 Set of elements used to provide details of a generic date time for the statement resource.
//
// swagger:model OBStatement2DetailStatementDateTimeItems0
type OBStatement2DetailStatementDateTimeItems0 struct {

	// date time
	// Required: true
	// Format: date-time
	DateTime DateTime `json:"DateTime"`

	// type
	// Required: true
	Type OBExternalStatementDateTimeType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail statement date time items0
func (m *OBStatement2DetailStatementDateTimeItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementDateTimeItems0) validateDateTime(formats strfmt.Registry) error {

	if err := m.DateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("DateTime")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementDateTimeItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementDateTimeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementDateTimeItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementDateTimeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementFeeItems0 Set of elements used to provide details of a fee for the statement resource.
//
// swagger:model OBStatement2DetailStatementFeeItems0
type OBStatement2DetailStatementFeeItems0 struct {

	// amount
	// Required: true
	Amount *OBActiveOrHistoricCurrencyAndAmount6 `json:"Amount"`

	// credit debit indicator
	// Required: true
	CreditDebitIndicator OBCreditDebitCode0 `json:"CreditDebitIndicator"`

	// description
	Description Description1 `json:"Description,omitempty"`

	// frequency
	Frequency OBExternalStatementFeeFrequency1Code `json:"Frequency,omitempty"`

	// rate
	Rate OBRate10 `json:"Rate,omitempty"`

	// rate type
	RateType OBExternalStatementFeeRateType1Code `json:"RateType,omitempty"`

	// type
	// Required: true
	Type OBExternalStatementFeeType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail statement fee items0
func (m *OBStatement2DetailStatementFeeItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditDebitIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateCreditDebitIndicator(formats strfmt.Registry) error {

	if err := m.CreditDebitIndicator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CreditDebitIndicator")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Description")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := m.Frequency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Frequency")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := m.Rate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Rate")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateRateType(formats strfmt.Registry) error {

	if swag.IsZero(m.RateType) { // not required
		return nil
	}

	if err := m.RateType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RateType")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementFeeItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementFeeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementFeeItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementFeeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementInterestItems0 Set of elements used to provide details of a generic interest amount related to the statement resource.
//
// swagger:model OBStatement2DetailStatementInterestItems0
type OBStatement2DetailStatementInterestItems0 struct {

	// amount
	// Required: true
	Amount *OBActiveOrHistoricCurrencyAndAmount7 `json:"Amount"`

	// credit debit indicator
	// Required: true
	CreditDebitIndicator OBCreditDebitCode0 `json:"CreditDebitIndicator"`

	// description
	Description Description2 `json:"Description,omitempty"`

	// frequency
	Frequency OBExternalStatementInterestFrequency1Code `json:"Frequency,omitempty"`

	// rate
	Rate OBRate11 `json:"Rate,omitempty"`

	// rate type
	RateType OBExternalStatementInterestRateType1Code `json:"RateType,omitempty"`

	// type
	// Required: true
	Type OBExternalStatementInterestType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail statement interest items0
func (m *OBStatement2DetailStatementInterestItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditDebitIndicator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFrequency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("Amount", "body", m.Amount); err != nil {
		return err
	}

	if m.Amount != nil {
		if err := m.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Amount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateCreditDebitIndicator(formats strfmt.Registry) error {

	if err := m.CreditDebitIndicator.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("CreditDebitIndicator")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := m.Description.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Description")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateFrequency(formats strfmt.Registry) error {

	if swag.IsZero(m.Frequency) { // not required
		return nil
	}

	if err := m.Frequency.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Frequency")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if err := m.Rate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Rate")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateRateType(formats strfmt.Registry) error {

	if swag.IsZero(m.RateType) { // not required
		return nil
	}

	if err := m.RateType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("RateType")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementInterestItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementInterestItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementInterestItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementInterestItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementRateItems0 Set of elements used to provide details of a generic rate related to the statement resource.
//
// swagger:model OBStatement2DetailStatementRateItems0
type OBStatement2DetailStatementRateItems0 struct {

	// rate
	// Required: true
	Rate Rate `json:"Rate"`

	// type
	// Required: true
	Type OBExternalStatementRateType1Code `json:"Type"`
}

// Validate validates this o b statement2 detail statement rate items0
func (m *OBStatement2DetailStatementRateItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementRateItems0) validateRate(formats strfmt.Registry) error {

	if err := m.Rate.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Rate")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementRateItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementRateItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementRateItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementRateItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// OBStatement2DetailStatementValueItems0 Set of elements used to provide details of a generic number value related to the statement resource.
//
// swagger:model OBStatement2DetailStatementValueItems0
type OBStatement2DetailStatementValueItems0 struct {

	// type
	// Required: true
	Type OBExternalStatementValueType1Code `json:"Type"`

	// value
	// Required: true
	Value Value `json:"Value"`
}

// Validate validates this o b statement2 detail statement value items0
func (m *OBStatement2DetailStatementValueItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStatement2DetailStatementValueItems0) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

func (m *OBStatement2DetailStatementValueItems0) validateValue(formats strfmt.Registry) error {

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStatement2DetailStatementValueItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStatement2DetailStatementValueItems0) UnmarshalBinary(b []byte) error {
	var res OBStatement2DetailStatementValueItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
