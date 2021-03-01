// Code generated by go-swagger; DO NOT EDIT.

package international_standing_orders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-sandbox/openbanking/paymentinitiation/models"
)

// CreateInternationalStandingOrdersReader is a Reader for the CreateInternationalStandingOrders structure.
type CreateInternationalStandingOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInternationalStandingOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInternationalStandingOrdersCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInternationalStandingOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInternationalStandingOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInternationalStandingOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInternationalStandingOrdersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateInternationalStandingOrdersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateInternationalStandingOrdersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateInternationalStandingOrdersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateInternationalStandingOrdersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInternationalStandingOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInternationalStandingOrdersCreated creates a CreateInternationalStandingOrdersCreated with default headers values
func NewCreateInternationalStandingOrdersCreated() *CreateInternationalStandingOrdersCreated {
	return &CreateInternationalStandingOrdersCreated{}
}

/* CreateInternationalStandingOrdersCreated describes a response with status code 201, with default header values.

International Standing Orders Created
*/
type CreateInternationalStandingOrdersCreated struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteInternationalStandingOrderResponse7
}

func (o *CreateInternationalStandingOrdersCreated) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersCreated  %+v", 201, o.Payload)
}
func (o *CreateInternationalStandingOrdersCreated) GetPayload() *models.OBWriteInternationalStandingOrderResponse7 {
	return o.Payload
}

func (o *CreateInternationalStandingOrdersCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBWriteInternationalStandingOrderResponse7)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrdersBadRequest creates a CreateInternationalStandingOrdersBadRequest with default headers values
func NewCreateInternationalStandingOrdersBadRequest() *CreateInternationalStandingOrdersBadRequest {
	return &CreateInternationalStandingOrdersBadRequest{}
}

/* CreateInternationalStandingOrdersBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateInternationalStandingOrdersBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalStandingOrdersBadRequest) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersBadRequest  %+v", 400, o.Payload)
}
func (o *CreateInternationalStandingOrdersBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalStandingOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrdersUnauthorized creates a CreateInternationalStandingOrdersUnauthorized with default headers values
func NewCreateInternationalStandingOrdersUnauthorized() *CreateInternationalStandingOrdersUnauthorized {
	return &CreateInternationalStandingOrdersUnauthorized{}
}

/* CreateInternationalStandingOrdersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateInternationalStandingOrdersUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalStandingOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersUnauthorized ", 401)
}

func (o *CreateInternationalStandingOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalStandingOrdersForbidden creates a CreateInternationalStandingOrdersForbidden with default headers values
func NewCreateInternationalStandingOrdersForbidden() *CreateInternationalStandingOrdersForbidden {
	return &CreateInternationalStandingOrdersForbidden{}
}

/* CreateInternationalStandingOrdersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateInternationalStandingOrdersForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalStandingOrdersForbidden) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersForbidden  %+v", 403, o.Payload)
}
func (o *CreateInternationalStandingOrdersForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalStandingOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInternationalStandingOrdersNotFound creates a CreateInternationalStandingOrdersNotFound with default headers values
func NewCreateInternationalStandingOrdersNotFound() *CreateInternationalStandingOrdersNotFound {
	return &CreateInternationalStandingOrdersNotFound{}
}

/* CreateInternationalStandingOrdersNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateInternationalStandingOrdersNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalStandingOrdersNotFound) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersNotFound ", 404)
}

func (o *CreateInternationalStandingOrdersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalStandingOrdersMethodNotAllowed creates a CreateInternationalStandingOrdersMethodNotAllowed with default headers values
func NewCreateInternationalStandingOrdersMethodNotAllowed() *CreateInternationalStandingOrdersMethodNotAllowed {
	return &CreateInternationalStandingOrdersMethodNotAllowed{}
}

/* CreateInternationalStandingOrdersMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateInternationalStandingOrdersMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalStandingOrdersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersMethodNotAllowed ", 405)
}

func (o *CreateInternationalStandingOrdersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalStandingOrdersNotAcceptable creates a CreateInternationalStandingOrdersNotAcceptable with default headers values
func NewCreateInternationalStandingOrdersNotAcceptable() *CreateInternationalStandingOrdersNotAcceptable {
	return &CreateInternationalStandingOrdersNotAcceptable{}
}

/* CreateInternationalStandingOrdersNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateInternationalStandingOrdersNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalStandingOrdersNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersNotAcceptable ", 406)
}

func (o *CreateInternationalStandingOrdersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalStandingOrdersUnsupportedMediaType creates a CreateInternationalStandingOrdersUnsupportedMediaType with default headers values
func NewCreateInternationalStandingOrdersUnsupportedMediaType() *CreateInternationalStandingOrdersUnsupportedMediaType {
	return &CreateInternationalStandingOrdersUnsupportedMediaType{}
}

/* CreateInternationalStandingOrdersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateInternationalStandingOrdersUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalStandingOrdersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersUnsupportedMediaType ", 415)
}

func (o *CreateInternationalStandingOrdersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalStandingOrdersTooManyRequests creates a CreateInternationalStandingOrdersTooManyRequests with default headers values
func NewCreateInternationalStandingOrdersTooManyRequests() *CreateInternationalStandingOrdersTooManyRequests {
	return &CreateInternationalStandingOrdersTooManyRequests{}
}

/* CreateInternationalStandingOrdersTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateInternationalStandingOrdersTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateInternationalStandingOrdersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersTooManyRequests ", 429)
}

func (o *CreateInternationalStandingOrdersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Retry-After
	hdrRetryAfter := response.GetHeader("Retry-After")

	if hdrRetryAfter != "" {
		valretryAfter, err := swag.ConvertInt64(hdrRetryAfter)
		if err != nil {
			return errors.InvalidType("Retry-After", "header", "int64", hdrRetryAfter)
		}
		o.RetryAfter = valretryAfter
	}

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateInternationalStandingOrdersInternalServerError creates a CreateInternationalStandingOrdersInternalServerError with default headers values
func NewCreateInternationalStandingOrdersInternalServerError() *CreateInternationalStandingOrdersInternalServerError {
	return &CreateInternationalStandingOrdersInternalServerError{}
}

/* CreateInternationalStandingOrdersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateInternationalStandingOrdersInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateInternationalStandingOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /international-standing-orders][%d] createInternationalStandingOrdersInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateInternationalStandingOrdersInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateInternationalStandingOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	// hydrates response header x-jws-signature
	hdrXJwsSignature := response.GetHeader("x-jws-signature")

	if hdrXJwsSignature != "" {
		o.XJwsSignature = hdrXJwsSignature
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
