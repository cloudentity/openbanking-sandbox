// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

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

// GetAccountsAccountIDDirectDebitsReader is a Reader for the GetAccountsAccountIDDirectDebits structure.
type GetAccountsAccountIDDirectDebitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDDirectDebitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDDirectDebitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDDirectDebitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDDirectDebitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDDirectDebitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDDirectDebitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDDirectDebitsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDDirectDebitsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDDirectDebitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDDirectDebitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDDirectDebitsOK creates a GetAccountsAccountIDDirectDebitsOK with default headers values
func NewGetAccountsAccountIDDirectDebitsOK() *GetAccountsAccountIDDirectDebitsOK {
	return &GetAccountsAccountIDDirectDebitsOK{}
}

/*GetAccountsAccountIDDirectDebitsOK handles this case with default header values.

Direct Debits Read
*/
type GetAccountsAccountIDDirectDebitsOK struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadDirectDebit2
}

func (o *GetAccountsAccountIDDirectDebitsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsAccountIDDirectDebitsOK) GetPayload() *models.OBReadDirectDebit2 {
	return o.Payload
}

func (o *GetAccountsAccountIDDirectDebitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBReadDirectDebit2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDDirectDebitsBadRequest creates a GetAccountsAccountIDDirectDebitsBadRequest with default headers values
func NewGetAccountsAccountIDDirectDebitsBadRequest() *GetAccountsAccountIDDirectDebitsBadRequest {
	return &GetAccountsAccountIDDirectDebitsBadRequest{}
}

/*GetAccountsAccountIDDirectDebitsBadRequest handles this case with default header values.

Bad request
*/
type GetAccountsAccountIDDirectDebitsBadRequest struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDDirectDebitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAccountsAccountIDDirectDebitsBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDDirectDebitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDDirectDebitsUnauthorized creates a GetAccountsAccountIDDirectDebitsUnauthorized with default headers values
func NewGetAccountsAccountIDDirectDebitsUnauthorized() *GetAccountsAccountIDDirectDebitsUnauthorized {
	return &GetAccountsAccountIDDirectDebitsUnauthorized{}
}

/*GetAccountsAccountIDDirectDebitsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountsAccountIDDirectDebitsUnauthorized struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDDirectDebitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsUnauthorized ", 401)
}

func (o *GetAccountsAccountIDDirectDebitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDDirectDebitsForbidden creates a GetAccountsAccountIDDirectDebitsForbidden with default headers values
func NewGetAccountsAccountIDDirectDebitsForbidden() *GetAccountsAccountIDDirectDebitsForbidden {
	return &GetAccountsAccountIDDirectDebitsForbidden{}
}

/*GetAccountsAccountIDDirectDebitsForbidden handles this case with default header values.

Forbidden
*/
type GetAccountsAccountIDDirectDebitsForbidden struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDDirectDebitsForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsForbidden  %+v", 403, o.Payload)
}

func (o *GetAccountsAccountIDDirectDebitsForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDDirectDebitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDDirectDebitsNotFound creates a GetAccountsAccountIDDirectDebitsNotFound with default headers values
func NewGetAccountsAccountIDDirectDebitsNotFound() *GetAccountsAccountIDDirectDebitsNotFound {
	return &GetAccountsAccountIDDirectDebitsNotFound{}
}

/*GetAccountsAccountIDDirectDebitsNotFound handles this case with default header values.

Not found
*/
type GetAccountsAccountIDDirectDebitsNotFound struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDDirectDebitsNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsNotFound ", 404)
}

func (o *GetAccountsAccountIDDirectDebitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDDirectDebitsMethodNotAllowed creates a GetAccountsAccountIDDirectDebitsMethodNotAllowed with default headers values
func NewGetAccountsAccountIDDirectDebitsMethodNotAllowed() *GetAccountsAccountIDDirectDebitsMethodNotAllowed {
	return &GetAccountsAccountIDDirectDebitsMethodNotAllowed{}
}

/*GetAccountsAccountIDDirectDebitsMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDDirectDebitsMethodNotAllowed struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDDirectDebitsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDDirectDebitsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDDirectDebitsNotAcceptable creates a GetAccountsAccountIDDirectDebitsNotAcceptable with default headers values
func NewGetAccountsAccountIDDirectDebitsNotAcceptable() *GetAccountsAccountIDDirectDebitsNotAcceptable {
	return &GetAccountsAccountIDDirectDebitsNotAcceptable{}
}

/*GetAccountsAccountIDDirectDebitsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GetAccountsAccountIDDirectDebitsNotAcceptable struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDDirectDebitsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDDirectDebitsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	return nil
}

// NewGetAccountsAccountIDDirectDebitsTooManyRequests creates a GetAccountsAccountIDDirectDebitsTooManyRequests with default headers values
func NewGetAccountsAccountIDDirectDebitsTooManyRequests() *GetAccountsAccountIDDirectDebitsTooManyRequests {
	return &GetAccountsAccountIDDirectDebitsTooManyRequests{}
}

/*GetAccountsAccountIDDirectDebitsTooManyRequests handles this case with default header values.

Too Many Requests
*/
type GetAccountsAccountIDDirectDebitsTooManyRequests struct {
	/*Number in seconds to wait
	 */
	RetryAfter int64
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDDirectDebitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDDirectDebitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDDirectDebitsInternalServerError creates a GetAccountsAccountIDDirectDebitsInternalServerError with default headers values
func NewGetAccountsAccountIDDirectDebitsInternalServerError() *GetAccountsAccountIDDirectDebitsInternalServerError {
	return &GetAccountsAccountIDDirectDebitsInternalServerError{}
}

/*GetAccountsAccountIDDirectDebitsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountsAccountIDDirectDebitsInternalServerError struct {
	/*An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDDirectDebitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/direct-debits][%d] getAccountsAccountIdDirectDebitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAccountsAccountIDDirectDebitsInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDDirectDebitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-fapi-interaction-id
	o.XFapiInteractionID = response.GetHeader("x-fapi-interaction-id")

	o.Payload = new(models.OBErrorResponse1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}