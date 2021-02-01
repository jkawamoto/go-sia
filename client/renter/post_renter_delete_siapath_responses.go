// Code generated by go-swagger; DO NOT EDIT.

package renter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/models"
)

// PostRenterDeleteSiapathReader is a Reader for the PostRenterDeleteSiapath structure.
type PostRenterDeleteSiapathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRenterDeleteSiapathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostRenterDeleteSiapathNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostRenterDeleteSiapathDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostRenterDeleteSiapathNoContent creates a PostRenterDeleteSiapathNoContent with default headers values
func NewPostRenterDeleteSiapathNoContent() *PostRenterDeleteSiapathNoContent {
	return &PostRenterDeleteSiapathNoContent{}
}

/* PostRenterDeleteSiapathNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type PostRenterDeleteSiapathNoContent struct {
}

func (o *PostRenterDeleteSiapathNoContent) Error() string {
	return fmt.Sprintf("[POST /renter/delete/{siapath}][%d] postRenterDeleteSiapathNoContent ", 204)
}

func (o *PostRenterDeleteSiapathNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRenterDeleteSiapathDefault creates a PostRenterDeleteSiapathDefault with default headers values
func NewPostRenterDeleteSiapathDefault(code int) *PostRenterDeleteSiapathDefault {
	return &PostRenterDeleteSiapathDefault{
		_statusCode: code,
	}
}

/* PostRenterDeleteSiapathDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostRenterDeleteSiapathDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post renter delete siapath default response
func (o *PostRenterDeleteSiapathDefault) Code() int {
	return o._statusCode
}

func (o *PostRenterDeleteSiapathDefault) Error() string {
	return fmt.Sprintf("[POST /renter/delete/{siapath}][%d] PostRenterDeleteSiapath default  %+v", o._statusCode, o.Payload)
}
func (o *PostRenterDeleteSiapathDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostRenterDeleteSiapathDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
