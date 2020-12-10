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

// UpdateGithubIDPReader is a Reader for the UpdateGithubIDP structure.
type UpdateGithubIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGithubIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGithubIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGithubIDPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateGithubIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGithubIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGithubIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateGithubIDPUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGithubIDPOK creates a UpdateGithubIDPOK with default headers values
func NewUpdateGithubIDPOK() *UpdateGithubIDPOK {
	return &UpdateGithubIDPOK{}
}

/*UpdateGithubIDPOK handles this case with default header values.

GithubIDP
*/
type UpdateGithubIDPOK struct {
	Payload *models.GithubIDP
}

func (o *UpdateGithubIDPOK) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/github/{iid}][%d] updateGithubIdPOK  %+v", 200, o.Payload)
}

func (o *UpdateGithubIDPOK) GetPayload() *models.GithubIDP {
	return o.Payload
}

func (o *UpdateGithubIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GithubIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGithubIDPBadRequest creates a UpdateGithubIDPBadRequest with default headers values
func NewUpdateGithubIDPBadRequest() *UpdateGithubIDPBadRequest {
	return &UpdateGithubIDPBadRequest{}
}

/*UpdateGithubIDPBadRequest handles this case with default header values.

HttpError
*/
type UpdateGithubIDPBadRequest struct {
	Payload *models.Error
}

func (o *UpdateGithubIDPBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/github/{iid}][%d] updateGithubIdPBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateGithubIDPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGithubIDPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGithubIDPUnauthorized creates a UpdateGithubIDPUnauthorized with default headers values
func NewUpdateGithubIDPUnauthorized() *UpdateGithubIDPUnauthorized {
	return &UpdateGithubIDPUnauthorized{}
}

/*UpdateGithubIDPUnauthorized handles this case with default header values.

HttpError
*/
type UpdateGithubIDPUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateGithubIDPUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/github/{iid}][%d] updateGithubIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateGithubIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGithubIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGithubIDPForbidden creates a UpdateGithubIDPForbidden with default headers values
func NewUpdateGithubIDPForbidden() *UpdateGithubIDPForbidden {
	return &UpdateGithubIDPForbidden{}
}

/*UpdateGithubIDPForbidden handles this case with default header values.

HttpError
*/
type UpdateGithubIDPForbidden struct {
	Payload *models.Error
}

func (o *UpdateGithubIDPForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/github/{iid}][%d] updateGithubIdPForbidden  %+v", 403, o.Payload)
}

func (o *UpdateGithubIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGithubIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGithubIDPNotFound creates a UpdateGithubIDPNotFound with default headers values
func NewUpdateGithubIDPNotFound() *UpdateGithubIDPNotFound {
	return &UpdateGithubIDPNotFound{}
}

/*UpdateGithubIDPNotFound handles this case with default header values.

HttpError
*/
type UpdateGithubIDPNotFound struct {
	Payload *models.Error
}

func (o *UpdateGithubIDPNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/github/{iid}][%d] updateGithubIdPNotFound  %+v", 404, o.Payload)
}

func (o *UpdateGithubIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGithubIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGithubIDPUnprocessableEntity creates a UpdateGithubIDPUnprocessableEntity with default headers values
func NewUpdateGithubIDPUnprocessableEntity() *UpdateGithubIDPUnprocessableEntity {
	return &UpdateGithubIDPUnprocessableEntity{}
}

/*UpdateGithubIDPUnprocessableEntity handles this case with default header values.

HttpError
*/
type UpdateGithubIDPUnprocessableEntity struct {
	Payload *models.Error
}

func (o *UpdateGithubIDPUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /api/admin/{tid}/servers/{aid}/idps/github/{iid}][%d] updateGithubIdPUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateGithubIDPUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateGithubIDPUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
