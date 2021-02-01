// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OBBeneficiary5Basic o b beneficiary5 basic
//
// swagger:model OBBeneficiary5Basic
type OBBeneficiary5Basic struct {

	// account Id
	AccountID AccountID `json:"AccountId,omitempty"`

	// beneficiary Id
	BeneficiaryID BeneficiaryID `json:"BeneficiaryId,omitempty"`

	// beneficiary type
	BeneficiaryType OBBeneficiaryType1Code `json:"BeneficiaryType,omitempty"`

	// reference
	Reference Reference `json:"Reference,omitempty"`

	// supplementary data
	SupplementaryData OBSupplementaryData1 `json:"SupplementaryData,omitempty"`
}

// Validate validates this o b beneficiary5 basic
func (m *OBBeneficiary5Basic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBBeneficiary5Basic) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := m.AccountID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Basic) validateBeneficiaryID(formats strfmt.Registry) error {
	if swag.IsZero(m.BeneficiaryID) { // not required
		return nil
	}

	if err := m.BeneficiaryID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Basic) validateBeneficiaryType(formats strfmt.Registry) error {
	if swag.IsZero(m.BeneficiaryType) { // not required
		return nil
	}

	if err := m.BeneficiaryType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryType")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Basic) validateReference(formats strfmt.Registry) error {
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

// ContextValidate validate this o b beneficiary5 basic based on the context it is used
func (m *OBBeneficiary5Basic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBeneficiaryID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBeneficiaryType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OBBeneficiary5Basic) contextValidateAccountID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.AccountID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AccountId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Basic) contextValidateBeneficiaryID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BeneficiaryID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryId")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Basic) contextValidateBeneficiaryType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.BeneficiaryType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("BeneficiaryType")
		}
		return err
	}

	return nil
}

func (m *OBBeneficiary5Basic) contextValidateReference(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Reference.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Reference")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OBBeneficiary5Basic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBBeneficiary5Basic) UnmarshalBinary(b []byte) error {
	var res OBBeneficiary5Basic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
