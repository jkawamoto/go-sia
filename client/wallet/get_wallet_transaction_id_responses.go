// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// GetWalletTransactionIDReader is a Reader for the GetWalletTransactionID structure.
type GetWalletTransactionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletTransactionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletTransactionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWalletTransactionIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletTransactionIDOK creates a GetWalletTransactionIDOK with default headers values
func NewGetWalletTransactionIDOK() *GetWalletTransactionIDOK {
	return &GetWalletTransactionIDOK{}
}

/* GetWalletTransactionIDOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetWalletTransactionIDOK struct {
	Payload *GetWalletTransactionIDOKBody
}

func (o *GetWalletTransactionIDOK) Error() string {
	return fmt.Sprintf("[GET /wallet/transaction/{id}][%d] getWalletTransactionIdOK  %+v", 200, o.Payload)
}
func (o *GetWalletTransactionIDOK) GetPayload() *GetWalletTransactionIDOKBody {
	return o.Payload
}

func (o *GetWalletTransactionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWalletTransactionIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletTransactionIDDefault creates a GetWalletTransactionIDDefault with default headers values
func NewGetWalletTransactionIDDefault(code int) *GetWalletTransactionIDDefault {
	return &GetWalletTransactionIDDefault{
		_statusCode: code,
	}
}

/* GetWalletTransactionIDDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetWalletTransactionIDDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get wallet transaction ID default response
func (o *GetWalletTransactionIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletTransactionIDDefault) Error() string {
	return fmt.Sprintf("[GET /wallet/transaction/{id}][%d] GetWalletTransactionID default  %+v", o._statusCode, o.Payload)
}
func (o *GetWalletTransactionIDDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetWalletTransactionIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWalletTransactionIDOKBody get wallet transaction ID o k body
swagger:model GetWalletTransactionIDOKBody
*/
type GetWalletTransactionIDOKBody struct {

	// transaction
	Transaction *models.Transaction `json:"transaction,omitempty"`
}

// Validate validates this get wallet transaction ID o k body
func (o *GetWalletTransactionIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTransaction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWalletTransactionIDOKBody) validateTransaction(formats strfmt.Registry) error {
	if swag.IsZero(o.Transaction) { // not required
		return nil
	}

	if o.Transaction != nil {
		if err := o.Transaction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWalletTransactionIdOK" + "." + "transaction")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get wallet transaction ID o k body based on the context it is used
func (o *GetWalletTransactionIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTransaction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWalletTransactionIDOKBody) contextValidateTransaction(ctx context.Context, formats strfmt.Registry) error {

	if o.Transaction != nil {
		if err := o.Transaction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getWalletTransactionIdOK" + "." + "transaction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWalletTransactionIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWalletTransactionIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetWalletTransactionIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
