// Code generated by go-swagger; DO NOT EDIT.

package statements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/models"
)

// GetAccountsAccountIDStatementsStatementIDReader is a Reader for the GetAccountsAccountIDStatementsStatementID structure.
type GetAccountsAccountIDStatementsStatementIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDStatementsStatementIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDStatementsStatementIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDStatementsStatementIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDStatementsStatementIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDStatementsStatementIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDStatementsStatementIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDStatementsStatementIDMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDStatementsStatementIDNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDStatementsStatementIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDStatementsStatementIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDStatementsStatementIDOK creates a GetAccountsAccountIDStatementsStatementIDOK with default headers values
func NewGetAccountsAccountIDStatementsStatementIDOK() *GetAccountsAccountIDStatementsStatementIDOK {
	return &GetAccountsAccountIDStatementsStatementIDOK{}
}

/* GetAccountsAccountIDStatementsStatementIDOK describes a response with status code 200, with default header values.

Statements Read
*/
type GetAccountsAccountIDStatementsStatementIDOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadStatement2
}

func (o *GetAccountsAccountIDStatementsStatementIDOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdOK  %+v", 200, o.Payload)
}
func (o *GetAccountsAccountIDStatementsStatementIDOK) GetPayload() *models.OBReadStatement2 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadStatement2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDBadRequest creates a GetAccountsAccountIDStatementsStatementIDBadRequest with default headers values
func NewGetAccountsAccountIDStatementsStatementIDBadRequest() *GetAccountsAccountIDStatementsStatementIDBadRequest {
	return &GetAccountsAccountIDStatementsStatementIDBadRequest{}
}

/* GetAccountsAccountIDStatementsStatementIDBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAccountsAccountIDStatementsStatementIDBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsStatementIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountsAccountIDStatementsStatementIDBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDStatementsStatementIDUnauthorized creates a GetAccountsAccountIDStatementsStatementIDUnauthorized with default headers values
func NewGetAccountsAccountIDStatementsStatementIDUnauthorized() *GetAccountsAccountIDStatementsStatementIDUnauthorized {
	return &GetAccountsAccountIDStatementsStatementIDUnauthorized{}
}

/* GetAccountsAccountIDStatementsStatementIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccountsAccountIDStatementsStatementIDUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdUnauthorized ", 401)
}

func (o *GetAccountsAccountIDStatementsStatementIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDForbidden creates a GetAccountsAccountIDStatementsStatementIDForbidden with default headers values
func NewGetAccountsAccountIDStatementsStatementIDForbidden() *GetAccountsAccountIDStatementsStatementIDForbidden {
	return &GetAccountsAccountIDStatementsStatementIDForbidden{}
}

/* GetAccountsAccountIDStatementsStatementIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountsAccountIDStatementsStatementIDForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsStatementIDForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountsAccountIDStatementsStatementIDForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDStatementsStatementIDNotFound creates a GetAccountsAccountIDStatementsStatementIDNotFound with default headers values
func NewGetAccountsAccountIDStatementsStatementIDNotFound() *GetAccountsAccountIDStatementsStatementIDNotFound {
	return &GetAccountsAccountIDStatementsStatementIDNotFound{}
}

/* GetAccountsAccountIDStatementsStatementIDNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAccountsAccountIDStatementsStatementIDNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdNotFound ", 404)
}

func (o *GetAccountsAccountIDStatementsStatementIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDMethodNotAllowed creates a GetAccountsAccountIDStatementsStatementIDMethodNotAllowed with default headers values
func NewGetAccountsAccountIDStatementsStatementIDMethodNotAllowed() *GetAccountsAccountIDStatementsStatementIDMethodNotAllowed {
	return &GetAccountsAccountIDStatementsStatementIDMethodNotAllowed{}
}

/* GetAccountsAccountIDStatementsStatementIDMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDStatementsStatementIDMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDStatementsStatementIDMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDNotAcceptable creates a GetAccountsAccountIDStatementsStatementIDNotAcceptable with default headers values
func NewGetAccountsAccountIDStatementsStatementIDNotAcceptable() *GetAccountsAccountIDStatementsStatementIDNotAcceptable {
	return &GetAccountsAccountIDStatementsStatementIDNotAcceptable{}
}

/* GetAccountsAccountIDStatementsStatementIDNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetAccountsAccountIDStatementsStatementIDNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDStatementsStatementIDNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDStatementsStatementIDTooManyRequests creates a GetAccountsAccountIDStatementsStatementIDTooManyRequests with default headers values
func NewGetAccountsAccountIDStatementsStatementIDTooManyRequests() *GetAccountsAccountIDStatementsStatementIDTooManyRequests {
	return &GetAccountsAccountIDStatementsStatementIDTooManyRequests{}
}

/* GetAccountsAccountIDStatementsStatementIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccountsAccountIDStatementsStatementIDTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsStatementIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDStatementsStatementIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDStatementsStatementIDInternalServerError creates a GetAccountsAccountIDStatementsStatementIDInternalServerError with default headers values
func NewGetAccountsAccountIDStatementsStatementIDInternalServerError() *GetAccountsAccountIDStatementsStatementIDInternalServerError {
	return &GetAccountsAccountIDStatementsStatementIDInternalServerError{}
}

/* GetAccountsAccountIDStatementsStatementIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountsAccountIDStatementsStatementIDInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsStatementIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements/{StatementId}][%d] getAccountsAccountIdStatementsStatementIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountsAccountIDStatementsStatementIDInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsStatementIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
