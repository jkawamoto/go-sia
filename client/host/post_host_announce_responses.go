// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// PostHostAnnounceReader is a Reader for the PostHostAnnounce structure.
type PostHostAnnounceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostAnnounceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostHostAnnounceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostHostAnnounceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHostAnnounceNoContent creates a PostHostAnnounceNoContent with default headers values
func NewPostHostAnnounceNoContent() *PostHostAnnounceNoContent {
	return &PostHostAnnounceNoContent{}
}

/*PostHostAnnounceNoContent handles this case with default header values.

Successful Response
*/
type PostHostAnnounceNoContent struct {
}

func (o *PostHostAnnounceNoContent) Error() string {
	return fmt.Sprintf("[POST /host/announce][%d] postHostAnnounceNoContent ", 204)
}

func (o *PostHostAnnounceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHostAnnounceDefault creates a PostHostAnnounceDefault with default headers values
func NewPostHostAnnounceDefault(code int) *PostHostAnnounceDefault {
	return &PostHostAnnounceDefault{
		_statusCode: code,
	}
}

/*PostHostAnnounceDefault handles this case with default header values.

Error Response
*/
type PostHostAnnounceDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post host announce default response
func (o *PostHostAnnounceDefault) Code() int {
	return o._statusCode
}

func (o *PostHostAnnounceDefault) Error() string {
	return fmt.Sprintf("[POST /host/announce][%d] PostHostAnnounce default  %+v", o._statusCode, o.Payload)
}

func (o *PostHostAnnounceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
