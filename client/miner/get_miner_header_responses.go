// Code generated by go-swagger; DO NOT EDIT.

package miner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/models"
)

// GetMinerHeaderReader is a Reader for the GetMinerHeader structure.
type GetMinerHeaderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMinerHeaderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMinerHeaderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMinerHeaderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMinerHeaderOK creates a GetMinerHeaderOK with default headers values
func NewGetMinerHeaderOK() *GetMinerHeaderOK {
	return &GetMinerHeaderOK{}
}

/* GetMinerHeaderOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetMinerHeaderOK struct {
}

func (o *GetMinerHeaderOK) Error() string {
	return fmt.Sprintf("[GET /miner/header][%d] getMinerHeaderOK ", 200)
}

func (o *GetMinerHeaderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMinerHeaderDefault creates a GetMinerHeaderDefault with default headers values
func NewGetMinerHeaderDefault(code int) *GetMinerHeaderDefault {
	return &GetMinerHeaderDefault{
		_statusCode: code,
	}
}

/* GetMinerHeaderDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetMinerHeaderDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get miner header default response
func (o *GetMinerHeaderDefault) Code() int {
	return o._statusCode
}

func (o *GetMinerHeaderDefault) Error() string {
	return fmt.Sprintf("[GET /miner/header][%d] GetMinerHeader default  %+v", o._statusCode, o.Payload)
}
func (o *GetMinerHeaderDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetMinerHeaderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
