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

// PostWalletInitSeedReader is a Reader for the PostWalletInitSeed structure.
type PostWalletInitSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWalletInitSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostWalletInitSeedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWalletInitSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWalletInitSeedNoContent creates a PostWalletInitSeedNoContent with default headers values
func NewPostWalletInitSeedNoContent() *PostWalletInitSeedNoContent {
	return &PostWalletInitSeedNoContent{}
}

/*PostWalletInitSeedNoContent handles this case with default header values.

Successful Response
*/
type PostWalletInitSeedNoContent struct {
}

func (o *PostWalletInitSeedNoContent) Error() string {
	return fmt.Sprintf("[POST /wallet/init/seed][%d] postWalletInitSeedNoContent ", 204)
}

func (o *PostWalletInitSeedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWalletInitSeedDefault creates a PostWalletInitSeedDefault with default headers values
func NewPostWalletInitSeedDefault(code int) *PostWalletInitSeedDefault {
	return &PostWalletInitSeedDefault{
		_statusCode: code,
	}
}

/*PostWalletInitSeedDefault handles this case with default header values.

Error Response
*/
type PostWalletInitSeedDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post wallet init seed default response
func (o *PostWalletInitSeedDefault) Code() int {
	return o._statusCode
}

func (o *PostWalletInitSeedDefault) Error() string {
	return fmt.Sprintf("[POST /wallet/init/seed][%d] PostWalletInitSeed default  %+v", o._statusCode, o.Payload)
}

func (o *PostWalletInitSeedDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostWalletInitSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
