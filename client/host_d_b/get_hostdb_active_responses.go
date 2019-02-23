// Code generated by go-swagger; DO NOT EDIT.

package host_d_b

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// GetHostdbActiveReader is a Reader for the GetHostdbActive structure.
type GetHostdbActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostdbActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHostdbActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetHostdbActiveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHostdbActiveOK creates a GetHostdbActiveOK with default headers values
func NewGetHostdbActiveOK() *GetHostdbActiveOK {
	return &GetHostdbActiveOK{}
}

/*GetHostdbActiveOK handles this case with default header values.

Successful Response
*/
type GetHostdbActiveOK struct {
	Payload *models.Hostdb
}

func (o *GetHostdbActiveOK) Error() string {
	return fmt.Sprintf("[GET /hostdb/active][%d] getHostdbActiveOK  %+v", 200, o.Payload)
}

func (o *GetHostdbActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Hostdb)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostdbActiveDefault creates a GetHostdbActiveDefault with default headers values
func NewGetHostdbActiveDefault(code int) *GetHostdbActiveDefault {
	return &GetHostdbActiveDefault{
		_statusCode: code,
	}
}

/*GetHostdbActiveDefault handles this case with default header values.

Error Response
*/
type GetHostdbActiveDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get hostdb active default response
func (o *GetHostdbActiveDefault) Code() int {
	return o._statusCode
}

func (o *GetHostdbActiveDefault) Error() string {
	return fmt.Sprintf("[GET /hostdb/active][%d] GetHostdbActive default  %+v", o._statusCode, o.Payload)
}

func (o *GetHostdbActiveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
