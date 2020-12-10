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

	"github.com/openbanking-sandbox/models"
)

// GetAccountsAccountIDStatementsReader is a Reader for the GetAccountsAccountIDStatements structure.
type GetAccountsAccountIDStatementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDStatementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDStatementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDStatementsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDStatementsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDStatementsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDStatementsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDStatementsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDStatementsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDStatementsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDStatementsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDStatementsOK creates a GetAccountsAccountIDStatementsOK with default headers values
func NewGetAccountsAccountIDStatementsOK() *GetAccountsAccountIDStatementsOK {
	return &GetAccountsAccountIDStatementsOK{}
}

/*GetAccountsAccountIDStatementsOK handles this case with default header values.

Statements Read
*/
type GetAccountsAccountIDStatementsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadStatement2
}

func (o *GetAccountsAccountIDStatementsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDStatementsOK) GetPayload() *models.OBReadStatement2 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadStatement2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsBadRequest creates a GetAccountsAccountIDStatementsBadRequest with default headers values
func NewGetAccountsAccountIDStatementsBadRequest() *GetAccountsAccountIDStatementsBadRequest {
	return &GetAccountsAccountIDStatementsBadRequest{}
}

/*GetAccountsAccountIDStatementsBadRequest handles this case with default header values.

Bad request
*/
type GetAccountsAccountIDStatementsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountsAccountIDStatementsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsUnauthorized creates a GetAccountsAccountIDStatementsUnauthorized with default headers values
func NewGetAccountsAccountIDStatementsUnauthorized() *GetAccountsAccountIDStatementsUnauthorized {
	return &GetAccountsAccountIDStatementsUnauthorized{}
}

/*GetAccountsAccountIDStatementsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountsAccountIDStatementsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsUnauthorized ", 401)
}

func (o *GetAccountsAccountIDStatementsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsForbidden creates a GetAccountsAccountIDStatementsForbidden with default headers values
func NewGetAccountsAccountIDStatementsForbidden() *GetAccountsAccountIDStatementsForbidden {
	return &GetAccountsAccountIDStatementsForbidden{}
}

/*GetAccountsAccountIDStatementsForbidden handles this case with default header values.

Forbidden
*/
type GetAccountsAccountIDStatementsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountsAccountIDStatementsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDStatementsNotFound creates a GetAccountsAccountIDStatementsNotFound with default headers values
func NewGetAccountsAccountIDStatementsNotFound() *GetAccountsAccountIDStatementsNotFound {
	return &GetAccountsAccountIDStatementsNotFound{}
}

/*GetAccountsAccountIDStatementsNotFound handles this case with default header values.

Not found
*/
type GetAccountsAccountIDStatementsNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsNotFound ", 404)
}

func (o *GetAccountsAccountIDStatementsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsMethodNotAllowed creates a GetAccountsAccountIDStatementsMethodNotAllowed with default headers values
func NewGetAccountsAccountIDStatementsMethodNotAllowed() *GetAccountsAccountIDStatementsMethodNotAllowed {
	return &GetAccountsAccountIDStatementsMethodNotAllowed{}
}

/*GetAccountsAccountIDStatementsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDStatementsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDStatementsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsNotAcceptable creates a GetAccountsAccountIDStatementsNotAcceptable with default headers values
func NewGetAccountsAccountIDStatementsNotAcceptable() *GetAccountsAccountIDStatementsNotAcceptable {
	return &GetAccountsAccountIDStatementsNotAcceptable{}
}

/*GetAccountsAccountIDStatementsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetAccountsAccountIDStatementsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDStatementsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDStatementsTooManyRequests creates a GetAccountsAccountIDStatementsTooManyRequests with default headers values
func NewGetAccountsAccountIDStatementsTooManyRequests() *GetAccountsAccountIDStatementsTooManyRequests {
	return &GetAccountsAccountIDStatementsTooManyRequests{}
}

/*GetAccountsAccountIDStatementsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetAccountsAccountIDStatementsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDStatementsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDStatementsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDStatementsInternalServerError creates a GetAccountsAccountIDStatementsInternalServerError with default headers values
func NewGetAccountsAccountIDStatementsInternalServerError() *GetAccountsAccountIDStatementsInternalServerError {
	return &GetAccountsAccountIDStatementsInternalServerError{}
}

/*GetAccountsAccountIDStatementsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountsAccountIDStatementsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDStatementsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/statements][%d] getAccountsAccountIdStatementsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountsAccountIDStatementsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDStatementsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
