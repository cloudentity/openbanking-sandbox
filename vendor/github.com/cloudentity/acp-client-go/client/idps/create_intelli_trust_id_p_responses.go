// Code generated by go-swagger; DO NOT EDIT.

package idps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// CreateIntelliTrustIDPReader is a Reader for the CreateIntelliTrustIDP structure.
type CreateIntelliTrustIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIntelliTrustIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateIntelliTrustIDPCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIntelliTrustIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateIntelliTrustIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateIntelliTrustIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateIntelliTrustIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateIntelliTrustIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateIntelliTrustIDPCreated creates a CreateIntelliTrustIDPCreated with default headers values
func NewCreateIntelliTrustIDPCreated() *CreateIntelliTrustIDPCreated {
	return &CreateIntelliTrustIDPCreated{}
}

/*CreateIntelliTrustIDPCreated handles this case with default header values.

IntelliTrustIDP
*/
type CreateIntelliTrustIDPCreated struct {
	Payload *models.IntelliTrustIDP
}

func (o *CreateIntelliTrustIDPCreated) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/intelli_trust][%d] createIntelliTrustIdPCreated  %+v", 201, o.Payload)
}

func (o *CreateIntelliTrustIDPCreated) GetPayload() *models.IntelliTrustIDP {
	return o.Payload
}

func (o *CreateIntelliTrustIDPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntelliTrustIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntelliTrustIDPBadRequest creates a CreateIntelliTrustIDPBadRequest with default headers values
func NewCreateIntelliTrustIDPBadRequest() *CreateIntelliTrustIDPBadRequest {
	return &CreateIntelliTrustIDPBadRequest{}
}

/*CreateIntelliTrustIDPBadRequest handles this case with default header values.

HttpError
*/
type CreateIntelliTrustIDPBadRequest struct {
	Payload *models.Error
}

func (o *CreateIntelliTrustIDPBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/intelli_trust][%d] createIntelliTrustIdPBadRequest  %+v", 400, o.Payload)
}

func (o *CreateIntelliTrustIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntelliTrustIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntelliTrustIDPUnauthorized creates a CreateIntelliTrustIDPUnauthorized with default headers values
func NewCreateIntelliTrustIDPUnauthorized() *CreateIntelliTrustIDPUnauthorized {
	return &CreateIntelliTrustIDPUnauthorized{}
}

/*CreateIntelliTrustIDPUnauthorized handles this case with default header values.

HttpError
*/
type CreateIntelliTrustIDPUnauthorized struct {
	Payload *models.Error
}

func (o *CreateIntelliTrustIDPUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/intelli_trust][%d] createIntelliTrustIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateIntelliTrustIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntelliTrustIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntelliTrustIDPForbidden creates a CreateIntelliTrustIDPForbidden with default headers values
func NewCreateIntelliTrustIDPForbidden() *CreateIntelliTrustIDPForbidden {
	return &CreateIntelliTrustIDPForbidden{}
}

/*CreateIntelliTrustIDPForbidden handles this case with default header values.

HttpError
*/
type CreateIntelliTrustIDPForbidden struct {
	Payload *models.Error
}

func (o *CreateIntelliTrustIDPForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/intelli_trust][%d] createIntelliTrustIdPForbidden  %+v", 403, o.Payload)
}

func (o *CreateIntelliTrustIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntelliTrustIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntelliTrustIDPNotFound creates a CreateIntelliTrustIDPNotFound with default headers values
func NewCreateIntelliTrustIDPNotFound() *CreateIntelliTrustIDPNotFound {
	return &CreateIntelliTrustIDPNotFound{}
}

/*CreateIntelliTrustIDPNotFound handles this case with default header values.

HttpError
*/
type CreateIntelliTrustIDPNotFound struct {
	Payload *models.Error
}

func (o *CreateIntelliTrustIDPNotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/intelli_trust][%d] createIntelliTrustIdPNotFound  %+v", 404, o.Payload)
}

func (o *CreateIntelliTrustIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntelliTrustIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIntelliTrustIDPUnprocessableEntity creates a CreateIntelliTrustIDPUnprocessableEntity with default headers values
func NewCreateIntelliTrustIDPUnprocessableEntity() *CreateIntelliTrustIDPUnprocessableEntity {
	return &CreateIntelliTrustIDPUnprocessableEntity{}
}

/*CreateIntelliTrustIDPUnprocessableEntity handles this case with default header values.

HttpError
*/
type CreateIntelliTrustIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateIntelliTrustIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/intelli_trust][%d] createIntelliTrustIdPUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateIntelliTrustIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateIntelliTrustIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
