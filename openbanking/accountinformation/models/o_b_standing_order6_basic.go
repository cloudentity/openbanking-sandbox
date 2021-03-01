// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBStandingOrder6Basic o b standing order6 basic
//
// swagger:model OBStandingOrder6Basic
type OBStandingOrder6Basic struct {

	// account Id
	// Required: true
	AccountID *AccountID `json:"AccountId"`

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
	Frequency *Frequency1 `json:"Frequency"`

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

// Validate validates this o b standing order6 basic
func (m *OBStandingOrder6Basic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
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

func (m *OBStandingOrder6Basic) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("AccountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.Required("AccountId", "body", m.AccountID); err != nil {
		return err
	}

	if m.AccountID != nil {
		if err := m.AccountID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountId")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) validateFinalPaymentAmount(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateFinalPaymentDateTime(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateFirstPaymentAmount(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateFirstPaymentDateTime(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateFrequency(formats strfmt.Registry) error {

	if err := validate.Required("Frequency", "body", m.Frequency); err != nil {
		return err
	}

	if err := validate.Required("Frequency", "body", m.Frequency); err != nil {
		return err
	}

	if m.Frequency != nil {
		if err := m.Frequency.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Frequency")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) validateLastPaymentAmount(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateLastPaymentDateTime(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateNextPaymentAmount(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateNextPaymentDateTime(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateNumberOfPayments(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateReference(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateStandingOrderID(formats strfmt.Registry) error {
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

func (m *OBStandingOrder6Basic) validateStandingOrderStatusCode(formats strfmt.Registry) error {
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

// ContextValidate validate this o b standing order6 basic based on the context it is used
func (m *OBStandingOrder6Basic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFirstPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFrequency(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextPaymentAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNextPaymentDateTime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNumberOfPayments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandingOrderID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStandingOrderStatusCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBStandingOrder6Basic) contextValidateAccountID(ctx context.Context, formats strfmt.Registry) error {

	if m.AccountID != nil {
		if err := m.AccountID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("AccountId")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateFinalPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.FinalPaymentAmount != nil {
		if err := m.FinalPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FinalPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateFinalPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FinalPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FinalPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateFirstPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.FirstPaymentAmount != nil {
		if err := m.FirstPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("FirstPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateFirstPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.FirstPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FirstPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateFrequency(ctx context.Context, formats strfmt.Registry) error {

	if m.Frequency != nil {
		if err := m.Frequency.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Frequency")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateLastPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.LastPaymentAmount != nil {
		if err := m.LastPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LastPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateLastPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LastPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LastPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateNextPaymentAmount(ctx context.Context, formats strfmt.Registry) error {

	if m.NextPaymentAmount != nil {
		if err := m.NextPaymentAmount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NextPaymentAmount")
			}
			return err
		}
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateNextPaymentDateTime(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NextPaymentDateTime.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NextPaymentDateTime")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateNumberOfPayments(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NumberOfPayments.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("NumberOfPayments")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateReference(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Reference.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateStandingOrderID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StandingOrderID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderId")
		}
		return err
	}

	return nil
}

func (m *OBStandingOrder6Basic) contextValidateStandingOrderStatusCode(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StandingOrderStatusCode.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StandingOrderStatusCode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBStandingOrder6Basic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBStandingOrder6Basic) UnmarshalBinary(b []byte) error {
	var res OBStandingOrder6Basic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
