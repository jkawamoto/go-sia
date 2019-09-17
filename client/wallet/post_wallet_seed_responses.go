// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// PostWalletSeedReader is a Reader for the PostWalletSeed structure.
type PostWalletSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWalletSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostWalletSeedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWalletSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWalletSeedNoContent creates a PostWalletSeedNoContent with default headers values
func NewPostWalletSeedNoContent() *PostWalletSeedNoContent {
	return &PostWalletSeedNoContent{}
}

/*PostWalletSeedNoContent handles this case with default header values.

Successful Response
*/
type PostWalletSeedNoContent struct {
}

func (o *PostWalletSeedNoContent) Error() string {
	return fmt.Sprintf("[POST /wallet/seed][%d] postWalletSeedNoContent ", 204)
}

func (o *PostWalletSeedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWalletSeedDefault creates a PostWalletSeedDefault with default headers values
func NewPostWalletSeedDefault(code int) *PostWalletSeedDefault {
	return &PostWalletSeedDefault{
		_statusCode: code,
	}
}

/*PostWalletSeedDefault handles this case with default header values.

Error Response
*/
type PostWalletSeedDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post wallet seed default response
func (o *PostWalletSeedDefault) Code() int {
	return o._statusCode
}

func (o *PostWalletSeedDefault) Error() string {
	return fmt.Sprintf("[POST /wallet/seed][%d] PostWalletSeed default  %+v", o._statusCode, o.Payload)
}

func (o *PostWalletSeedDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostWalletSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
