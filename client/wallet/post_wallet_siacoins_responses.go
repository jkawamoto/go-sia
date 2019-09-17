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

// PostWalletSiacoinsReader is a Reader for the PostWalletSiacoins structure.
type PostWalletSiacoinsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWalletSiacoinsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWalletSiacoinsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostWalletSiacoinsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostWalletSiacoinsOK creates a PostWalletSiacoinsOK with default headers values
func NewPostWalletSiacoinsOK() *PostWalletSiacoinsOK {
	return &PostWalletSiacoinsOK{}
}

/*PostWalletSiacoinsOK handles this case with default header values.

Successful Response
*/
type PostWalletSiacoinsOK struct {
	Payload *PostWalletSiacoinsOKBody
}

func (o *PostWalletSiacoinsOK) Error() string {
	return fmt.Sprintf("[POST /wallet/siacoins][%d] postWalletSiacoinsOK  %+v", 200, o.Payload)
}

func (o *PostWalletSiacoinsOK) GetPayload() *PostWalletSiacoinsOKBody {
	return o.Payload
}

func (o *PostWalletSiacoinsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWalletSiacoinsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWalletSiacoinsDefault creates a PostWalletSiacoinsDefault with default headers values
func NewPostWalletSiacoinsDefault(code int) *PostWalletSiacoinsDefault {
	return &PostWalletSiacoinsDefault{
		_statusCode: code,
	}
}

/*PostWalletSiacoinsDefault handles this case with default header values.

Error Response
*/
type PostWalletSiacoinsDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post wallet siacoins default response
func (o *PostWalletSiacoinsDefault) Code() int {
	return o._statusCode
}

func (o *PostWalletSiacoinsDefault) Error() string {
	return fmt.Sprintf("[POST /wallet/siacoins][%d] PostWalletSiacoins default  %+v", o._statusCode, o.Payload)
}

func (o *PostWalletSiacoinsDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostWalletSiacoinsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostWalletSiacoinsOKBody post wallet siacoins o k body
swagger:model PostWalletSiacoinsOKBody
*/
type PostWalletSiacoinsOKBody struct {

	// Array of IDs of the transactions that were created when sending the coins. The last transaction contains the output headed to the 'destination'. Transaction IDs are 64 character long hex strings.
	Transactionids []string `json:"transactionids"`
}

// Validate validates this post wallet siacoins o k body
func (o *PostWalletSiacoinsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWalletSiacoinsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWalletSiacoinsOKBody) UnmarshalBinary(b []byte) error {
	var res PostWalletSiacoinsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
