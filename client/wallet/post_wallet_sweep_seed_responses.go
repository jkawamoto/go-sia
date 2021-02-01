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

// PostWalletSweepSeedReader is a Reader for the PostWalletSweepSeed structure.
type PostWalletSweepSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWalletSweepSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWalletSweepSeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWalletSweepSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWalletSweepSeedOK creates a PostWalletSweepSeedOK with default headers values
func NewPostWalletSweepSeedOK() *PostWalletSweepSeedOK {
	return &PostWalletSweepSeedOK{}
}

/* PostWalletSweepSeedOK describes a response with status code 200, with default header values.

Successful Response
*/
type PostWalletSweepSeedOK struct {
	Payload *PostWalletSweepSeedOKBody
}

func (o *PostWalletSweepSeedOK) Error() string {
	return fmt.Sprintf("[POST /wallet/sweep/seed][%d] postWalletSweepSeedOK  %+v", 200, o.Payload)
}
func (o *PostWalletSweepSeedOK) GetPayload() *PostWalletSweepSeedOKBody {
	return o.Payload
}

func (o *PostWalletSweepSeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWalletSweepSeedOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWalletSweepSeedDefault creates a PostWalletSweepSeedDefault with default headers values
func NewPostWalletSweepSeedDefault(code int) *PostWalletSweepSeedDefault {
	return &PostWalletSweepSeedDefault{
		_statusCode: code,
	}
}

/* PostWalletSweepSeedDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostWalletSweepSeedDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post wallet sweep seed default response
func (o *PostWalletSweepSeedDefault) Code() int {
	return o._statusCode
}

func (o *PostWalletSweepSeedDefault) Error() string {
	return fmt.Sprintf("[POST /wallet/sweep/seed][%d] PostWalletSweepSeed default  %+v", o._statusCode, o.Payload)
}
func (o *PostWalletSweepSeedDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostWalletSweepSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostWalletSweepSeedOKBody post wallet sweep seed o k body
swagger:model PostWalletSweepSeedOKBody
*/
type PostWalletSweepSeedOKBody struct {

	// Number of siacoins, in hastings, transferred to the wallet as a result of the sweep.
	// Example: 123456
	Coins string `json:"coins,omitempty"`

	// Number of siafunds transferred to the wallet as a result of the sweep.
	// Example: 1
	Funds string `json:"funds,omitempty"`
}

// Validate validates this post wallet sweep seed o k body
func (o *PostWalletSweepSeedOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wallet sweep seed o k body based on context it is used
func (o *PostWalletSweepSeedOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWalletSweepSeedOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWalletSweepSeedOKBody) UnmarshalBinary(b []byte) error {
	var res PostWalletSweepSeedOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
