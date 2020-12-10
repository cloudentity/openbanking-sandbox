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

// GetOktaIDPReader is a Reader for the GetOktaIDP structure.
type GetOktaIDPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOktaIDPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOktaIDPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOktaIDPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOktaIDPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOktaIDPNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOktaIDPOK creates a GetOktaIDPOK with default headers values
func NewGetOktaIDPOK() *GetOktaIDPOK {
	return &GetOktaIDPOK{}
}

/*GetOktaIDPOK handles this case with default header values.

OktaIDP
*/
type GetOktaIDPOK struct {
	Payload *models.OktaIDP
}

func (o *GetOktaIDPOK) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] getOktaIdPOK  %+v", 200, o.Payload)
}

func (o *GetOktaIDPOK) GetPayload() *models.OktaIDP {
	return o.Payload
}

func (o *GetOktaIDPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OktaIDP)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOktaIDPUnauthorized creates a GetOktaIDPUnauthorized with default headers values
func NewGetOktaIDPUnauthorized() *GetOktaIDPUnauthorized {
	return &GetOktaIDPUnauthorized{}
}

/*GetOktaIDPUnauthorized handles this case with default header values.

HttpError
*/
type GetOktaIDPUnauthorized struct {
	Payload *models.Error
}

func (o *GetOktaIDPUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] getOktaIdPUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOktaIDPUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOktaIDPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOktaIDPForbidden creates a GetOktaIDPForbidden with default headers values
func NewGetOktaIDPForbidden() *GetOktaIDPForbidden {
	return &GetOktaIDPForbidden{}
}

/*GetOktaIDPForbidden handles this case with default header values.

HttpError
*/
type GetOktaIDPForbidden struct {
	Payload *models.Error
}

func (o *GetOktaIDPForbidden) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] getOktaIdPForbidden  %+v", 403, o.Payload)
}

func (o *GetOktaIDPForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOktaIDPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOktaIDPNotFound creates a GetOktaIDPNotFound with default headers values
func NewGetOktaIDPNotFound() *GetOktaIDPNotFound {
	return &GetOktaIDPNotFound{}
}

/*GetOktaIDPNotFound handles this case with default header values.

HttpError
*/
type GetOktaIDPNotFound struct {
	Payload *models.Error
}

func (o *GetOktaIDPNotFound) Error() string {
	return fmt.Sprintf("[GET /api/admin/{tid}/servers/{aid}/idps/okta/{iid}][%d] getOktaIdPNotFound  %+v", 404, o.Payload)
}

func (o *GetOktaIDPNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOktaIDPNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
