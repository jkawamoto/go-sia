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

// GetWalletVerifyAddressAddrReader is a Reader for the GetWalletVerifyAddressAddr structure.
type GetWalletVerifyAddressAddrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletVerifyAddressAddrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletVerifyAddressAddrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWalletVerifyAddressAddrDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletVerifyAddressAddrOK creates a GetWalletVerifyAddressAddrOK with default headers values
func NewGetWalletVerifyAddressAddrOK() *GetWalletVerifyAddressAddrOK {
	return &GetWalletVerifyAddressAddrOK{}
}

/*GetWalletVerifyAddressAddrOK handles this case with default header values.

Successful Response
*/
type GetWalletVerifyAddressAddrOK struct {
	Payload *GetWalletVerifyAddressAddrOKBody
}

func (o *GetWalletVerifyAddressAddrOK) Error() string {
	return fmt.Sprintf("[GET /wallet/verify/address/{addr}][%d] getWalletVerifyAddressAddrOK  %+v", 200, o.Payload)
}

func (o *GetWalletVerifyAddressAddrOK) GetPayload() *GetWalletVerifyAddressAddrOKBody {
	return o.Payload
}

func (o *GetWalletVerifyAddressAddrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWalletVerifyAddressAddrOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletVerifyAddressAddrDefault creates a GetWalletVerifyAddressAddrDefault with default headers values
func NewGetWalletVerifyAddressAddrDefault(code int) *GetWalletVerifyAddressAddrDefault {
	return &GetWalletVerifyAddressAddrDefault{
		_statusCode: code,
	}
}

/*GetWalletVerifyAddressAddrDefault handles this case with default header values.

Error Response
*/
type GetWalletVerifyAddressAddrDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get wallet verify address addr default response
func (o *GetWalletVerifyAddressAddrDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletVerifyAddressAddrDefault) Error() string {
	return fmt.Sprintf("[GET /wallet/verify/address/{addr}][%d] GetWalletVerifyAddressAddr default  %+v", o._statusCode, o.Payload)
}

func (o *GetWalletVerifyAddressAddrDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetWalletVerifyAddressAddrDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWalletVerifyAddressAddrOKBody get wallet verify address addr o k body
swagger:model GetWalletVerifyAddressAddrOKBody
*/
type GetWalletVerifyAddressAddrOKBody struct {

	// valid indicates if the address supplied to addr is a valid UnlockHash.
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this get wallet verify address addr o k body
func (o *GetWalletVerifyAddressAddrOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWalletVerifyAddressAddrOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWalletVerifyAddressAddrOKBody) UnmarshalBinary(b []byte) error {
	var res GetWalletVerifyAddressAddrOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
