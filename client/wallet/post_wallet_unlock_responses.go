// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/models"
)

// PostWalletUnlockReader is a Reader for the PostWalletUnlock structure.
type PostWalletUnlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWalletUnlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostWalletUnlockNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWalletUnlockDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWalletUnlockNoContent creates a PostWalletUnlockNoContent with default headers values
func NewPostWalletUnlockNoContent() *PostWalletUnlockNoContent {
	return &PostWalletUnlockNoContent{}
}

/* PostWalletUnlockNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type PostWalletUnlockNoContent struct {
}

func (o *PostWalletUnlockNoContent) Error() string {
	return fmt.Sprintf("[POST /wallet/unlock][%d] postWalletUnlockNoContent ", 204)
}

func (o *PostWalletUnlockNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWalletUnlockDefault creates a PostWalletUnlockDefault with default headers values
func NewPostWalletUnlockDefault(code int) *PostWalletUnlockDefault {
	return &PostWalletUnlockDefault{
		_statusCode: code,
	}
}

/* PostWalletUnlockDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostWalletUnlockDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post wallet unlock default response
func (o *PostWalletUnlockDefault) Code() int {
	return o._statusCode
}

func (o *PostWalletUnlockDefault) Error() string {
	return fmt.Sprintf("[POST /wallet/unlock][%d] PostWalletUnlock default  %+v", o._statusCode, o.Payload)
}
func (o *PostWalletUnlockDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostWalletUnlockDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
