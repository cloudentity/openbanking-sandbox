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

// UpdateAzureIDPReader is a Reader for the UpdateAzureIDP structure.
type UpdateAzureIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAzureIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAzureIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAzureIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateAzureIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAzureIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAzureIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateAzureIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAzureIDPOK creates a UpdateAzureIDPOK with default headers values
func NewUpdateAzureIDPOK() *UpdateAzureIDPOK {
	return &UpdateAzureIDPOK{}
}

/*UpdateAzureIDPOK handles this case with default header values.

AzureIDP
*/
type UpdateAzureIDPOK struct {
	Payload *models.AzureIDP
}

func (o *UpdateAzureIDPOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/azure/{iid}][%d] updateAzureIdPOK  %+v", 200, o.Payload)
}

func (o *UpdateAzureIDPOK) GetPayload() *models.AzureIDP {
	return o.Payload
}

func (o *UpdateAzureIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureIDPBadRequest creates a UpdateAzureIDPBadRequest with default headers values
func NewUpdateAzureIDPBadRequest() *UpdateAzureIDPBadRequest {
	return &UpdateAzureIDPBadRequest{}
}

/*UpdateAzureIDPBadRequest handles this case with default header values.

HttpError
*/
type UpdateAzureIDPBadRequest struct {
	Payload *models.Error
}

func (o *UpdateAzureIDPBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/azure/{iid}][%d] updateAzureIdPBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateAzureIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAzureIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureIDPUnauthorized creates a UpdateAzureIDPUnauthorized with default headers values
func NewUpdateAzureIDPUnauthorized() *UpdateAzureIDPUnauthorized {
	return &UpdateAzureIDPUnauthorized{}
}

/*UpdateAzureIDPUnauthorized handles this case with default header values.

HttpError
*/
type UpdateAzureIDPUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateAzureIDPUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/azure/{iid}][%d] updateAzureIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateAzureIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAzureIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureIDPForbidden creates a UpdateAzureIDPForbidden with default headers values
func NewUpdateAzureIDPForbidden() *UpdateAzureIDPForbidden {
	return &UpdateAzureIDPForbidden{}
}

/*UpdateAzureIDPForbidden handles this case with default header values.

HttpError
*/
type UpdateAzureIDPForbidden struct {
	Payload *models.Error
}

func (o *UpdateAzureIDPForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/azure/{iid}][%d] updateAzureIdPForbidden  %+v", 403, o.Payload)
}

func (o *UpdateAzureIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAzureIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureIDPNotFound creates a UpdateAzureIDPNotFound with default headers values
func NewUpdateAzureIDPNotFound() *UpdateAzureIDPNotFound {
	return &UpdateAzureIDPNotFound{}
}

/*UpdateAzureIDPNotFound handles this case with default header values.

HttpError
*/
type UpdateAzureIDPNotFound struct {
	Payload *models.Error
}

func (o *UpdateAzureIDPNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/azure/{iid}][%d] updateAzureIdPNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAzureIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAzureIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureIDPUnprocessableEntity creates a UpdateAzureIDPUnprocessableEntity with default headers values
func NewUpdateAzureIDPUnprocessableEntity() *UpdateAzureIDPUnprocessableEntity {
	return &UpdateAzureIDPUnprocessableEntity{}
}

/*UpdateAzureIDPUnprocessableEntity handles this case with default header values.

HttpError
*/
type UpdateAzureIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateAzureIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/azure/{iid}][%d] updateAzureIdPUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateAzureIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAzureIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}