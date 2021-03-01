// Code generated by go-swagger; DO NOT EDIT.

package file_payments

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

// CreateFilePaymentConsentsConsentIDFileReader is a Reader for the CreateFilePaymentConsentsConsentIDFile structure.
type CreateFilePaymentConsentsConsentIDFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFilePaymentConsentsConsentIDFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateFilePaymentConsentsConsentIDFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFilePaymentConsentsConsentIDFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateFilePaymentConsentsConsentIDFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFilePaymentConsentsConsentIDFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateFilePaymentConsentsConsentIDFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateFilePaymentConsentsConsentIDFileMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewCreateFilePaymentConsentsConsentIDFileNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewCreateFilePaymentConsentsConsentIDFileUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateFilePaymentConsentsConsentIDFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateFilePaymentConsentsConsentIDFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFilePaymentConsentsConsentIDFileOK creates a CreateFilePaymentConsentsConsentIDFileOK with default headers values
func NewCreateFilePaymentConsentsConsentIDFileOK() *CreateFilePaymentConsentsConsentIDFileOK {
	return &CreateFilePaymentConsentsConsentIDFileOK{}
}

/* CreateFilePaymentConsentsConsentIDFileOK describes a response with status code 200, with default header values.

File Payment Consents Created
*/
type CreateFilePaymentConsentsConsentIDFileOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileOK) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileOK ", 200)
}

func (o *CreateFilePaymentConsentsConsentIDFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsConsentIDFileBadRequest creates a CreateFilePaymentConsentsConsentIDFileBadRequest with default headers values
func NewCreateFilePaymentConsentsConsentIDFileBadRequest() *CreateFilePaymentConsentsConsentIDFileBadRequest {
	return &CreateFilePaymentConsentsConsentIDFileBadRequest{}
}

/* CreateFilePaymentConsentsConsentIDFileBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateFilePaymentConsentsConsentIDFileBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentConsentsConsentIDFileBadRequest) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFilePaymentConsentsConsentIDFileBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsConsentIDFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsConsentIDFileUnauthorized creates a CreateFilePaymentConsentsConsentIDFileUnauthorized with default headers values
func NewCreateFilePaymentConsentsConsentIDFileUnauthorized() *CreateFilePaymentConsentsConsentIDFileUnauthorized {
	return &CreateFilePaymentConsentsConsentIDFileUnauthorized{}
}

/* CreateFilePaymentConsentsConsentIDFileUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateFilePaymentConsentsConsentIDFileUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileUnauthorized) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileUnauthorized ", 401)
}

func (o *CreateFilePaymentConsentsConsentIDFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsConsentIDFileForbidden creates a CreateFilePaymentConsentsConsentIDFileForbidden with default headers values
func NewCreateFilePaymentConsentsConsentIDFileForbidden() *CreateFilePaymentConsentsConsentIDFileForbidden {
	return &CreateFilePaymentConsentsConsentIDFileForbidden{}
}

/* CreateFilePaymentConsentsConsentIDFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateFilePaymentConsentsConsentIDFileForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentConsentsConsentIDFileForbidden) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileForbidden  %+v", 403, o.Payload)
}
func (o *CreateFilePaymentConsentsConsentIDFileForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsConsentIDFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsConsentIDFileNotFound creates a CreateFilePaymentConsentsConsentIDFileNotFound with default headers values
func NewCreateFilePaymentConsentsConsentIDFileNotFound() *CreateFilePaymentConsentsConsentIDFileNotFound {
	return &CreateFilePaymentConsentsConsentIDFileNotFound{}
}

/* CreateFilePaymentConsentsConsentIDFileNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateFilePaymentConsentsConsentIDFileNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileNotFound) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileNotFound ", 404)
}

func (o *CreateFilePaymentConsentsConsentIDFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsConsentIDFileMethodNotAllowed creates a CreateFilePaymentConsentsConsentIDFileMethodNotAllowed with default headers values
func NewCreateFilePaymentConsentsConsentIDFileMethodNotAllowed() *CreateFilePaymentConsentsConsentIDFileMethodNotAllowed {
	return &CreateFilePaymentConsentsConsentIDFileMethodNotAllowed{}
}

/* CreateFilePaymentConsentsConsentIDFileMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type CreateFilePaymentConsentsConsentIDFileMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileMethodNotAllowed ", 405)
}

func (o *CreateFilePaymentConsentsConsentIDFileMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsConsentIDFileNotAcceptable creates a CreateFilePaymentConsentsConsentIDFileNotAcceptable with default headers values
func NewCreateFilePaymentConsentsConsentIDFileNotAcceptable() *CreateFilePaymentConsentsConsentIDFileNotAcceptable {
	return &CreateFilePaymentConsentsConsentIDFileNotAcceptable{}
}

/* CreateFilePaymentConsentsConsentIDFileNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type CreateFilePaymentConsentsConsentIDFileNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileNotAcceptable ", 406)
}

func (o *CreateFilePaymentConsentsConsentIDFileNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsConsentIDFileUnsupportedMediaType creates a CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType with default headers values
func NewCreateFilePaymentConsentsConsentIDFileUnsupportedMediaType() *CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType {
	return &CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType{}
}

/* CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type
*/
type CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileUnsupportedMediaType ", 415)
}

func (o *CreateFilePaymentConsentsConsentIDFileUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewCreateFilePaymentConsentsConsentIDFileTooManyRequests creates a CreateFilePaymentConsentsConsentIDFileTooManyRequests with default headers values
func NewCreateFilePaymentConsentsConsentIDFileTooManyRequests() *CreateFilePaymentConsentsConsentIDFileTooManyRequests {
	return &CreateFilePaymentConsentsConsentIDFileTooManyRequests{}
}

/* CreateFilePaymentConsentsConsentIDFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateFilePaymentConsentsConsentIDFileTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *CreateFilePaymentConsentsConsentIDFileTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileTooManyRequests ", 429)
}

func (o *CreateFilePaymentConsentsConsentIDFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateFilePaymentConsentsConsentIDFileInternalServerError creates a CreateFilePaymentConsentsConsentIDFileInternalServerError with default headers values
func NewCreateFilePaymentConsentsConsentIDFileInternalServerError() *CreateFilePaymentConsentsConsentIDFileInternalServerError {
	return &CreateFilePaymentConsentsConsentIDFileInternalServerError{}
}

/* CreateFilePaymentConsentsConsentIDFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateFilePaymentConsentsConsentIDFileInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	/* Header containing a detached JWS signature of the body of the payload.

	 */
	XJwsSignature string

	Payload *models.OBErrorResponse1
}

func (o *CreateFilePaymentConsentsConsentIDFileInternalServerError) Error() string {
	return fmt.Sprintf("[POST /file-payment-consents/{ConsentId}/file][%d] createFilePaymentConsentsConsentIdFileInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateFilePaymentConsentsConsentIDFileInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *CreateFilePaymentConsentsConsentIDFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
