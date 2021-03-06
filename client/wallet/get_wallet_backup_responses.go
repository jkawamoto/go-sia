// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/models"
)

// GetWalletBackupReader is a Reader for the GetWalletBackup structure.
type GetWalletBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWalletBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetWalletBackupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWalletBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWalletBackupNoContent creates a GetWalletBackupNoContent with default headers values
func NewGetWalletBackupNoContent() *GetWalletBackupNoContent {
	return &GetWalletBackupNoContent{}
}

/* GetWalletBackupNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type GetWalletBackupNoContent struct {
}

func (o *GetWalletBackupNoContent) Error() string {
	return fmt.Sprintf("[GET /wallet/backup][%d] getWalletBackupNoContent ", 204)
}

func (o *GetWalletBackupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWalletBackupDefault creates a GetWalletBackupDefault with default headers values
func NewGetWalletBackupDefault(code int) *GetWalletBackupDefault {
	return &GetWalletBackupDefault{
		_statusCode: code,
	}
}

/* GetWalletBackupDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetWalletBackupDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get wallet backup default response
func (o *GetWalletBackupDefault) Code() int {
	return o._statusCode
}

func (o *GetWalletBackupDefault) Error() string {
	return fmt.Sprintf("[GET /wallet/backup][%d] GetWalletBackup default  %+v", o._statusCode, o.Payload)
}
func (o *GetWalletBackupDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetWalletBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
