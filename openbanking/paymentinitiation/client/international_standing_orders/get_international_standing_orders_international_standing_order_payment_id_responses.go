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

// GetInternationalStandingOrdersInternationalStandingOrderPaymentIDReader is a Reader for the GetInternationalStandingOrdersInternationalStandingOrderPaymentID structure.
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK describes a response with status code 200, with default header values.

International Standing Orders Read
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWriteInternationalStandingOrderResponse7
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdOK  %+v", 200, o.Payload)
}
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK) GetPayload() *models.OBWriteInternationalStandingOrderResponse7 {
	return o.Payload
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdUnauthorized ", 401)
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdForbidden  %+v", 403, o.Payload)
}
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdNotFound ", 404)
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdMethodNotAllowed ", 405)
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdNotAcceptable ", 406)
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdTooManyRequests ", 429)
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError creates a GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError with default headers values
func NewGetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError() *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError {
	return &GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError{}
}

/* GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /international-standing-orders/{InternationalStandingOrderPaymentId}][%d] getInternationalStandingOrdersInternationalStandingOrderPaymentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetInternationalStandingOrdersInternationalStandingOrderPaymentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
