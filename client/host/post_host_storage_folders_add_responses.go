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

// PostHostStorageFoldersAddReader is a Reader for the PostHostStorageFoldersAdd structure.
type PostHostStorageFoldersAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostStorageFoldersAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostHostStorageFoldersAddNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostHostStorageFoldersAddDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHostStorageFoldersAddNoContent creates a PostHostStorageFoldersAddNoContent with default headers values
func NewPostHostStorageFoldersAddNoContent() *PostHostStorageFoldersAddNoContent {
	return &PostHostStorageFoldersAddNoContent{}
}

/*PostHostStorageFoldersAddNoContent handles this case with default header values.

Successful Response
*/
type PostHostStorageFoldersAddNoContent struct {
}

func (o *PostHostStorageFoldersAddNoContent) Error() string {
	return fmt.Sprintf("[POST /host/storage/folders/add][%d] postHostStorageFoldersAddNoContent ", 204)
}

func (o *PostHostStorageFoldersAddNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHostStorageFoldersAddDefault creates a PostHostStorageFoldersAddDefault with default headers values
func NewPostHostStorageFoldersAddDefault(code int) *PostHostStorageFoldersAddDefault {
	return &PostHostStorageFoldersAddDefault{
		_statusCode: code,
	}
}

/*PostHostStorageFoldersAddDefault handles this case with default header values.

Error Response
*/
type PostHostStorageFoldersAddDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post host storage folders add default response
func (o *PostHostStorageFoldersAddDefault) Code() int {
	return o._statusCode
}

func (o *PostHostStorageFoldersAddDefault) Error() string {
	return fmt.Sprintf("[POST /host/storage/folders/add][%d] PostHostStorageFoldersAdd default  %+v", o._statusCode, o.Payload)
}

func (o *PostHostStorageFoldersAddDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
