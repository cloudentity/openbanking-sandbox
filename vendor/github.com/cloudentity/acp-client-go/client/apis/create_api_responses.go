// Code generated by go-swagger; DO NOT EDIT.

package apis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// CreateAPIReader is a Reader for the CreateAPI structure.
type CreateAPIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAPICreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAPIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAPIUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAPIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAPINotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateAPIUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAPICreated creates a CreateAPICreated with default headers values
func NewCreateAPICreated() *CreateAPICreated {
	return &CreateAPICreated{}
}

/*CreateAPICreated handles this case with default header values.

API
*/
type CreateAPICreated struct {
	Payload *models.API
}

func (o *CreateAPICreated) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/apis][%d] createApiCreated  %+v", 201, o.Payload)
}

func (o *CreateAPICreated) GetPayload() *models.API {
	return o.Payload
}

func (o *CreateAPICreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.API)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIBadRequest creates a CreateAPIBadRequest with default headers values
func NewCreateAPIBadRequest() *CreateAPIBadRequest {
	return &CreateAPIBadRequest{}
}

/*CreateAPIBadRequest handles this case with default header values.

HttpError
*/
type CreateAPIBadRequest struct {
	Payload *models.Error
}

func (o *CreateAPIBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/apis][%d] createApiBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAPIBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIUnauthorized creates a CreateAPIUnauthorized with default headers values
func NewCreateAPIUnauthorized() *CreateAPIUnauthorized {
	return &CreateAPIUnauthorized{}
}

/*CreateAPIUnauthorized handles this case with default header values.

HttpError
*/
type CreateAPIUnauthorized struct {
	Payload *models.Error
}

func (o *CreateAPIUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/apis][%d] createApiUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAPIUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPIUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIForbidden creates a CreateAPIForbidden with default headers values
func NewCreateAPIForbidden() *CreateAPIForbidden {
	return &CreateAPIForbidden{}
}

/*CreateAPIForbidden handles this case with default header values.

HttpError
*/
type CreateAPIForbidden struct {
	Payload *models.Error
}

func (o *CreateAPIForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/apis][%d] createApiForbidden  %+v", 403, o.Payload)
}

func (o *CreateAPIForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPINotFound creates a CreateAPINotFound with default headers values
func NewCreateAPINotFound() *CreateAPINotFound {
	return &CreateAPINotFound{}
}

/*CreateAPINotFound handles this case with default header values.

HttpError
*/
type CreateAPINotFound struct {
	Payload *models.Error
}

func (o *CreateAPINotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/apis][%d] createApiNotFound  %+v", 404, o.Payload)
}

func (o *CreateAPINotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPINotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIUnprocessableEntity creates a CreateAPIUnprocessableEntity with default headers values
func NewCreateAPIUnprocessableEntity() *CreateAPIUnprocessableEntity {
	return &CreateAPIUnprocessableEntity{}
}

/*CreateAPIUnprocessableEntity handles this case with default header values.

HttpError
*/
type CreateAPIUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateAPIUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/apis][%d] createApiUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateAPIUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAPIUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
