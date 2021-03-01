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

	"github.com/cloudentity/openbanking-sandbox/openbanking/accountinformation/models"
)

// GetAccountsAccountIDPartyReader is a Reader for the GetAccountsAccountIDParty structure.
type GetAccountsAccountIDPartyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsAccountIDPartyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountsAccountIDPartyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountsAccountIDPartyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAccountsAccountIDPartyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAccountsAccountIDPartyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountsAccountIDPartyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetAccountsAccountIDPartyMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 406:
		result := NewGetAccountsAccountIDPartyNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAccountsAccountIDPartyTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAccountsAccountIDPartyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountsAccountIDPartyOK creates a GetAccountsAccountIDPartyOK with default headers values
func NewGetAccountsAccountIDPartyOK() *GetAccountsAccountIDPartyOK {
	return &GetAccountsAccountIDPartyOK{}
}

/* GetAccountsAccountIDPartyOK describes a response with status code 200, with default header values.

Parties Read
*/
type GetAccountsAccountIDPartyOK struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBReadParty2
}

func (o *GetAccountsAccountIDPartyOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyOK  %+v", 200, o.Payload)
}
func (o *GetAccountsAccountIDPartyOK) GetPayload() *models.OBReadParty2 {
	return o.Payload
}

func (o *GetAccountsAccountIDPartyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	o.Payload = new(models.OBReadParty2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountsAccountIDPartyBadRequest creates a GetAccountsAccountIDPartyBadRequest with default headers values
func NewGetAccountsAccountIDPartyBadRequest() *GetAccountsAccountIDPartyBadRequest {
	return &GetAccountsAccountIDPartyBadRequest{}
}

/* GetAccountsAccountIDPartyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAccountsAccountIDPartyBadRequest struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDPartyBadRequest) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyBadRequest  %+v", 400, o.Payload)
}
func (o *GetAccountsAccountIDPartyBadRequest) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDPartyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDPartyUnauthorized creates a GetAccountsAccountIDPartyUnauthorized with default headers values
func NewGetAccountsAccountIDPartyUnauthorized() *GetAccountsAccountIDPartyUnauthorized {
	return &GetAccountsAccountIDPartyUnauthorized{}
}

/* GetAccountsAccountIDPartyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAccountsAccountIDPartyUnauthorized struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDPartyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyUnauthorized ", 401)
}

func (o *GetAccountsAccountIDPartyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDPartyForbidden creates a GetAccountsAccountIDPartyForbidden with default headers values
func NewGetAccountsAccountIDPartyForbidden() *GetAccountsAccountIDPartyForbidden {
	return &GetAccountsAccountIDPartyForbidden{}
}

/* GetAccountsAccountIDPartyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAccountsAccountIDPartyForbidden struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDPartyForbidden) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyForbidden  %+v", 403, o.Payload)
}
func (o *GetAccountsAccountIDPartyForbidden) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDPartyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDPartyNotFound creates a GetAccountsAccountIDPartyNotFound with default headers values
func NewGetAccountsAccountIDPartyNotFound() *GetAccountsAccountIDPartyNotFound {
	return &GetAccountsAccountIDPartyNotFound{}
}

/* GetAccountsAccountIDPartyNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAccountsAccountIDPartyNotFound struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDPartyNotFound) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyNotFound ", 404)
}

func (o *GetAccountsAccountIDPartyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDPartyMethodNotAllowed creates a GetAccountsAccountIDPartyMethodNotAllowed with default headers values
func NewGetAccountsAccountIDPartyMethodNotAllowed() *GetAccountsAccountIDPartyMethodNotAllowed {
	return &GetAccountsAccountIDPartyMethodNotAllowed{}
}

/* GetAccountsAccountIDPartyMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetAccountsAccountIDPartyMethodNotAllowed struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDPartyMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyMethodNotAllowed ", 405)
}

func (o *GetAccountsAccountIDPartyMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDPartyNotAcceptable creates a GetAccountsAccountIDPartyNotAcceptable with default headers values
func NewGetAccountsAccountIDPartyNotAcceptable() *GetAccountsAccountIDPartyNotAcceptable {
	return &GetAccountsAccountIDPartyNotAcceptable{}
}

/* GetAccountsAccountIDPartyNotAcceptable describes a response with status code 406, with default header values.

Not Acceptable
*/
type GetAccountsAccountIDPartyNotAcceptable struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDPartyNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyNotAcceptable ", 406)
}

func (o *GetAccountsAccountIDPartyNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-fapi-interaction-id
	hdrXFapiInteractionID := response.GetHeader("x-fapi-interaction-id")

	if hdrXFapiInteractionID != "" {
		o.XFapiInteractionID = hdrXFapiInteractionID
	}

	return nil
}

// NewGetAccountsAccountIDPartyTooManyRequests creates a GetAccountsAccountIDPartyTooManyRequests with default headers values
func NewGetAccountsAccountIDPartyTooManyRequests() *GetAccountsAccountIDPartyTooManyRequests {
	return &GetAccountsAccountIDPartyTooManyRequests{}
}

/* GetAccountsAccountIDPartyTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetAccountsAccountIDPartyTooManyRequests struct {

	/* Number in seconds to wait
	 */
	RetryAfter int64

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string
}

func (o *GetAccountsAccountIDPartyTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyTooManyRequests ", 429)
}

func (o *GetAccountsAccountIDPartyTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAccountsAccountIDPartyInternalServerError creates a GetAccountsAccountIDPartyInternalServerError with default headers values
func NewGetAccountsAccountIDPartyInternalServerError() *GetAccountsAccountIDPartyInternalServerError {
	return &GetAccountsAccountIDPartyInternalServerError{}
}

/* GetAccountsAccountIDPartyInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAccountsAccountIDPartyInternalServerError struct {

	/* An RFC4122 UID used as a correlation id.
	 */
	XFapiInteractionID string

	Payload *models.OBErrorResponse1
}

func (o *GetAccountsAccountIDPartyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{AccountId}/party][%d] getAccountsAccountIdPartyInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAccountsAccountIDPartyInternalServerError) GetPayload() *models.OBErrorResponse1 {
	return o.Payload
}

func (o *GetAccountsAccountIDPartyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
