// Code generated by go-swagger; DO NOT EDIT.

package renter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/jkawamoto/go-sia/models"
)

// GetRenterDirSiapathReader is a Reader for the GetRenterDirSiapath structure.
type GetRenterDirSiapathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRenterDirSiapathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRenterDirSiapathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRenterDirSiapathDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRenterDirSiapathOK creates a GetRenterDirSiapathOK with default headers values
func NewGetRenterDirSiapathOK() *GetRenterDirSiapathOK {
	return &GetRenterDirSiapathOK{}
}

/* GetRenterDirSiapathOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetRenterDirSiapathOK struct {
	Payload *GetRenterDirSiapathOKBody
}

func (o *GetRenterDirSiapathOK) Error() string {
	return fmt.Sprintf("[GET /renter/dir/{siapath}][%d] getRenterDirSiapathOK  %+v", 200, o.Payload)
}
func (o *GetRenterDirSiapathOK) GetPayload() *GetRenterDirSiapathOKBody {
	return o.Payload
}

func (o *GetRenterDirSiapathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRenterDirSiapathOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRenterDirSiapathDefault creates a GetRenterDirSiapathDefault with default headers values
func NewGetRenterDirSiapathDefault(code int) *GetRenterDirSiapathDefault {
	return &GetRenterDirSiapathDefault{
		_statusCode: code,
	}
}

/* GetRenterDirSiapathDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetRenterDirSiapathDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get renter dir siapath default response
func (o *GetRenterDirSiapathDefault) Code() int {
	return o._statusCode
}

func (o *GetRenterDirSiapathDefault) Error() string {
	return fmt.Sprintf("[GET /renter/dir/{siapath}][%d] GetRenterDirSiapath default  %+v", o._statusCode, o.Payload)
}
func (o *GetRenterDirSiapathDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetRenterDirSiapathDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRenterDirSiapathOKBody get renter dir siapath o k body
swagger:model GetRenterDirSiapathOKBody
*/
type GetRenterDirSiapathOKBody struct {

	// directories
	Directories []*GetRenterDirSiapathOKBodyDirectoriesItems0 `json:"directories"`

	// files
	Files []*models.FileInfo `json:"files"`
}

// Validate validates this get renter dir siapath o k body
func (o *GetRenterDirSiapathOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDirectories(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterDirSiapathOKBody) validateDirectories(formats strfmt.Registry) error {
	if swag.IsZero(o.Directories) { // not required
		return nil
	}

	for i := 0; i < len(o.Directories); i++ {
		if swag.IsZero(o.Directories[i]) { // not required
			continue
		}

		if o.Directories[i] != nil {
			if err := o.Directories[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRenterDirSiapathOK" + "." + "directories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetRenterDirSiapathOKBody) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(o.Files) { // not required
		return nil
	}

	for i := 0; i < len(o.Files); i++ {
		if swag.IsZero(o.Files[i]) { // not required
			continue
		}

		if o.Files[i] != nil {
			if err := o.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRenterDirSiapathOK" + "." + "files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get renter dir siapath o k body based on the context it is used
func (o *GetRenterDirSiapathOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDirectories(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterDirSiapathOKBody) contextValidateDirectories(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Directories); i++ {

		if o.Directories[i] != nil {
			if err := o.Directories[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRenterDirSiapathOK" + "." + "directories" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetRenterDirSiapathOKBody) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Files); i++ {

		if o.Files[i] != nil {
			if err := o.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRenterDirSiapathOK" + "." + "files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterDirSiapathOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterDirSiapathOKBody) UnmarshalBinary(b []byte) error {
	var res GetRenterDirSiapathOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRenterDirSiapathOKBodyDirectoriesItems0 get renter dir siapath o k body directories items0
swagger:model GetRenterDirSiapathOKBodyDirectoriesItems0
*/
type GetRenterDirSiapathOKBodyDirectoriesItems0 struct {

	// the total number of files in the sub directory tree
	// Example: 2
	Aggregatenumfiles int64 `json:"aggregatenumfiles,omitempty"`

	// aggregatenumsize
	// Example: 4096
	Aggregatenumsize int64 `json:"aggregatenumsize,omitempty"`

	// the total number of stuck chunks in the sub directory tree
	// Example: 4
	Aggregatenumstuckchunks int64 `json:"aggregatenumstuckchunks,omitempty"`

	// This is the worst health of any of the files or subdirectories. Health is the percent of parity pieces missing. - health = 0 is full redundancy - health <= 1 is recoverable - health > 1 needs to be repaired from disk
	// Example: 1
	Health float64 `json:"health,omitempty"`

	// The oldest time that the health of the directory or any of its files or sub directories' health was checked.
	// Example: 2018-09-23T08:00:00.000000000+04:00
	// Format: date-time
	Lasthealthchecktime strfmt.DateTime `json:"lasthealthchecktime,omitempty"`

	// This is the worst health when comparing stuck health vs health
	// Example: 0.5
	Maxhealth float64 `json:"maxhealth,omitempty"`

	// the lowest redundancy of any file or directory in the sub directory tree
	// Example: 2.6
	Minredundancy float64 `json:"minredundancy,omitempty"`

	// the most recent mod time of any file or directory in the sub directory tree
	// Example: 2018-09-23T08:00:00.000000000+04:00
	// Format: date-time
	Mostrecentmodtime strfmt.DateTime `json:"mostrecentmodtime,omitempty"`

	// the number of files in the directory
	// Example: 3
	Numfiles int64 `json:"numfiles,omitempty"`

	// the number of directories in the directory
	// Example: 2
	Numsubdirs int64 `json:"numsubdirs,omitempty"`

	// The path to the directory on the sia network
	// Example: foo/bar
	Siapath string `json:"siapath,omitempty"`
}

// Validate validates this get renter dir siapath o k body directories items0
func (o *GetRenterDirSiapathOKBodyDirectoriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLasthealthchecktime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMostrecentmodtime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterDirSiapathOKBodyDirectoriesItems0) validateLasthealthchecktime(formats strfmt.Registry) error {
	if swag.IsZero(o.Lasthealthchecktime) { // not required
		return nil
	}

	if err := validate.FormatOf("lasthealthchecktime", "body", "date-time", o.Lasthealthchecktime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetRenterDirSiapathOKBodyDirectoriesItems0) validateMostrecentmodtime(formats strfmt.Registry) error {
	if swag.IsZero(o.Mostrecentmodtime) { // not required
		return nil
	}

	if err := validate.FormatOf("mostrecentmodtime", "body", "date-time", o.Mostrecentmodtime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get renter dir siapath o k body directories items0 based on context it is used
func (o *GetRenterDirSiapathOKBodyDirectoriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterDirSiapathOKBodyDirectoriesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterDirSiapathOKBodyDirectoriesItems0) UnmarshalBinary(b []byte) error {
	var res GetRenterDirSiapathOKBodyDirectoriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
