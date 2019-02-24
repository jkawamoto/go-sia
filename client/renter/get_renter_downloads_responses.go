// Code generated by go-swagger; DO NOT EDIT.

package renter

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

// GetRenterDownloadsReader is a Reader for the GetRenterDownloads structure.
type GetRenterDownloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRenterDownloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRenterDownloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetRenterDownloadsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRenterDownloadsOK creates a GetRenterDownloadsOK with default headers values
func NewGetRenterDownloadsOK() *GetRenterDownloadsOK {
	return &GetRenterDownloadsOK{}
}

/*GetRenterDownloadsOK handles this case with default header values.

Successful Response
*/
type GetRenterDownloadsOK struct {
	Payload *GetRenterDownloadsOKBody
}

func (o *GetRenterDownloadsOK) Error() string {
	return fmt.Sprintf("[GET /renter/downloads][%d] getRenterDownloadsOK  %+v", 200, o.Payload)
}

func (o *GetRenterDownloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRenterDownloadsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRenterDownloadsDefault creates a GetRenterDownloadsDefault with default headers values
func NewGetRenterDownloadsDefault(code int) *GetRenterDownloadsDefault {
	return &GetRenterDownloadsDefault{
		_statusCode: code,
	}
}

/*GetRenterDownloadsDefault handles this case with default header values.

Error Response
*/
type GetRenterDownloadsDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get renter downloads default response
func (o *GetRenterDownloadsDefault) Code() int {
	return o._statusCode
}

func (o *GetRenterDownloadsDefault) Error() string {
	return fmt.Sprintf("[GET /renter/downloads][%d] GetRenterDownloads default  %+v", o._statusCode, o.Payload)
}

func (o *GetRenterDownloadsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DownloadsItems0 downloads items0
swagger:model DownloadsItems0
*/
type DownloadsItems0 struct {

	// Whether or not the download has completed. Will be false initially, and set to true immediately as the download has been fully written out to the file, to the http stream, or to the in-memory buffer. Completed will also be set to true if there is an error that causes the download to fail.
	Completed bool `json:"completed,omitempty"`

	// Local path that the file will be downloaded to.
	Destination string `json:"destination,omitempty"`

	// What type of destination was used. Can be "file", indicating a download to disk, can be "buffer", indicating a download to memory, and can be "http stream", indicating that the download was streamed through the http API.
	Destinationtype string `json:"destinationtype,omitempty"`

	// Time at which the download completed in RFC 3339 time. Will be zero if the download has not yet completed.
	Endtime string `json:"endtime,omitempty"`

	// Error encountered while downloading. If there was no error (yet), it will be the empty string.
	Error string `json:"error,omitempty"`

	// Length of the download. If the download was a partial download, this will indicate the length of the partial download, and not the length of the full file.
	Length int64 `json:"length,omitempty"`

	// Offset within the file of the download. For full file downloads, the offset will be '0'. For partial downloads, the offset may be anywhere within the file. offset+length will never exceed the full file size.
	Offset int64 `json:"offset,omitempty"`

	// Number of bytes downloaded thus far. Will only be updated as segments of the file complete fully. This typically has a resolution of tens of megabytes.
	Received int64 `json:"received,omitempty"`

	// Siapath given to the file when it was uploaded.
	Siapath string `json:"siapath,omitempty"`

	// Time at which the download was initiated in RFC 3339.
	Starttime string `json:"starttime,omitempty"`

	// The total amount of data transfered when downloading the file. This will eventually include data transferred during contract + payment negotiation, as well as data from failed piece downloads.
	Totaldatatransfered int64 `json:"totaldatatransfered,omitempty"`
}

// Validate validates this downloads items0
func (o *DownloadsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DownloadsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DownloadsItems0) UnmarshalBinary(b []byte) error {
	var res DownloadsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRenterDownloadsOKBody get renter downloads o k body
swagger:model GetRenterDownloadsOKBody
*/
type GetRenterDownloadsOKBody struct {

	// downloads
	Downloads []*DownloadsItems0 `json:"downloads"`
}

// Validate validates this get renter downloads o k body
func (o *GetRenterDownloadsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDownloads(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterDownloadsOKBody) validateDownloads(formats strfmt.Registry) error {

	if swag.IsZero(o.Downloads) { // not required
		return nil
	}

	for i := 0; i < len(o.Downloads); i++ {
		if swag.IsZero(o.Downloads[i]) { // not required
			continue
		}

		if o.Downloads[i] != nil {
			if err := o.Downloads[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRenterDownloadsOK" + "." + "downloads" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterDownloadsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterDownloadsOKBody) UnmarshalBinary(b []byte) error {
	var res GetRenterDownloadsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}