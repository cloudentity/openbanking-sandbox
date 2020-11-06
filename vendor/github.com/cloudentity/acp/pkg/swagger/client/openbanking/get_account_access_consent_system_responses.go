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

// GetAccountAccessConsentSystemReader is a Reader for the GetAccountAccessConsentSystem structure.
type GetAccountAccessConsentSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAccessConsentSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountAccessConsentSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAccountAccessConsentSystemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountAccessConsentSystemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountAccessConsentSystemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountAccessConsentSystemOK creates a GetAccountAccessConsentSystemOK with default headers values
func NewGetAccountAccessConsentSystemOK() *GetAccountAccessConsentSystemOK {
	return &GetAccountAccessConsentSystemOK{}
}

/*GetAccountAccessConsentSystemOK handles this case with default header values.

GetAccountAccessConsentResponse
*/
type GetAccountAccessConsentSystemOK struct {
	Payload *models.GetAccountAccessConsentResponse
}

func (o *GetAccountAccessConsentSystemOK) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/account-access-consent/{login}][%d] getAccountAccessConsentSystemOK  %+v", 200, o.Payload)
}

func (o *GetAccountAccessConsentSystemOK) GetPayload() *models.GetAccountAccessConsentResponse {
	return o.Payload
}

func (o *GetAccountAccessConsentSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAccountAccessConsentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAccessConsentSystemUnauthorized creates a GetAccountAccessConsentSystemUnauthorized with default headers values
func NewGetAccountAccessConsentSystemUnauthorized() *GetAccountAccessConsentSystemUnauthorized {
	return &GetAccountAccessConsentSystemUnauthorized{}
}

/*GetAccountAccessConsentSystemUnauthorized handles this case with default header values.

HttpError
*/
type GetAccountAccessConsentSystemUnauthorized struct {
	Payload *models.Error
}

func (o *GetAccountAccessConsentSystemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/account-access-consent/{login}][%d] getAccountAccessConsentSystemUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccountAccessConsentSystemUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountAccessConsentSystemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAccessConsentSystemForbidden creates a GetAccountAccessConsentSystemForbidden with default headers values
func NewGetAccountAccessConsentSystemForbidden() *GetAccountAccessConsentSystemForbidden {
	return &GetAccountAccessConsentSystemForbidden{}
}

/*GetAccountAccessConsentSystemForbidden handles this case with default header values.

HttpError
*/
type GetAccountAccessConsentSystemForbidden struct {
	Payload *models.Error
}

func (o *GetAccountAccessConsentSystemForbidden) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/account-access-consent/{login}][%d] getAccountAccessConsentSystemForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountAccessConsentSystemForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountAccessConsentSystemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAccessConsentSystemNotFound creates a GetAccountAccessConsentSystemNotFound with default headers values
func NewGetAccountAccessConsentSystemNotFound() *GetAccountAccessConsentSystemNotFound {
	return &GetAccountAccessConsentSystemNotFound{}
}

/*GetAccountAccessConsentSystemNotFound handles this case with default header values.

HttpError
*/
type GetAccountAccessConsentSystemNotFound struct {
	Payload *models.Error
}

func (o *GetAccountAccessConsentSystemNotFound) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/account-access-consent/{login}][%d] getAccountAccessConsentSystemNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountAccessConsentSystemNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountAccessConsentSystemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
