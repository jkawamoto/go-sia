// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// GetWalletAddressesReader is a Reader for the GetWalletAddresses structure.
type GetWalletAddressesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletAddressesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletAddressesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWalletAddressesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletAddressesOK creates a GetWalletAddressesOK with default headers values
func NewGetWalletAddressesOK() *GetWalletAddressesOK {
	return &GetWalletAddressesOK{}
}

/* GetWalletAddressesOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetWalletAddressesOK struct {
	Payload *GetWalletAddressesOKBody
}

func (o *GetWalletAddressesOK) Error() string {
	return fmt.Sprintf("[GET /wallet/addresses][%d] getWalletAddressesOK  %+v", 200, o.Payload)
}
func (o *GetWalletAddressesOK) GetPayload() *GetWalletAddressesOKBody {
	return o.Payload
}

func (o *GetWalletAddressesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWalletAddressesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletAddressesDefault creates a GetWalletAddressesDefault with default headers values
func NewGetWalletAddressesDefault(code int) *GetWalletAddressesDefault {
	return &GetWalletAddressesDefault{
		_statusCode: code,
	}
}

/* GetWalletAddressesDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetWalletAddressesDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get wallet addresses default response
func (o *GetWalletAddressesDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletAddressesDefault) Error() string {
	return fmt.Sprintf("[GET /wallet/addresses][%d] GetWalletAddresses default  %+v", o._statusCode, o.Payload)
}
func (o *GetWalletAddressesDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetWalletAddressesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWalletAddressesOKBody get wallet addresses o k body
swagger:model GetWalletAddressesOKBody
*/
type GetWalletAddressesOKBody struct {

	// Array of wallet addresses owned by the wallet.
	// Example: {"addresses":["1234567890abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789ab","aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"]}
	Addresses []string `json:"addresses"`
}

// Validate validates this get wallet addresses o k body
func (o *GetWalletAddressesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wallet addresses o k body based on context it is used
func (o *GetWalletAddressesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWalletAddressesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWalletAddressesOKBody) UnmarshalBinary(b []byte) error {
	var res GetWalletAddressesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
