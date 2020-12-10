// Code generated by go-swagger; DO NOT EDIT.

package openbanking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// RevokeOpenbankingConsentReader is a Reader for the RevokeOpenbankingConsent structure.
type RevokeOpenbankingConsentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeOpenbankingConsentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRevokeOpenbankingConsentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeOpenbankingConsentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeOpenbankingConsentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeOpenbankingConsentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeOpenbankingConsentNoContent creates a RevokeOpenbankingConsentNoContent with default headers values
func NewRevokeOpenbankingConsentNoContent() *RevokeOpenbankingConsentNoContent {
	return &RevokeOpenbankingConsentNoContent{}
}

/*RevokeOpenbankingConsentNoContent handles this case with default header values.

Consent has been revoked
*/
type RevokeOpenbankingConsentNoContent struct {
}

func (o *RevokeOpenbankingConsentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/system/{tid}/open-banking/consents/{consentID}][%d] revokeOpenbankingConsentNoContent ", 204)
}

func (o *RevokeOpenbankingConsentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeOpenbankingConsentUnauthorized creates a RevokeOpenbankingConsentUnauthorized with default headers values
func NewRevokeOpenbankingConsentUnauthorized() *RevokeOpenbankingConsentUnauthorized {
	return &RevokeOpenbankingConsentUnauthorized{}
}

/*RevokeOpenbankingConsentUnauthorized handles this case with default header values.

HttpError
*/
type RevokeOpenbankingConsentUnauthorized struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/system/{tid}/open-banking/consents/{consentID}][%d] revokeOpenbankingConsentUnauthorized  %+v", 401, o.Payload)
}

func (o *RevokeOpenbankingConsentUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentForbidden creates a RevokeOpenbankingConsentForbidden with default headers values
func NewRevokeOpenbankingConsentForbidden() *RevokeOpenbankingConsentForbidden {
	return &RevokeOpenbankingConsentForbidden{}
}

/*RevokeOpenbankingConsentForbidden handles this case with default header values.

HttpError
*/
type RevokeOpenbankingConsentForbidden struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/system/{tid}/open-banking/consents/{consentID}][%d] revokeOpenbankingConsentForbidden  %+v", 403, o.Payload)
}

func (o *RevokeOpenbankingConsentForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeOpenbankingConsentNotFound creates a RevokeOpenbankingConsentNotFound with default headers values
func NewRevokeOpenbankingConsentNotFound() *RevokeOpenbankingConsentNotFound {
	return &RevokeOpenbankingConsentNotFound{}
}

/*RevokeOpenbankingConsentNotFound handles this case with default header values.

HttpError
*/
type RevokeOpenbankingConsentNotFound struct {
	Payload *models.Error
}

func (o *RevokeOpenbankingConsentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/system/{tid}/open-banking/consents/{consentID}][%d] revokeOpenbankingConsentNotFound  %+v", 404, o.Payload)
}

func (o *RevokeOpenbankingConsentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeOpenbankingConsentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
