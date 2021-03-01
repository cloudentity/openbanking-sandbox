// Code generated by go-swagger; DO NOT EDIT.

package account_access

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

// DeleteAccountAccessConsentsConsentIDReader is a Reader for the DeleteAccountAccessConsentsConsentID structure.
type DeleteAccountAccessConsentsConsentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountAccessConsentsConsentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAccountAccessConsentsConsentIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAccountAccessConsentsConsentIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAccountAccessConsentsConsentIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAccountAccessConsentsConsentIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteAccountAccessConsentsConsentIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewDeleteAccountAccessConsentsConsentIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteAccountAccessConsentsConsentIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAccountAccessConsentsConsentIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAccountAccessConsentsConsentIDNoContent creates a DeleteAccountAccessConsentsConsentIDNoContent with default headers values
func NewDeleteAccountAccessConsentsConsentIDNoContent() *DeleteAccountAccessConsentsConsentIDNoContent {
	return &DeleteAccountAccessConsentsConsentIDNoContent{}
}

/* DeleteAccountAccessConsentsConsentIDNoContent describes a response with status code 204, with default header values.

Account Access Consents Deleted
*/
type DeleteAccountAccessConsentsConsentIDNoContent struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *DeleteAccountAccessConsentsConsentIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdNoContent ", 204)
}

func (o *DeleteAccountAccessConsentsConsentIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewDeleteAccountAccessConsentsConsentIDBadRequest creates a DeleteAccountAccessConsentsConsentIDBadRequest with default headers values
func NewDeleteAccountAccessConsentsConsentIDBadRequest() *DeleteAccountAccessConsentsConsentIDBadRequest {
	return &DeleteAccountAccessConsentsConsentIDBadRequest{}
}

/* DeleteAccountAccessConsentsConsentIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteAccountAccessConsentsConsentIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *DeleteAccountAccessConsentsConsentIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteAccountAccessConsentsConsentIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *DeleteAccountAccessConsentsConsentIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountAccessConsentsConsentIDUnauthorized creates a DeleteAccountAccessConsentsConsentIDUnauthorized with default headers values
func NewDeleteAccountAccessConsentsConsentIDUnauthorized() *DeleteAccountAccessConsentsConsentIDUnauthorized {
	return &DeleteAccountAccessConsentsConsentIDUnauthorized{}
}

/* DeleteAccountAccessConsentsConsentIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteAccountAccessConsentsConsentIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *DeleteAccountAccessConsentsConsentIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdUnauthorized ", 401)
}

func (o *DeleteAccountAccessConsentsConsentIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewDeleteAccountAccessConsentsConsentIDForbidden creates a DeleteAccountAccessConsentsConsentIDForbidden with default headers values
func NewDeleteAccountAccessConsentsConsentIDForbidden() *DeleteAccountAccessConsentsConsentIDForbidden {
	return &DeleteAccountAccessConsentsConsentIDForbidden{}
}

/* DeleteAccountAccessConsentsConsentIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAccountAccessConsentsConsentIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *DeleteAccountAccessConsentsConsentIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdForbidden  %+v", 403, o.Payload)
}
func (o *DeleteAccountAccessConsentsConsentIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *DeleteAccountAccessConsentsConsentIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAccountAccessConsentsConsentIDMethodNotAllowed creates a DeleteAccountAccessConsentsConsentIDMethodNotAllowed with default headers values
func NewDeleteAccountAccessConsentsConsentIDMethodNotAllowed() *DeleteAccountAccessConsentsConsentIDMethodNotAllowed {
	return &DeleteAccountAccessConsentsConsentIDMethodNotAllowed{}
}

/* DeleteAccountAccessConsentsConsentIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type DeleteAccountAccessConsentsConsentIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *DeleteAccountAccessConsentsConsentIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdMethodNotAllowed ", 405)
}

func (o *DeleteAccountAccessConsentsConsentIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewDeleteAccountAccessConsentsConsentIDNotAcceptable creates a DeleteAccountAccessConsentsConsentIDNotAcceptable with default headers values
func NewDeleteAccountAccessConsentsConsentIDNotAcceptable() *DeleteAccountAccessConsentsConsentIDNotAcceptable {
	return &DeleteAccountAccessConsentsConsentIDNotAcceptable{}
}

/* DeleteAccountAccessConsentsConsentIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type DeleteAccountAccessConsentsConsentIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *DeleteAccountAccessConsentsConsentIDNotAcceptable) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdNotAcceptable ", 406)
}

func (o *DeleteAccountAccessConsentsConsentIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewDeleteAccountAccessConsentsConsentIDTooManyRequests creates a DeleteAccountAccessConsentsConsentIDTooManyRequests with default headers values
func NewDeleteAccountAccessConsentsConsentIDTooManyRequests() *DeleteAccountAccessConsentsConsentIDTooManyRequests {
	return &DeleteAccountAccessConsentsConsentIDTooManyRequests{}
}

/* DeleteAccountAccessConsentsConsentIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteAccountAccessConsentsConsentIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *DeleteAccountAccessConsentsConsentIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdTooManyRequests ", 429)
}

func (o *DeleteAccountAccessConsentsConsentIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteAccountAccessConsentsConsentIDInternalServerError creates a DeleteAccountAccessConsentsConsentIDInternalServerError with default headers values
func NewDeleteAccountAccessConsentsConsentIDInternalServerError() *DeleteAccountAccessConsentsConsentIDInternalServerError {
	return &DeleteAccountAccessConsentsConsentIDInternalServerError{}
}

/* DeleteAccountAccessConsentsConsentIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteAccountAccessConsentsConsentIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *DeleteAccountAccessConsentsConsentIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /account-access-consents/{ConsentId}][%d] deleteAccountAccessConsentsConsentIdInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteAccountAccessConsentsConsentIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *DeleteAccountAccessConsentsConsentIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
