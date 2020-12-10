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

// UpdateOktaIDPReader is a Reader for the UpdateOktaIDP structure.
type UpdateOktaIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOktaIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOktaIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateOktaIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateOktaIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateOktaIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateOktaIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateOktaIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOktaIDPOK creates a UpdateOktaIDPOK with default headers values
func NewUpdateOktaIDPOK() *UpdateOktaIDPOK {
	return &UpdateOktaIDPOK{}
}

/*UpdateOktaIDPOK handles this case with default header values.

OktaIDP
*/
type UpdateOktaIDPOK struct {
	Payload *models.OktaIDP
}

func (o *UpdateOktaIDPOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] updateOktaIdPOK  %+v", 200, o.Payload)
}

func (o *UpdateOktaIDPOK) GetPayload() *models.OktaIDP {
	return o.Payload
}

func (o *UpdateOktaIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OktaIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOktaIDPBadRequest creates a UpdateOktaIDPBadRequest with default headers values
func NewUpdateOktaIDPBadRequest() *UpdateOktaIDPBadRequest {
	return &UpdateOktaIDPBadRequest{}
}

/*UpdateOktaIDPBadRequest handles this case with default header values.

HttpError
*/
type UpdateOktaIDPBadRequest struct {
	Payload *models.Error
}

func (o *UpdateOktaIDPBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] updateOktaIdPBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateOktaIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOktaIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOktaIDPUnauthorized creates a UpdateOktaIDPUnauthorized with default headers values
func NewUpdateOktaIDPUnauthorized() *UpdateOktaIDPUnauthorized {
	return &UpdateOktaIDPUnauthorized{}
}

/*UpdateOktaIDPUnauthorized handles this case with default header values.

HttpError
*/
type UpdateOktaIDPUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateOktaIDPUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] updateOktaIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateOktaIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOktaIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOktaIDPForbidden creates a UpdateOktaIDPForbidden with default headers values
func NewUpdateOktaIDPForbidden() *UpdateOktaIDPForbidden {
	return &UpdateOktaIDPForbidden{}
}

/*UpdateOktaIDPForbidden handles this case with default header values.

HttpError
*/
type UpdateOktaIDPForbidden struct {
	Payload *models.Error
}

func (o *UpdateOktaIDPForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] updateOktaIdPForbidden  %+v", 403, o.Payload)
}

func (o *UpdateOktaIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOktaIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOktaIDPNotFound creates a UpdateOktaIDPNotFound with default headers values
func NewUpdateOktaIDPNotFound() *UpdateOktaIDPNotFound {
	return &UpdateOktaIDPNotFound{}
}

/*UpdateOktaIDPNotFound handles this case with default header values.

HttpError
*/
type UpdateOktaIDPNotFound struct {
	Payload *models.Error
}

func (o *UpdateOktaIDPNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] updateOktaIdPNotFound  %+v", 404, o.Payload)
}

func (o *UpdateOktaIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOktaIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOktaIDPUnprocessableEntity creates a UpdateOktaIDPUnprocessableEntity with default headers values
func NewUpdateOktaIDPUnprocessableEntity() *UpdateOktaIDPUnprocessableEntity {
	return &UpdateOktaIDPUnprocessableEntity{}
}

/*UpdateOktaIDPUnprocessableEntity handles this case with default header values.

HttpError
*/
type UpdateOktaIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateOktaIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] updateOktaIdPUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateOktaIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateOktaIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
