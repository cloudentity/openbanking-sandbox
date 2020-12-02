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

// ListConsentsByClientIDReader is a Reader for the ListConsentsByClientID structure.
type ListConsentsByClientIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConsentsByClientIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConsentsByClientIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListConsentsByClientIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListConsentsByClientIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListConsentsByClientIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListConsentsByClientIDOK creates a ListConsentsByClientIDOK with default headers values
func NewListConsentsByClientIDOK() *ListConsentsByClientIDOK {
	return &ListConsentsByClientIDOK{}
}

/*ListConsentsByClientIDOK handles this case with default header values.

ListAccountAccessConsents
*/
type ListConsentsByClientIDOK struct {
	Payload *models.ListAccountAccessConsents
}

func (o *ListConsentsByClientIDOK) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/consents/by-client/{clientID}][%d] listConsentsByClientIdOK  %+v", 200, o.Payload)
}

func (o *ListConsentsByClientIDOK) GetPayload() *models.ListAccountAccessConsents {
	return o.Payload
}

func (o *ListConsentsByClientIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListAccountAccessConsents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsentsByClientIDUnauthorized creates a ListConsentsByClientIDUnauthorized with default headers values
func NewListConsentsByClientIDUnauthorized() *ListConsentsByClientIDUnauthorized {
	return &ListConsentsByClientIDUnauthorized{}
}

/*ListConsentsByClientIDUnauthorized handles this case with default header values.

HttpError
*/
type ListConsentsByClientIDUnauthorized struct {
	Payload *models.Error
}

func (o *ListConsentsByClientIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/consents/by-client/{clientID}][%d] listConsentsByClientIdUnauthorized  %+v", 401, o.Payload)
}

func (o *ListConsentsByClientIDUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListConsentsByClientIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsentsByClientIDForbidden creates a ListConsentsByClientIDForbidden with default headers values
func NewListConsentsByClientIDForbidden() *ListConsentsByClientIDForbidden {
	return &ListConsentsByClientIDForbidden{}
}

/*ListConsentsByClientIDForbidden handles this case with default header values.

HttpError
*/
type ListConsentsByClientIDForbidden struct {
	Payload *models.Error
}

func (o *ListConsentsByClientIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/consents/by-client/{clientID}][%d] listConsentsByClientIdForbidden  %+v", 403, o.Payload)
}

func (o *ListConsentsByClientIDForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListConsentsByClientIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsentsByClientIDNotFound creates a ListConsentsByClientIDNotFound with default headers values
func NewListConsentsByClientIDNotFound() *ListConsentsByClientIDNotFound {
	return &ListConsentsByClientIDNotFound{}
}

/*ListConsentsByClientIDNotFound handles this case with default header values.

HttpError
*/
type ListConsentsByClientIDNotFound struct {
	Payload *models.Error
}

func (o *ListConsentsByClientIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/system/{tid}/open-banking/consents/by-client/{clientID}][%d] listConsentsByClientIdNotFound  %+v", 404, o.Payload)
}

func (o *ListConsentsByClientIDNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListConsentsByClientIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
