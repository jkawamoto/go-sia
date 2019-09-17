// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// GetWalletTransactionsReader is a Reader for the GetWalletTransactions structure.
type GetWalletTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWalletTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWalletTransactionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletTransactionsOK creates a GetWalletTransactionsOK with default headers values
func NewGetWalletTransactionsOK() *GetWalletTransactionsOK {
	return &GetWalletTransactionsOK{}
}

/*GetWalletTransactionsOK handles this case with default header values.

Successful Response
*/
type GetWalletTransactionsOK struct {
	Payload *GetWalletTransactionsOKBody
}

func (o *GetWalletTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /wallet/transactions][%d] getWalletTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetWalletTransactionsOK) GetPayload() *GetWalletTransactionsOKBody {
	return o.Payload
}

func (o *GetWalletTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWalletTransactionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWalletTransactionsDefault creates a GetWalletTransactionsDefault with default headers values
func NewGetWalletTransactionsDefault(code int) *GetWalletTransactionsDefault {
	return &GetWalletTransactionsDefault{
		_statusCode: code,
	}
}

/*GetWalletTransactionsDefault handles this case with default header values.

Error Response
*/
type GetWalletTransactionsDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get wallet transactions default response
func (o *GetWalletTransactionsDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletTransactionsDefault) Error() string {
	return fmt.Sprintf("[GET /wallet/transactions][%d] GetWalletTransactions default  %+v", o._statusCode, o.Payload)
}

func (o *GetWalletTransactionsDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetWalletTransactionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWalletTransactionsOKBody get wallet transactions o k body
swagger:model GetWalletTransactionsOKBody
*/
type GetWalletTransactionsOKBody struct {

	// All of the confirmed transactions appearing between height 'startheight' and height 'endheight' (inclusive).
	Confirmedtransactions []*models.Transaction `json:"confirmedtransactions"`

	// All of the unconfirmed transactions.
	Unconfirmedtransactions []*models.Transaction `json:"unconfirmedtransactions"`
}

// Validate validates this get wallet transactions o k body
func (o *GetWalletTransactionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConfirmedtransactions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUnconfirmedtransactions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWalletTransactionsOKBody) validateConfirmedtransactions(formats strfmt.Registry) error {

	if swag.IsZero(o.Confirmedtransactions) { // not required
		return nil
	}

	for i := 0; i < len(o.Confirmedtransactions); i++ {
		if swag.IsZero(o.Confirmedtransactions[i]) { // not required
			continue
		}

		if o.Confirmedtransactions[i] != nil {
			if err := o.Confirmedtransactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWalletTransactionsOK" + "." + "confirmedtransactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetWalletTransactionsOKBody) validateUnconfirmedtransactions(formats strfmt.Registry) error {

	if swag.IsZero(o.Unconfirmedtransactions) { // not required
		return nil
	}

	for i := 0; i < len(o.Unconfirmedtransactions); i++ {
		if swag.IsZero(o.Unconfirmedtransactions[i]) { // not required
			continue
		}

		if o.Unconfirmedtransactions[i] != nil {
			if err := o.Unconfirmedtransactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWalletTransactionsOK" + "." + "unconfirmedtransactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWalletTransactionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWalletTransactionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetWalletTransactionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
