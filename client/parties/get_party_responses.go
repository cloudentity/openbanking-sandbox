// Code generated by go-swagger; DO NOT EDIT.

package parties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-sandbox/models"
)

// GetPartyReader is a Reader for the GetParty structure.
type GetPartyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPartyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPartyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPartyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPartyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPartyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPartyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetPartyMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetPartyNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetPartyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPartyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPartyOK creates a GetPartyOK with default headers values
func NewGetPartyOK() *GetPartyOK {
	return &GetPartyOK{}
}

/*GetPartyOK handles this case with default header values.

Parties Read
*/
type GetPartyOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadParty2
}

func (o *GetPartyOK) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyOK  %+v", 200, o.Payload)
}

func (o *GetPartyOK) GetPayload() *models.OBReadParty2 {
	return o.Payload
}

func (o *GetPartyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadParty2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartyBadRequest creates a GetPartyBadRequest with default headers values
func NewGetPartyBadRequest() *GetPartyBadRequest {
	return &GetPartyBadRequest{}
}

/*GetPartyBadRequest handles this case with default header values.

Bad request
*/
type GetPartyBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetPartyBadRequest) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyBadRequest  %+v", 400, o.Payload)
}

func (o *GetPartyBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetPartyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartyUnauthorized creates a GetPartyUnauthorized with default headers values
func NewGetPartyUnauthorized() *GetPartyUnauthorized {
	return &GetPartyUnauthorized{}
}

/*GetPartyUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPartyUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetPartyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyUnauthorized ", 401)
}

func (o *GetPartyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetPartyForbidden creates a GetPartyForbidden with default headers values
func NewGetPartyForbidden() *GetPartyForbidden {
	return &GetPartyForbidden{}
}

/*GetPartyForbidden handles this case with default header values.

Forbidden
*/
type GetPartyForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetPartyForbidden) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyForbidden  %+v", 403, o.Payload)
}

func (o *GetPartyForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetPartyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPartyNotFound creates a GetPartyNotFound with default headers values
func NewGetPartyNotFound() *GetPartyNotFound {
	return &GetPartyNotFound{}
}

/*GetPartyNotFound handles this case with default header values.

Not found
*/
type GetPartyNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetPartyNotFound) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyNotFound ", 404)
}

func (o *GetPartyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetPartyMethodNotAllowed creates a GetPartyMethodNotAllowed with default headers values
func NewGetPartyMethodNotAllowed() *GetPartyMethodNotAllowed {
	return &GetPartyMethodNotAllowed{}
}

/*GetPartyMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetPartyMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetPartyMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyMethodNotAllowed ", 405)
}

func (o *GetPartyMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetPartyNotAcceptable creates a GetPartyNotAcceptable with default headers values
func NewGetPartyNotAcceptable() *GetPartyNotAcceptable {
	return &GetPartyNotAcceptable{}
}

/*GetPartyNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetPartyNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetPartyNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyNotAcceptable ", 406)
}

func (o *GetPartyNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetPartyTooManyRequests creates a GetPartyTooManyRequests with default headers values
func NewGetPartyTooManyRequests() *GetPartyTooManyRequests {
	return &GetPartyTooManyRequests{}
}

/*GetPartyTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetPartyTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetPartyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyTooManyRequests ", 429)
}

func (o *GetPartyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Retry-After
	retryAfter, err := swag.ConvertInt64(response.GetHeader("Retry-After"))
	if err != nil {
		return errors.InvalidType("Retry-After", "header", "int64", response.GetHeader("Retry-After"))
	}
	o.RetryAfter = retryAfter

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetPartyInternalServerError creates a GetPartyInternalServerError with default headers values
func NewGetPartyInternalServerError() *GetPartyInternalServerError {
	return &GetPartyInternalServerError{}
}

/*GetPartyInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPartyInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetPartyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /party][%d] getPartyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPartyInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetPartyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
