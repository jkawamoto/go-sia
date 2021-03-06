// Code generated by go-swagger; DO NOT EDIT.

package consensus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/models"
)

// PostConsensusValidateTransactionsetReader is a Reader for the PostConsensusValidateTransactionset structure.
type PostConsensusValidateTransactionsetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConsensusValidateTransactionsetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostConsensusValidateTransactionsetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostConsensusValidateTransactionsetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostConsensusValidateTransactionsetNoContent creates a PostConsensusValidateTransactionsetNoContent with default headers values
func NewPostConsensusValidateTransactionsetNoContent() *PostConsensusValidateTransactionsetNoContent {
	return &PostConsensusValidateTransactionsetNoContent{}
}

/* PostConsensusValidateTransactionsetNoContent describes a response with status code 204, with default header values.

Successful Response
*/
type PostConsensusValidateTransactionsetNoContent struct {
}

func (o *PostConsensusValidateTransactionsetNoContent) Error() string {
	return fmt.Sprintf("[POST /consensus/validate/transactionset][%d] postConsensusValidateTransactionsetNoContent ", 204)
}

func (o *PostConsensusValidateTransactionsetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConsensusValidateTransactionsetDefault creates a PostConsensusValidateTransactionsetDefault with default headers values
func NewPostConsensusValidateTransactionsetDefault(code int) *PostConsensusValidateTransactionsetDefault {
	return &PostConsensusValidateTransactionsetDefault{
		_statusCode: code,
	}
}

/* PostConsensusValidateTransactionsetDefault describes a response with status code -1, with default header values.

Error Response
*/
type PostConsensusValidateTransactionsetDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the post consensus validate transactionset default response
func (o *PostConsensusValidateTransactionsetDefault) Code() int {
	return o._statusCode
}

func (o *PostConsensusValidateTransactionsetDefault) Error() string {
	return fmt.Sprintf("[POST /consensus/validate/transactionset][%d] PostConsensusValidateTransactionset default  %+v", o._statusCode, o.Payload)
}
func (o *PostConsensusValidateTransactionsetDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *PostConsensusValidateTransactionsetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
