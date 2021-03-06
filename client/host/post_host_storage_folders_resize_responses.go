// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/models"
)

// PostHostStorageFoldersResizeReader is a Reader for the PostHostStorageFoldersResize structure.
type PostHostStorageFoldersResizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostStorageFoldersResizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostHostStorageFoldersResizeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHostStorageFoldersResizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHostStorageFoldersResizeNoContent creates a PostHostStorageFoldersResizeNoContent with default headers values
func NewPostHostStorageFoldersResizeNoContent() *PostHostStorageFoldersResizeNoContent {
	return &PostHostStorageFoldersResizeNoContent{}
}

/* PostHostStorageFoldersResizeNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type PostHostStorageFoldersResizeNoContent struct {
}

func (o *PostHostStorageFoldersResizeNoContent) Error() string {
	return fmt.Sprintf("[POST /host/storage/folders/resize][%d] postHostStorageFoldersResizeNoContent ", 204)
}

func (o *PostHostStorageFoldersResizeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHostStorageFoldersResizeDefault creates a PostHostStorageFoldersResizeDefault with default headers values
func NewPostHostStorageFoldersResizeDefault(code int) *PostHostStorageFoldersResizeDefault {
	return &PostHostStorageFoldersResizeDefault{
		_statusCode: code,
	}
}

/* PostHostStorageFoldersResizeDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostHostStorageFoldersResizeDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post host storage folders resize default response
func (o *PostHostStorageFoldersResizeDefault) Code() int {
	return o._statusCode
}

func (o *PostHostStorageFoldersResizeDefault) Error() string {
	return fmt.Sprintf("[POST /host/storage/folders/resize][%d] PostHostStorageFoldersResize default  %+v", o._statusCode, o.Payload)
}
func (o *PostHostStorageFoldersResizeDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostHostStorageFoldersResizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
