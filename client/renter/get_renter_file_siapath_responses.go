// Code generated by go-swagger; DO NOT EDIT.

package renter

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

	"github.com/jkawamoto/go-sia/models"
)

// GetRenterFileSiapathReader is a Reader for the GetRenterFileSiapath structure.
type GetRenterFileSiapathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRenterFileSiapathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRenterFileSiapathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRenterFileSiapathDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRenterFileSiapathOK creates a GetRenterFileSiapathOK with default headers values
func NewGetRenterFileSiapathOK() *GetRenterFileSiapathOK {
	return &GetRenterFileSiapathOK{}
}

/* GetRenterFileSiapathOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetRenterFileSiapathOK struct {
	Payload *GetRenterFileSiapathOKBody
}

func (o *GetRenterFileSiapathOK) Error() string {
	return fmt.Sprintf("[GET /renter/file/{siapath}][%d] getRenterFileSiapathOK  %+v", 200, o.Payload)
}
func (o *GetRenterFileSiapathOK) GetPayload() *GetRenterFileSiapathOKBody {
	return o.Payload
}

func (o *GetRenterFileSiapathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRenterFileSiapathOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRenterFileSiapathDefault creates a GetRenterFileSiapathDefault with default headers values
func NewGetRenterFileSiapathDefault(code int) *GetRenterFileSiapathDefault {
	return &GetRenterFileSiapathDefault{
		_statusCode: code,
	}
}

/* GetRenterFileSiapathDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetRenterFileSiapathDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get renter file siapath default response
func (o *GetRenterFileSiapathDefault) Code() int {
	return o._statusCode
}

func (o *GetRenterFileSiapathDefault) Error() string {
	return fmt.Sprintf("[GET /renter/file/{siapath}][%d] GetRenterFileSiapath default  %+v", o._statusCode, o.Payload)
}
func (o *GetRenterFileSiapathDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetRenterFileSiapathDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRenterFileSiapathOKBody get renter file siapath o k body
swagger:model GetRenterFileSiapathOKBody
*/
type GetRenterFileSiapathOKBody struct {

	// file
	File *models.FileInfo `json:"file,omitempty"`
}

// Validate validates this get renter file siapath o k body
func (o *GetRenterFileSiapathOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterFileSiapathOKBody) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(o.File) { // not required
		return nil
	}

	if o.File != nil {
		if err := o.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRenterFileSiapathOK" + "." + "file")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get renter file siapath o k body based on the context it is used
func (o *GetRenterFileSiapathOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterFileSiapathOKBody) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if o.File != nil {
		if err := o.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRenterFileSiapathOK" + "." + "file")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterFileSiapathOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterFileSiapathOKBody) UnmarshalBinary(b []byte) error {
	var res GetRenterFileSiapathOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
