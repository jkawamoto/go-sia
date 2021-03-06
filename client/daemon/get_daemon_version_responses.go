// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/jkawamoto/go-sia/models"
)

// GetDaemonVersionReader is a Reader for the GetDaemonVersion structure.
type GetDaemonVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDaemonVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDaemonVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDaemonVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDaemonVersionOK creates a GetDaemonVersionOK with default headers values
func NewGetDaemonVersionOK() *GetDaemonVersionOK {
	return &GetDaemonVersionOK{}
}

/* GetDaemonVersionOK describes a response with status code 200, with default header values.

returns the version of the Sia daemon currently running.
*/
type GetDaemonVersionOK struct {
	Payload *GetDaemonVersionOKBody
}

func (o *GetDaemonVersionOK) Error() string {
	return fmt.Sprintf("[GET /daemon/version][%d] getDaemonVersionOK  %+v", 200, o.Payload)
}
func (o *GetDaemonVersionOK) GetPayload() *GetDaemonVersionOKBody {
	return o.Payload
}

func (o *GetDaemonVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDaemonVersionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDaemonVersionDefault creates a GetDaemonVersionDefault with default headers values
func NewGetDaemonVersionDefault(code int) *GetDaemonVersionDefault {
	return &GetDaemonVersionDefault{
		_statusCode: code,
	}
}

/* GetDaemonVersionDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetDaemonVersionDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get daemon version default response
func (o *GetDaemonVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetDaemonVersionDefault) Error() string {
	return fmt.Sprintf("[GET /daemon/version][%d] GetDaemonVersion default  %+v", o._statusCode, o.Payload)
}
func (o *GetDaemonVersionDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetDaemonVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetDaemonVersionOKBody get daemon version o k body
swagger:model GetDaemonVersionOKBody
*/
type GetDaemonVersionOKBody struct {

	// Version number of the running Sia Daemon. This number is visible to its
	// peers on the network.
	//
	// Example: 1.0.0
	Version string `json:"version,omitempty"`
}

// Validate validates this get daemon version o k body
func (o *GetDaemonVersionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get daemon version o k body based on context it is used
func (o *GetDaemonVersionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDaemonVersionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDaemonVersionOKBody) UnmarshalBinary(b []byte) error {
	var res GetDaemonVersionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
