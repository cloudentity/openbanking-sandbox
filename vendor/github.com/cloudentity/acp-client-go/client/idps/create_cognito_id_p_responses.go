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

// CreateCognitoIDPReader is a Reader for the CreateCognitoIDP structure.
type CreateCognitoIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCognitoIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCognitoIDPCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCognitoIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCognitoIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCognitoIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCognitoIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateCognitoIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCognitoIDPCreated creates a CreateCognitoIDPCreated with default headers values
func NewCreateCognitoIDPCreated() *CreateCognitoIDPCreated {
	return &CreateCognitoIDPCreated{}
}

/*CreateCognitoIDPCreated handles this case with default header values.

CognitoIDP
*/
type CreateCognitoIDPCreated struct {
	Payload *models.CognitoIDP
}

func (o *CreateCognitoIDPCreated) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/cognito][%d] createCognitoIdPCreated  %+v", 201, o.Payload)
}

func (o *CreateCognitoIDPCreated) GetPayload() *models.CognitoIDP {
	return o.Payload
}

func (o *CreateCognitoIDPCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CognitoIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCognitoIDPBadRequest creates a CreateCognitoIDPBadRequest with default headers values
func NewCreateCognitoIDPBadRequest() *CreateCognitoIDPBadRequest {
	return &CreateCognitoIDPBadRequest{}
}

/*CreateCognitoIDPBadRequest handles this case with default header values.

HttpError
*/
type CreateCognitoIDPBadRequest struct {
	Payload *models.Error
}

func (o *CreateCognitoIDPBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/cognito][%d] createCognitoIdPBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCognitoIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCognitoIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCognitoIDPUnauthorized creates a CreateCognitoIDPUnauthorized with default headers values
func NewCreateCognitoIDPUnauthorized() *CreateCognitoIDPUnauthorized {
	return &CreateCognitoIDPUnauthorized{}
}

/*CreateCognitoIDPUnauthorized handles this case with default header values.

HttpError
*/
type CreateCognitoIDPUnauthorized struct {
	Payload *models.Error
}

func (o *CreateCognitoIDPUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/cognito][%d] createCognitoIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCognitoIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCognitoIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCognitoIDPForbidden creates a CreateCognitoIDPForbidden with default headers values
func NewCreateCognitoIDPForbidden() *CreateCognitoIDPForbidden {
	return &CreateCognitoIDPForbidden{}
}

/*CreateCognitoIDPForbidden handles this case with default header values.

HttpError
*/
type CreateCognitoIDPForbidden struct {
	Payload *models.Error
}

func (o *CreateCognitoIDPForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/cognito][%d] createCognitoIdPForbidden  %+v", 403, o.Payload)
}

func (o *CreateCognitoIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCognitoIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCognitoIDPNotFound creates a CreateCognitoIDPNotFound with default headers values
func NewCreateCognitoIDPNotFound() *CreateCognitoIDPNotFound {
	return &CreateCognitoIDPNotFound{}
}

/*CreateCognitoIDPNotFound handles this case with default header values.

HttpError
*/
type CreateCognitoIDPNotFound struct {
	Payload *models.Error
}

func (o *CreateCognitoIDPNotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/cognito][%d] createCognitoIdPNotFound  %+v", 404, o.Payload)
}

func (o *CreateCognitoIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCognitoIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCognitoIDPUnprocessableEntity creates a CreateCognitoIDPUnprocessableEntity with default headers values
func NewCreateCognitoIDPUnprocessableEntity() *CreateCognitoIDPUnprocessableEntity {
	return &CreateCognitoIDPUnprocessableEntity{}
}

/*CreateCognitoIDPUnprocessableEntity handles this case with default header values.

HttpError
*/
type CreateCognitoIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *CreateCognitoIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/servers/{aid}/idps/cognito][%d] createCognitoIdPUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateCognitoIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCognitoIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
