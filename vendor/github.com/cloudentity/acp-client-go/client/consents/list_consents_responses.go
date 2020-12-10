// Code generated by go-swagger; DO NOT EDIT.

package consents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// ListConsentsReader is a Reader for the ListConsents structure.
type ListConsentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListConsentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListConsentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListConsentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListConsentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListConsentsOK creates a ListConsentsOK with default headers values
func NewListConsentsOK() *ListConsentsOK {
	return &ListConsentsOK{}
}

/*ListConsentsOK handles this case with default header values.

Consents
*/
type ListConsentsOK struct {
	Payload *models.Consents
}

func (o *ListConsentsOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/consents][%d] listConsentsOK  %+v", 200, o.Payload)
}

func (o *ListConsentsOK) GetPayload() *models.Consents {
	return o.Payload
}

func (o *ListConsentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Consents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsentsUnauthorized creates a ListConsentsUnauthorized with default headers values
func NewListConsentsUnauthorized() *ListConsentsUnauthorized {
	return &ListConsentsUnauthorized{}
}

/*ListConsentsUnauthorized handles this case with default header values.

HttpError
*/
type ListConsentsUnauthorized struct {
	Payload *models.Error
}

func (o *ListConsentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/consents][%d] listConsentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListConsentsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListConsentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListConsentsForbidden creates a ListConsentsForbidden with default headers values
func NewListConsentsForbidden() *ListConsentsForbidden {
	return &ListConsentsForbidden{}
}

/*ListConsentsForbidden handles this case with default header values.

HttpError
*/
type ListConsentsForbidden struct {
	Payload *models.Error
}

func (o *ListConsentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/consents][%d] listConsentsForbidden  %+v", 403, o.Payload)
}

func (o *ListConsentsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListConsentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
