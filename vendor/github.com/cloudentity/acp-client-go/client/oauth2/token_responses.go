// Code generated by go-swagger; DO NOT EDIT.

package oauth2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudentity/acp-client-go/models"
)

// TokenReader is a Reader for the Token structure.
type TokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTokenOK creates a TokenOK with default headers values
func NewTokenOK() *TokenOK {
	return &TokenOK{}
}

/*TokenOK handles this case with default header values.

tokenResponse
*/
type TokenOK struct {
	Payload *models.TokenResponse
}

func (o *TokenOK) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/token][%d] tokenOK  %+v", 200, o.Payload)
}

func (o *TokenOK) GetPayload() *models.TokenResponse {
	return o.Payload
}

func (o *TokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenUnauthorized creates a TokenUnauthorized with default headers values
func NewTokenUnauthorized() *TokenUnauthorized {
	return &TokenUnauthorized{}
}

/*TokenUnauthorized handles this case with default header values.

genericError
*/
type TokenUnauthorized struct {
	Payload *models.GenericError
}

func (o *TokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/token][%d] tokenUnauthorized  %+v", 401, o.Payload)
}

func (o *TokenUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *TokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenNotFound creates a TokenNotFound with default headers values
func NewTokenNotFound() *TokenNotFound {
	return &TokenNotFound{}
}

/*TokenNotFound handles this case with default header values.

genericError
*/
type TokenNotFound struct {
	Payload *models.GenericError
}

func (o *TokenNotFound) Error() string {
	return fmt.Sprintf("[POST /{tid}/{aid}/oauth2/token][%d] tokenNotFound  %+v", 404, o.Payload)
}

func (o *TokenNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *TokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
