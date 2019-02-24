// Code generated by go-swagger; DO NOT EDIT.

package renter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// GetRenterDownloadasyncSiapathReader is a Reader for the GetRenterDownloadasyncSiapath structure.
type GetRenterDownloadasyncSiapathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRenterDownloadasyncSiapathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewGetRenterDownloadasyncSiapathNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetRenterDownloadasyncSiapathDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRenterDownloadasyncSiapathNoContent creates a GetRenterDownloadasyncSiapathNoContent with default headers values
func NewGetRenterDownloadasyncSiapathNoContent() *GetRenterDownloadasyncSiapathNoContent {
	return &GetRenterDownloadasyncSiapathNoContent{}
}

/*GetRenterDownloadasyncSiapathNoContent handles this case with default header values.

Successful Response
*/
type GetRenterDownloadasyncSiapathNoContent struct {
}

func (o *GetRenterDownloadasyncSiapathNoContent) Error() string {
	return fmt.Sprintf("[GET /renter/downloadasync/{siapath}][%d] getRenterDownloadasyncSiapathNoContent ", 204)
}

func (o *GetRenterDownloadasyncSiapathNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRenterDownloadasyncSiapathDefault creates a GetRenterDownloadasyncSiapathDefault with default headers values
func NewGetRenterDownloadasyncSiapathDefault(code int) *GetRenterDownloadasyncSiapathDefault {
	return &GetRenterDownloadasyncSiapathDefault{
		_statusCode: code,
	}
}

/*GetRenterDownloadasyncSiapathDefault handles this case with default header values.

Error Response
*/
type GetRenterDownloadasyncSiapathDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get renter downloadasync siapath default response
func (o *GetRenterDownloadasyncSiapathDefault) Code() int {
	return o._statusCode
}

func (o *GetRenterDownloadasyncSiapathDefault) Error() string {
	return fmt.Sprintf("[GET /renter/downloadasync/{siapath}][%d] GetRenterDownloadasyncSiapath default  %+v", o._statusCode, o.Payload)
}

func (o *GetRenterDownloadasyncSiapathDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}