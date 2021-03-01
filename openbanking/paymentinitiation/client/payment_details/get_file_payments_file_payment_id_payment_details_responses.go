// Code generated by go-swagger; DO NOT EDIT.

package payment_details

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

// GetFilePaymentsFilePaymentIDPaymentDetailsReader is a Reader for the GetFilePaymentsFilePaymentIDPaymentDetails structure.
type GetFilePaymentsFilePaymentIDPaymentDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsOK creates a GetFilePaymentsFilePaymentIDPaymentDetailsOK with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsOK() *GetFilePaymentsFilePaymentIDPaymentDetailsOK {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsOK{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsOK describes a response with status code 200, with default header values.

Payment Details Read
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBWritePaymentDetailsResponse1
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsOK) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsOK  %+v", 200, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsOK) GetPayload() *models.OBWritePaymentDetailsResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.OBWritePaymentDetailsResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsBadRequest creates a GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsBadRequest() *GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsBadRequest  %+v", 400, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized creates a GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized() *GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsUnauthorized ", 401)
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsForbidden creates a GetFilePaymentsFilePaymentIDPaymentDetailsForbidden with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsForbidden() *GetFilePaymentsFilePaymentIDPaymentDetailsForbidden {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsForbidden{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsForbidden) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsForbidden  %+v", 403, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentsFilePaymentIDPaymentDetailsNotFound creates a GetFilePaymentsFilePaymentIDPaymentDetailsNotFound with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsNotFound() *GetFilePaymentsFilePaymentIDPaymentDetailsNotFound {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsNotFound{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsNotFound ", 404)
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed creates a GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed() *GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsMethodNotAllowed ", 405)
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable creates a GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable() *GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsNotAcceptable ", 406)
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests creates a GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests() *GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsTooManyRequests ", 429)
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError creates a GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError with default headers values
func NewGetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError() *GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError {
	return &GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError{}
}

/* GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /file-payments/{FilePaymentId}/payment-details][%d] getFilePaymentsFilePaymentIdPaymentDetailsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetFilePaymentsFilePaymentIDPaymentDetailsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
