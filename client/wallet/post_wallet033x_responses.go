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

// PostWallet033xReader is a Reader for the PostWallet033x structure.
type PostWallet033xReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWallet033xReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostWallet033xNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostWallet033xDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWallet033xNoContent creates a PostWallet033xNoContent with default headers values
func NewPostWallet033xNoContent() *PostWallet033xNoContent {
	return &PostWallet033xNoContent{}
}

/*PostWallet033xNoContent handles this case with default header values.

Successful Response
*/
type PostWallet033xNoContent struct {
}

func (o *PostWallet033xNoContent) Error() string {
	return fmt.Sprintf("[POST /wallet/033x][%d] postWallet033xNoContent ", 204)
}

func (o *PostWallet033xNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWallet033xDefault creates a PostWallet033xDefault with default headers values
func NewPostWallet033xDefault(code int) *PostWallet033xDefault {
	return &PostWallet033xDefault{
		_statusCode: code,
	}
}

/*PostWallet033xDefault handles this case with default header values.

Error Response
*/
type PostWallet033xDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post wallet033x default response
func (o *PostWallet033xDefault) Code() int {
	return o._statusCode
}

func (o *PostWallet033xDefault) Error() string {
	return fmt.Sprintf("[POST /wallet/033x][%d] PostWallet033x default  %+v", o._statusCode, o.Payload)
}

func (o *PostWallet033xDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
