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

// PostRenterBackupsRestoreReader is a Reader for the PostRenterBackupsRestore structure.
type PostRenterBackupsRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRenterBackupsRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostRenterBackupsRestoreNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostRenterBackupsRestoreDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostRenterBackupsRestoreNoContent creates a PostRenterBackupsRestoreNoContent with default headers values
func NewPostRenterBackupsRestoreNoContent() *PostRenterBackupsRestoreNoContent {
	return &PostRenterBackupsRestoreNoContent{}
}

/* PostRenterBackupsRestoreNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type PostRenterBackupsRestoreNoContent struct {
}

func (o *PostRenterBackupsRestoreNoContent) Error() string {
	return fmt.Sprintf("[POST /renter/backups/restore][%d] postRenterBackupsRestoreNoContent ", 204)
}

func (o *PostRenterBackupsRestoreNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRenterBackupsRestoreDefault creates a PostRenterBackupsRestoreDefault with default headers values
func NewPostRenterBackupsRestoreDefault(code int) *PostRenterBackupsRestoreDefault {
	return &PostRenterBackupsRestoreDefault{
		_statusCode: code,
	}
}

/* PostRenterBackupsRestoreDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostRenterBackupsRestoreDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post renter backups restore default response
func (o *PostRenterBackupsRestoreDefault) Code() int {
	return o._statusCode
}

func (o *PostRenterBackupsRestoreDefault) Error() string {
	return fmt.Sprintf("[POST /renter/backups/restore][%d] PostRenterBackupsRestore default  %+v", o._statusCode, o.Payload)
}
func (o *PostRenterBackupsRestoreDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostRenterBackupsRestoreDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
