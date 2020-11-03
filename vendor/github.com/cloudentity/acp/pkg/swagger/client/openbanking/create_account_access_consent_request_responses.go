// Code generated by go-swagger; DO NOT EDIT.

package openbanking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp/pkg/swagger/models"
)

// CreateAccountAccessConsentRequestReader is a Reader for the CreateAccountAccessConsentRequest structure.
type CreateAccountAccessConsentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountAccessConsentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAccountAccessConsentRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAccountAccessConsentRequestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAccountAccessConsentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAccountAccessConsentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateAccountAccessConsentRequestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateAccountAccessConsentRequestNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateAccountAccessConsentRequestUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateAccountAccessConsentRequestTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAccountAccessConsentRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountAccessConsentRequestCreated creates a CreateAccountAccessConsentRequestCreated with default headers values
func NewCreateAccountAccessConsentRequestCreated() *CreateAccountAccessConsentRequestCreated {
	return &CreateAccountAccessConsentRequestCreated{}
}

/*CreateAccountAccessConsentRequestCreated handles this case with default header values.

AccountAccessConsentResponse
*/
type CreateAccountAccessConsentRequestCreated struct {
	Payload *models.AccountAccessConsentResponse
}

func (o *CreateAccountAccessConsentRequestCreated) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateAccountAccessConsentRequestCreated) GetPayload() *models.AccountAccessConsentResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestBadRequest creates a CreateAccountAccessConsentRequestBadRequest with default headers values
func NewCreateAccountAccessConsentRequestBadRequest() *CreateAccountAccessConsentRequestBadRequest {
	return &CreateAccountAccessConsentRequestBadRequest{}
}

/*CreateAccountAccessConsentRequestBadRequest handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestBadRequest) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAccountAccessConsentRequestBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestUnauthorized creates a CreateAccountAccessConsentRequestUnauthorized with default headers values
func NewCreateAccountAccessConsentRequestUnauthorized() *CreateAccountAccessConsentRequestUnauthorized {
	return &CreateAccountAccessConsentRequestUnauthorized{}
}

/*CreateAccountAccessConsentRequestUnauthorized handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestUnauthorized struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAccountAccessConsentRequestUnauthorized) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestForbidden creates a CreateAccountAccessConsentRequestForbidden with default headers values
func NewCreateAccountAccessConsentRequestForbidden() *CreateAccountAccessConsentRequestForbidden {
	return &CreateAccountAccessConsentRequestForbidden{}
}

/*CreateAccountAccessConsentRequestForbidden handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestForbidden struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestForbidden  %+v", 403, o.Payload)
}

func (o *CreateAccountAccessConsentRequestForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestMethodNotAllowed creates a CreateAccountAccessConsentRequestMethodNotAllowed with default headers values
func NewCreateAccountAccessConsentRequestMethodNotAllowed() *CreateAccountAccessConsentRequestMethodNotAllowed {
	return &CreateAccountAccessConsentRequestMethodNotAllowed{}
}

/*CreateAccountAccessConsentRequestMethodNotAllowed handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestMethodNotAllowed struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateAccountAccessConsentRequestMethodNotAllowed) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestNotAcceptable creates a CreateAccountAccessConsentRequestNotAcceptable with default headers values
func NewCreateAccountAccessConsentRequestNotAcceptable() *CreateAccountAccessConsentRequestNotAcceptable {
	return &CreateAccountAccessConsentRequestNotAcceptable{}
}

/*CreateAccountAccessConsentRequestNotAcceptable handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestNotAcceptable struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestNotAcceptable  %+v", 406, o.Payload)
}

func (o *CreateAccountAccessConsentRequestNotAcceptable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestUnsupportedMediaType creates a CreateAccountAccessConsentRequestUnsupportedMediaType with default headers values
func NewCreateAccountAccessConsentRequestUnsupportedMediaType() *CreateAccountAccessConsentRequestUnsupportedMediaType {
	return &CreateAccountAccessConsentRequestUnsupportedMediaType{}
}

/*CreateAccountAccessConsentRequestUnsupportedMediaType handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestUnsupportedMediaType struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateAccountAccessConsentRequestUnsupportedMediaType) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestTooManyRequests creates a CreateAccountAccessConsentRequestTooManyRequests with default headers values
func NewCreateAccountAccessConsentRequestTooManyRequests() *CreateAccountAccessConsentRequestTooManyRequests {
	return &CreateAccountAccessConsentRequestTooManyRequests{}
}

/*CreateAccountAccessConsentRequestTooManyRequests handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestTooManyRequests struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateAccountAccessConsentRequestTooManyRequests) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountAccessConsentRequestInternalServerError creates a CreateAccountAccessConsentRequestInternalServerError with default headers values
func NewCreateAccountAccessConsentRequestInternalServerError() *CreateAccountAccessConsentRequestInternalServerError {
	return &CreateAccountAccessConsentRequestInternalServerError{}
}

/*CreateAccountAccessConsentRequestInternalServerError handles this case with default header values.

ErrorResponse
*/
type CreateAccountAccessConsentRequestInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *CreateAccountAccessConsentRequestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/open-banking/v3.1/aisp/account-access-consents][%d] createAccountAccessConsentRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAccountAccessConsentRequestInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateAccountAccessConsentRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
