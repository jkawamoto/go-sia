// Code generated by go-swagger; DO NOT EDIT.

package miner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// PostMinerHeaderReader is a Reader for the PostMinerHeader structure.
type PostMinerHeaderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMinerHeaderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMinerHeaderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostMinerHeaderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostMinerHeaderOK creates a PostMinerHeaderOK with default headers values
func NewPostMinerHeaderOK() *PostMinerHeaderOK {
	return &PostMinerHeaderOK{}
}

/*PostMinerHeaderOK handles this case with default header values.

Successful Response
*/
type PostMinerHeaderOK struct {
}

func (o *PostMinerHeaderOK) Error() string {
	return fmt.Sprintf("[POST /miner/header][%d] postMinerHeaderOK ", 200)
}

func (o *PostMinerHeaderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostMinerHeaderDefault creates a PostMinerHeaderDefault with default headers values
func NewPostMinerHeaderDefault(code int) *PostMinerHeaderDefault {
	return &PostMinerHeaderDefault{
		_statusCode: code,
	}
}

/*PostMinerHeaderDefault handles this case with default header values.

Error Response
*/
type PostMinerHeaderDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post miner header default response
func (o *PostMinerHeaderDefault) Code() int {
	return o._statusCode
}

func (o *PostMinerHeaderDefault) Error() string {
	return fmt.Sprintf("[POST /miner/header][%d] PostMinerHeader default  %+v", o._statusCode, o.Payload)
}

func (o *PostMinerHeaderDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostMinerHeaderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
