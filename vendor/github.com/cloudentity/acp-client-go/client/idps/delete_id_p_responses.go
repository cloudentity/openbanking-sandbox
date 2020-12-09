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

// DeleteIDPReader is a Reader for the DeleteIDP structure.
type DeleteIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteIDPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewDeleteIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIDPNoContent creates a DeleteIDPNoContent with default headers values
func NewDeleteIDPNoContent() *DeleteIDPNoContent {
	return &DeleteIDPNoContent{}
}

/*DeleteIDPNoContent handles this case with default header values.

IDP has been deleted
*/
type DeleteIDPNoContent struct {
}

func (o *DeleteIDPNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/servers/{aid}/idps/{iid}][%d] deleteIdPNoContent ", 204)
}

func (o *DeleteIDPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIDPBadRequest creates a DeleteIDPBadRequest with default headers values
func NewDeleteIDPBadRequest() *DeleteIDPBadRequest {
	return &DeleteIDPBadRequest{}
}

/*DeleteIDPBadRequest handles this case with default header values.

HttpError
*/
type DeleteIDPBadRequest struct {
	Payload *models.Error
}

func (o *DeleteIDPBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/servers/{aid}/idps/{iid}][%d] deleteIdPBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIDPUnauthorized creates a DeleteIDPUnauthorized with default headers values
func NewDeleteIDPUnauthorized() *DeleteIDPUnauthorized {
	return &DeleteIDPUnauthorized{}
}

/*DeleteIDPUnauthorized handles this case with default header values.

HttpError
*/
type DeleteIDPUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteIDPUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/servers/{aid}/idps/{iid}][%d] deleteIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIDPForbidden creates a DeleteIDPForbidden with default headers values
func NewDeleteIDPForbidden() *DeleteIDPForbidden {
	return &DeleteIDPForbidden{}
}

/*DeleteIDPForbidden handles this case with default header values.

HttpError
*/
type DeleteIDPForbidden struct {
	Payload *models.Error
}

func (o *DeleteIDPForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/servers/{aid}/idps/{iid}][%d] deleteIdPForbidden  %+v", 403, o.Payload)
}

func (o *DeleteIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIDPNotFound creates a DeleteIDPNotFound with default headers values
func NewDeleteIDPNotFound() *DeleteIDPNotFound {
	return &DeleteIDPNotFound{}
}

/*DeleteIDPNotFound handles this case with default header values.

HttpError
*/
type DeleteIDPNotFound struct {
	Payload *models.Error
}

func (o *DeleteIDPNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/servers/{aid}/idps/{iid}][%d] deleteIdPNotFound  %+v", 404, o.Payload)
}

func (o *DeleteIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIDPUnprocessableEntity creates a DeleteIDPUnprocessableEntity with default headers values
func NewDeleteIDPUnprocessableEntity() *DeleteIDPUnprocessableEntity {
	return &DeleteIDPUnprocessableEntity{}
}

/*DeleteIDPUnprocessableEntity handles this case with default header values.

HttpError
*/
type DeleteIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *DeleteIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /api/admin/{tid}/servers/{aid}/idps/{iid}][%d] deleteIdPUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
