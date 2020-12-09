// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// ImportSpecificationFromFileReader is a Reader for the ImportSpecificationFromFile structure.
type ImportSpecificationFromFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportSpecificationFromFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportSpecificationFromFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportSpecificationFromFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportSpecificationFromFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportSpecificationFromFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportSpecificationFromFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImportSpecificationFromFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewImportSpecificationFromFileUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportSpecificationFromFileOK creates a ImportSpecificationFromFileOK with default headers values
func NewImportSpecificationFromFileOK() *ImportSpecificationFromFileOK {
	return &ImportSpecificationFromFileOK{}
}

/*ImportSpecificationFromFileOK handles this case with default header values.

ImportServiceConfigurationResult
*/
type ImportSpecificationFromFileOK struct {
	Payload *models.ImportServiceConfigurationResult
}

func (o *ImportSpecificationFromFileOK) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileOK  %+v", 200, o.Payload)
}

func (o *ImportSpecificationFromFileOK) GetPayload() *models.ImportServiceConfigurationResult {
	return o.Payload
}

func (o *ImportSpecificationFromFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportServiceConfigurationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromFileBadRequest creates a ImportSpecificationFromFileBadRequest with default headers values
func NewImportSpecificationFromFileBadRequest() *ImportSpecificationFromFileBadRequest {
	return &ImportSpecificationFromFileBadRequest{}
}

/*ImportSpecificationFromFileBadRequest handles this case with default header values.

HttpError
*/
type ImportSpecificationFromFileBadRequest struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileBadRequest  %+v", 400, o.Payload)
}

func (o *ImportSpecificationFromFileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromFileUnauthorized creates a ImportSpecificationFromFileUnauthorized with default headers values
func NewImportSpecificationFromFileUnauthorized() *ImportSpecificationFromFileUnauthorized {
	return &ImportSpecificationFromFileUnauthorized{}
}

/*ImportSpecificationFromFileUnauthorized handles this case with default header values.

HttpError
*/
type ImportSpecificationFromFileUnauthorized struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromFileUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileUnauthorized  %+v", 401, o.Payload)
}

func (o *ImportSpecificationFromFileUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromFileForbidden creates a ImportSpecificationFromFileForbidden with default headers values
func NewImportSpecificationFromFileForbidden() *ImportSpecificationFromFileForbidden {
	return &ImportSpecificationFromFileForbidden{}
}

/*ImportSpecificationFromFileForbidden handles this case with default header values.

HttpError
*/
type ImportSpecificationFromFileForbidden struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromFileForbidden) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileForbidden  %+v", 403, o.Payload)
}

func (o *ImportSpecificationFromFileForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromFileNotFound creates a ImportSpecificationFromFileNotFound with default headers values
func NewImportSpecificationFromFileNotFound() *ImportSpecificationFromFileNotFound {
	return &ImportSpecificationFromFileNotFound{}
}

/*ImportSpecificationFromFileNotFound handles this case with default header values.

HttpError
*/
type ImportSpecificationFromFileNotFound struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromFileNotFound) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileNotFound  %+v", 404, o.Payload)
}

func (o *ImportSpecificationFromFileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromFileConflict creates a ImportSpecificationFromFileConflict with default headers values
func NewImportSpecificationFromFileConflict() *ImportSpecificationFromFileConflict {
	return &ImportSpecificationFromFileConflict{}
}

/*ImportSpecificationFromFileConflict handles this case with default header values.

HttpError
*/
type ImportSpecificationFromFileConflict struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromFileConflict) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileConflict  %+v", 409, o.Payload)
}

func (o *ImportSpecificationFromFileConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportSpecificationFromFileUnprocessableEntity creates a ImportSpecificationFromFileUnprocessableEntity with default headers values
func NewImportSpecificationFromFileUnprocessableEntity() *ImportSpecificationFromFileUnprocessableEntity {
	return &ImportSpecificationFromFileUnprocessableEntity{}
}

/*ImportSpecificationFromFileUnprocessableEntity handles this case with default header values.

HttpError
*/
type ImportSpecificationFromFileUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ImportSpecificationFromFileUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/admin/{tid}/services/{sid}/apis/import/file][%d] importSpecificationFromFileUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ImportSpecificationFromFileUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ImportSpecificationFromFileUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
