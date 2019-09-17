// Code generated by go-swagger; DO NOT EDIT.

package renter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jkawamoto/go-sia/models"
)

// GetRenterReader is a Reader for the GetRenter structure.
type GetRenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRenterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRenterOK creates a GetRenterOK with default headers values
func NewGetRenterOK() *GetRenterOK {
	return &GetRenterOK{}
}

/*GetRenterOK handles this case with default header values.

Successful Response
*/
type GetRenterOK struct {
	Payload *GetRenterOKBody
}

func (o *GetRenterOK) Error() string {
	return fmt.Sprintf("[GET /renter][%d] getRenterOK  %+v", 200, o.Payload)
}

func (o *GetRenterOK) GetPayload() *GetRenterOKBody {
	return o.Payload
}

func (o *GetRenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRenterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRenterDefault creates a GetRenterDefault with default headers values
func NewGetRenterDefault(code int) *GetRenterDefault {
	return &GetRenterDefault{
		_statusCode: code,
	}
}

/*GetRenterDefault handles this case with default header values.

Error Response
*/
type GetRenterDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get renter default response
func (o *GetRenterDefault) Code() int {
	return o._statusCode
}

func (o *GetRenterDefault) Error() string {
	return fmt.Sprintf("[GET /renter][%d] GetRenter default  %+v", o._statusCode, o.Payload)
}

func (o *GetRenterDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetRenterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetRenterOKBody get renter o k body
swagger:model GetRenterOKBody
*/
type GetRenterOKBody struct {

	// Height at which the current allowance period began.
	Currentperiod int64 `json:"currentperiod,omitempty"`

	// financialmetrics
	Financialmetrics *GetRenterOKBodyFinancialmetrics `json:"financialmetrics,omitempty"`

	// settings
	Settings *GetRenterOKBodySettings `json:"settings,omitempty"`
}

// Validate validates this get renter o k body
func (o *GetRenterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFinancialmetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterOKBody) validateFinancialmetrics(formats strfmt.Registry) error {

	if swag.IsZero(o.Financialmetrics) { // not required
		return nil
	}

	if o.Financialmetrics != nil {
		if err := o.Financialmetrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRenterOK" + "." + "financialmetrics")
			}
			return err
		}
	}

	return nil
}

func (o *GetRenterOKBody) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(o.Settings) { // not required
		return nil
	}

	if o.Settings != nil {
		if err := o.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRenterOK" + "." + "settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterOKBody) UnmarshalBinary(b []byte) error {
	var res GetRenterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRenterOKBodyFinancialmetrics Metrics about how much the Renter has spent on storage, uploads, and downloads.
swagger:model GetRenterOKBodyFinancialmetrics
*/
type GetRenterOKBodyFinancialmetrics struct {

	// Amount of money spent on contract fees, transaction fees and siafund fees. In hastings.
	Contractfees string `json:"contractfees,omitempty"`

	// How much money, in hastings, the Renter has spent on file contracts, including fees. (deprecated, now totalallocated)
	Contractspending string `json:"contractspending,omitempty"`

	// Amount of money spent on downloads.
	Downloadspending string `json:"downloadspending,omitempty"`

	// Amount of money spend on storage.
	Storagespending string `json:"storagespending,omitempty"`

	// Total amount of money that the renter has put into contracts. Includes spent money and also money that will be returned to the renter. In hastings.
	Totalallocated string `json:"totalallocated,omitempty"`

	// Amount of money in the allowance that has not been spent.
	Unspent string `json:"unspent,omitempty"`

	// Amount of money spent on uploads.
	Uploadspending string `json:"uploadspending,omitempty"`
}

// Validate validates this get renter o k body financialmetrics
func (o *GetRenterOKBodyFinancialmetrics) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterOKBodyFinancialmetrics) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterOKBodyFinancialmetrics) UnmarshalBinary(b []byte) error {
	var res GetRenterOKBodyFinancialmetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRenterOKBodySettings Settings that control the behavior of the renter.
swagger:model GetRenterOKBodySettings
*/
type GetRenterOKBodySettings struct {

	// allowance
	Allowance *GetRenterOKBodySettingsAllowance `json:"allowance,omitempty"`

	// bytes per second. MaxDownloadSpeed by defaul is unlimited but can be set by the user to manage bandwidth
	Maxdownloadspeed int64 `json:"maxdownloadspeed,omitempty"`

	// bytes per second. MaxUploadSpeed by defaul is unlimited but can be set by the user to manage bandwidth
	Maxuploadspeed int64 `json:"maxuploadspeed,omitempty"`

	// The StreamCacheSize is the number of data chunks that will be cached during streaming
	Streamcachesize int64 `json:"streamcachesize,omitempty"`
}

// Validate validates this get renter o k body settings
func (o *GetRenterOKBodySettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllowance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRenterOKBodySettings) validateAllowance(formats strfmt.Registry) error {

	if swag.IsZero(o.Allowance) { // not required
		return nil
	}

	if o.Allowance != nil {
		if err := o.Allowance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getRenterOK" + "." + "settings" + "." + "allowance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterOKBodySettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterOKBodySettings) UnmarshalBinary(b []byte) error {
	var res GetRenterOKBodySettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetRenterOKBodySettingsAllowance Allowance dictates how much the renter is allowed to spend in a given period. Note that funds are spent on both storage and bandwidth.
swagger:model GetRenterOKBodySettingsAllowance
*/
type GetRenterOKBodySettingsAllowance struct {

	// Amount of money allocated for contracts. Funds are spent on both storage and bandwidth.
	Funds string `json:"funds,omitempty"`

	// Number of hosts that contracts will be formed with.
	Hosts int64 `json:"hosts,omitempty"`

	// Duration of contracts formed, in number of blocks.
	Period int64 `json:"period,omitempty"`

	// If the current blockheight + the renew window >= the height the contract is scheduled to end, the contract is renewed automatically. Is always nonzero.
	Renewwindow int64 `json:"renewwindow,omitempty"`
}

// Validate validates this get renter o k body settings allowance
func (o *GetRenterOKBodySettingsAllowance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRenterOKBodySettingsAllowance) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRenterOKBodySettingsAllowance) UnmarshalBinary(b []byte) error {
	var res GetRenterOKBodySettingsAllowance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
