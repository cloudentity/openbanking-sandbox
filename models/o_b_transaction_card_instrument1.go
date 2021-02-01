// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OBTransactionCardInstrument1 Set of elements to describe the card instrument used in the transaction.
//
// swagger:model OBTransactionCardInstrument1
type OBTransactionCardInstrument1 struct {

	// The card authorisation type.
	// Enum: [ConsumerDevice Contactless None PIN]
	AuthorisationType string `json:"AuthorisationType,omitempty"`

	// Name of the card scheme.
	// Required: true
	// Enum: [AmericanExpress Diners Discover MasterCard VISA]
	CardSchemeName *string `json:"CardSchemeName"`

	// Identification assigned by an institution to identify the card instrument used in the transaction. This identification is known by the account owner, and may be masked.
	// Max Length: 34
	// Min Length: 1
	Identification string `json:"Identification,omitempty"`

	// Name of the cardholder using the card instrument.
	// Max Length: 70
	// Min Length: 1
	Name string `json:"Name,omitempty"`
}

// Validate validates this o b transaction card instrument1
func (m *OBTransactionCardInstrument1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorisationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardSchemeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var oBTransactionCardInstrument1TypeAuthorisationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ConsumerDevice","Contactless","None","PIN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBTransactionCardInstrument1TypeAuthorisationTypePropEnum = append(oBTransactionCardInstrument1TypeAuthorisationTypePropEnum, v)
	}
}

const (

	// OBTransactionCardInstrument1AuthorisationTypeConsumerDevice captures enum value "ConsumerDevice"
	OBTransactionCardInstrument1AuthorisationTypeConsumerDevice string = "ConsumerDevice"

	// OBTransactionCardInstrument1AuthorisationTypeContactless captures enum value "Contactless"
	OBTransactionCardInstrument1AuthorisationTypeContactless string = "Contactless"

	// OBTransactionCardInstrument1AuthorisationTypeNone captures enum value "None"
	OBTransactionCardInstrument1AuthorisationTypeNone string = "None"

	// OBTransactionCardInstrument1AuthorisationTypePIN captures enum value "PIN"
	OBTransactionCardInstrument1AuthorisationTypePIN string = "PIN"
)

// prop value enum
func (m *OBTransactionCardInstrument1) validateAuthorisationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oBTransactionCardInstrument1TypeAuthorisationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OBTransactionCardInstrument1) validateAuthorisationType(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthorisationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthorisationTypeEnum("AuthorisationType", "body", m.AuthorisationType); err != nil {
		return err
	}

	return nil
}

var oBTransactionCardInstrument1TypeCardSchemeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AmericanExpress","Diners","Discover","MasterCard","VISA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		oBTransactionCardInstrument1TypeCardSchemeNamePropEnum = append(oBTransactionCardInstrument1TypeCardSchemeNamePropEnum, v)
	}
}

const (

	// OBTransactionCardInstrument1CardSchemeNameAmericanExpress captures enum value "AmericanExpress"
	OBTransactionCardInstrument1CardSchemeNameAmericanExpress string = "AmericanExpress"

	// OBTransactionCardInstrument1CardSchemeNameDiners captures enum value "Diners"
	OBTransactionCardInstrument1CardSchemeNameDiners string = "Diners"

	// OBTransactionCardInstrument1CardSchemeNameDiscover captures enum value "Discover"
	OBTransactionCardInstrument1CardSchemeNameDiscover string = "Discover"

	// OBTransactionCardInstrument1CardSchemeNameMasterCard captures enum value "MasterCard"
	OBTransactionCardInstrument1CardSchemeNameMasterCard string = "MasterCard"

	// OBTransactionCardInstrument1CardSchemeNameVISA captures enum value "VISA"
	OBTransactionCardInstrument1CardSchemeNameVISA string = "VISA"
)

// prop value enum
func (m *OBTransactionCardInstrument1) validateCardSchemeNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, oBTransactionCardInstrument1TypeCardSchemeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OBTransactionCardInstrument1) validateCardSchemeName(formats strfmt.Registry) error {

	if err := validate.Required("CardSchemeName", "body", m.CardSchemeName); err != nil {
		return err
	}

	// value enum
	if err := m.validateCardSchemeNameEnum("CardSchemeName", "body", *m.CardSchemeName); err != nil {
		return err
	}

	return nil
}

func (m *OBTransactionCardInstrument1) validateIdentification(formats strfmt.Registry) error {
	if swag.IsZero(m.Identification) { // not required
		return nil
	}

	if err := validate.MinLength("Identification", "body", m.Identification, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Identification", "body", m.Identification, 34); err != nil {
		return err
	}

	return nil
}

func (m *OBTransactionCardInstrument1) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("Name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("Name", "body", m.Name, 70); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this o b transaction card instrument1 based on context it is used
func (m *OBTransactionCardInstrument1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OBTransactionCardInstrument1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OBTransactionCardInstrument1) UnmarshalBinary(b []byte) error {
	var res OBTransactionCardInstrument1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
