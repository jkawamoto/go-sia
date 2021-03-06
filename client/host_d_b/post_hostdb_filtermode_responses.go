// Code generated by go-swagger; DO NOT EDIT.

package host_d_b

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jkawamoto/go-sia/models"
)

// PostHostdbFiltermodeReader is a Reader for the PostHostdbFiltermode structure.
type PostHostdbFiltermodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostdbFiltermodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostHostdbFiltermodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostHostdbFiltermodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostHostdbFiltermodeNoContent creates a PostHostdbFiltermodeNoContent with default headers values
func NewPostHostdbFiltermodeNoContent() *PostHostdbFiltermodeNoContent {
	return &PostHostdbFiltermodeNoContent{}
}

/* PostHostdbFiltermodeNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type PostHostdbFiltermodeNoContent struct {
}

func (o *PostHostdbFiltermodeNoContent) Error() string {
	return fmt.Sprintf("[POST /hostdb/filtermode][%d] postHostdbFiltermodeNoContent ", 204)
}

func (o *PostHostdbFiltermodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostHostdbFiltermodeDefault creates a PostHostdbFiltermodeDefault with default headers values
func NewPostHostdbFiltermodeDefault(code int) *PostHostdbFiltermodeDefault {
	return &PostHostdbFiltermodeDefault{
		_statusCode: code,
	}
}

/* PostHostdbFiltermodeDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostHostdbFiltermodeDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post hostdb filtermode default response
func (o *PostHostdbFiltermodeDefault) Code() int {
	return o._statusCode
}

func (o *PostHostdbFiltermodeDefault) Error() string {
	return fmt.Sprintf("[POST /hostdb/filtermode][%d] PostHostdbFiltermode default  %+v", o._statusCode, o.Payload)
}
func (o *PostHostdbFiltermodeDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostHostdbFiltermodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostHostdbFiltermodeBody post hostdb filtermode body
swagger:model PostHostdbFiltermodeBody
*/
type PostHostdbFiltermodeBody struct {

	// Can be either whitelist, blacklist, or disable.
	// Required: true
	Filtermode *string `json:"filtermode"`

	// pubkeys.
	Hosts []string `json:"hosts"`
}

// Validate validates this post hostdb filtermode body
func (o *PostHostdbFiltermodeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFiltermode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostHostdbFiltermodeBody) validateFiltermode(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"filtermode", "body", o.Filtermode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post hostdb filtermode body based on context it is used
func (o *PostHostdbFiltermodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostHostdbFiltermodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostHostdbFiltermodeBody) UnmarshalBinary(b []byte) error {
	var res PostHostdbFiltermodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
