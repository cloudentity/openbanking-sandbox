// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/cloudentity/acp/pkg/openbanking/models"
)

// GetAccountsAccountIDTransactionsReader is a Reader for the GetAccountsAccountIDTransactions structure.
type GetAccountsAccountIDTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDTransactionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDTransactionsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDTransactionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountsAccountIDTransactionsOK creates a GetAccountsAccountIDTransactionsOK with default headers values
func NewGetAccountsAccountIDTransactionsOK() *GetAccountsAccountIDTransactionsOK {
	return &GetAccountsAccountIDTransactionsOK{}
}

/*GetAccountsAccountIDTransactionsOK handles this case with default header values.

Transactions Read
*/
type GetAccountsAccountIDTransactionsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadTransaction6
}

func (o *GetAccountsAccountIDTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDTransactionsOK) GetPayload() *models.OBReadTransaction6 {
	return o.Payload
}

func (o *GetAccountsAccountIDTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadTransaction6)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDTransactionsBadRequest creates a GetAccountsAccountIDTransactionsBadRequest with default headers values
func NewGetAccountsAccountIDTransactionsBadRequest() *GetAccountsAccountIDTransactionsBadRequest {
	return &GetAccountsAccountIDTransactionsBadRequest{}
}

/*GetAccountsAccountIDTransactionsBadRequest handles this case with default header values.

Bad request
*/
type GetAccountsAccountIDTransactionsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountsAccountIDTransactionsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDTransactionsUnauthorized creates a GetAccountsAccountIDTransactionsUnauthorized with default headers values
func NewGetAccountsAccountIDTransactionsUnauthorized() *GetAccountsAccountIDTransactionsUnauthorized {
	return &GetAccountsAccountIDTransactionsUnauthorized{}
}

/*GetAccountsAccountIDTransactionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountsAccountIDTransactionsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsUnauthorized ", 401)
}

func (o *GetAccountsAccountIDTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDTransactionsForbidden creates a GetAccountsAccountIDTransactionsForbidden with default headers values
func NewGetAccountsAccountIDTransactionsForbidden() *GetAccountsAccountIDTransactionsForbidden {
	return &GetAccountsAccountIDTransactionsForbidden{}
}

/*GetAccountsAccountIDTransactionsForbidden handles this case with default header values.

Forbidden
*/
type GetAccountsAccountIDTransactionsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountsAccountIDTransactionsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDTransactionsMethodNotAllowed creates a GetAccountsAccountIDTransactionsMethodNotAllowed with default headers values
func NewGetAccountsAccountIDTransactionsMethodNotAllowed() *GetAccountsAccountIDTransactionsMethodNotAllowed {
	return &GetAccountsAccountIDTransactionsMethodNotAllowed{}
}

/*GetAccountsAccountIDTransactionsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDTransactionsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDTransactionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDTransactionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDTransactionsNotAcceptable creates a GetAccountsAccountIDTransactionsNotAcceptable with default headers values
func NewGetAccountsAccountIDTransactionsNotAcceptable() *GetAccountsAccountIDTransactionsNotAcceptable {
	return &GetAccountsAccountIDTransactionsNotAcceptable{}
}

/*GetAccountsAccountIDTransactionsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetAccountsAccountIDTransactionsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDTransactionsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDTransactionsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDTransactionsTooManyRequests creates a GetAccountsAccountIDTransactionsTooManyRequests with default headers values
func NewGetAccountsAccountIDTransactionsTooManyRequests() *GetAccountsAccountIDTransactionsTooManyRequests {
	return &GetAccountsAccountIDTransactionsTooManyRequests{}
}

/*GetAccountsAccountIDTransactionsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetAccountsAccountIDTransactionsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDTransactionsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDTransactionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDTransactionsInternalServerError creates a GetAccountsAccountIDTransactionsInternalServerError with default headers values
func NewGetAccountsAccountIDTransactionsInternalServerError() *GetAccountsAccountIDTransactionsInternalServerError {
	return &GetAccountsAccountIDTransactionsInternalServerError{}
}

/*GetAccountsAccountIDTransactionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountsAccountIDTransactionsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/transactions][%d] getAccountsAccountIdTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountsAccountIDTransactionsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}