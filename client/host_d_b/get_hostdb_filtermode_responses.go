// Code generated by go-swagger; DO NOT EDIT.

package host_d_b

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

// GetHostdbFiltermodeReader is a Reader for the GetHostdbFiltermode structure.
type GetHostdbFiltermodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostdbFiltermodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostdbFiltermodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHostdbFiltermodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHostdbFiltermodeOK creates a GetHostdbFiltermodeOK with default headers values
func NewGetHostdbFiltermodeOK() *GetHostdbFiltermodeOK {
	return &GetHostdbFiltermodeOK{}
}

/* GetHostdbFiltermodeOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetHostdbFiltermodeOK struct {
	Payload *GetHostdbFiltermodeOKBody
}

func (o *GetHostdbFiltermodeOK) Error() string {
	return fmt.Sprintf("[GET /hostdb/filtermode][%d] getHostdbFiltermodeOK  %+v", 200, o.Payload)
}
func (o *GetHostdbFiltermodeOK) GetPayload() *GetHostdbFiltermodeOKBody {
	return o.Payload
}

func (o *GetHostdbFiltermodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHostdbFiltermodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostdbFiltermodeDefault creates a GetHostdbFiltermodeDefault with default headers values
func NewGetHostdbFiltermodeDefault(code int) *GetHostdbFiltermodeDefault {
	return &GetHostdbFiltermodeDefault{
		_statusCode: code,
	}
}

/* GetHostdbFiltermodeDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetHostdbFiltermodeDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get hostdb filtermode default response
func (o *GetHostdbFiltermodeDefault) Code() int {
	return o._statusCode
}

func (o *GetHostdbFiltermodeDefault) Error() string {
	return fmt.Sprintf("[GET /hostdb/filtermode][%d] GetHostdbFiltermode default  %+v", o._statusCode, o.Payload)
}
func (o *GetHostdbFiltermodeDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetHostdbFiltermodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHostdbFiltermodeOKBody get hostdb filtermode o k body
swagger:model GetHostdbFiltermodeOKBody
*/
type GetHostdbFiltermodeOKBody struct {

	// Can be either whitelist, blacklist, or disable.
	Filtermode string `json:"filtermode,omitempty"`

	// array of strings Comma separated pubkeys.
	Hosts []string `json:"hosts"`
}

// Validate validates this get hostdb filtermode o k body
func (o *GetHostdbFiltermodeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get hostdb filtermode o k body based on context it is used
func (o *GetHostdbFiltermodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHostdbFiltermodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostdbFiltermodeOKBody) UnmarshalBinary(b []byte) error {
	var res GetHostdbFiltermodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
