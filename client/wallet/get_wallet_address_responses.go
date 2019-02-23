// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// GetWalletAddressReader is a Reader for the GetWalletAddress structure.
type GetWalletAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWalletAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWalletAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletAddressOK creates a GetWalletAddressOK with default headers values
func NewGetWalletAddressOK() *GetWalletAddressOK {
	return &GetWalletAddressOK{}
}

/*GetWalletAddressOK handles this case with default header values.

Successful Response
*/
type GetWalletAddressOK struct {
	Payload *GetWalletAddressOKBody
}

func (o *GetWalletAddressOK) Error() string {
	return fmt.Sprintf("[GET /wallet/address][%d] getWalletAddressOK  %+v", 200, o.Payload)
}

func (o *GetWalletAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWalletAddressOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletAddressDefault creates a GetWalletAddressDefault with default headers values
func NewGetWalletAddressDefault(code int) *GetWalletAddressDefault {
	return &GetWalletAddressDefault{
		_statusCode: code,
	}
}

/*GetWalletAddressDefault handles this case with default header values.

Error Response
*/
type GetWalletAddressDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get wallet address default response
func (o *GetWalletAddressDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletAddressDefault) Error() string {
	return fmt.Sprintf("[GET /wallet/address][%d] GetWalletAddress default  %+v", o._statusCode, o.Payload)
}

func (o *GetWalletAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWalletAddressOKBody get wallet address o k body
swagger:model GetWalletAddressOKBody
*/
type GetWalletAddressOKBody struct {

	// Wallet address that can receive siacoins or siafunds. Addresses are 76 character long hex strings.
	Address string `json:"address,omitempty"`
}

// Validate validates this get wallet address o k body
func (o *GetWalletAddressOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWalletAddressOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWalletAddressOKBody) UnmarshalBinary(b []byte) error {
	var res GetWalletAddressOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
