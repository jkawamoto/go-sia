// Code generated by go-swagger; DO NOT EDIT.

package host_d_b

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

// GetHostdbHostsPubkeyReader is a Reader for the GetHostdbHostsPubkey structure.
type GetHostdbHostsPubkeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostdbHostsPubkeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostdbHostsPubkeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetHostdbHostsPubkeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetHostdbHostsPubkeyOK creates a GetHostdbHostsPubkeyOK with default headers values
func NewGetHostdbHostsPubkeyOK() *GetHostdbHostsPubkeyOK {
	return &GetHostdbHostsPubkeyOK{}
}

/* GetHostdbHostsPubkeyOK describes a response with status code 200, with default header values.

Successful Response
*/
type GetHostdbHostsPubkeyOK struct {
	Payload *GetHostdbHostsPubkeyOKBody
}

func (o *GetHostdbHostsPubkeyOK) Error() string {
	return fmt.Sprintf("[GET /hostdb/hosts/{pubkey}][%d] getHostdbHostsPubkeyOK  %+v", 200, o.Payload)
}
func (o *GetHostdbHostsPubkeyOK) GetPayload() *GetHostdbHostsPubkeyOKBody {
	return o.Payload
}

func (o *GetHostdbHostsPubkeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetHostdbHostsPubkeyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostdbHostsPubkeyDefault creates a GetHostdbHostsPubkeyDefault with default headers values
func NewGetHostdbHostsPubkeyDefault(code int) *GetHostdbHostsPubkeyDefault {
	return &GetHostdbHostsPubkeyDefault{
		_statusCode: code,
	}
}

/* GetHostdbHostsPubkeyDefault describes a response with status code -1, with default header values.

Error Response
*/
type GetHostdbHostsPubkeyDefault struct {
	_statusCode int

	Payload *models.StandardError
}

// Code gets the status code for the get hostdb hosts pubkey default response
func (o *GetHostdbHostsPubkeyDefault) Code() int {
	return o._statusCode
}

func (o *GetHostdbHostsPubkeyDefault) Error() string {
	return fmt.Sprintf("[GET /hostdb/hosts/{pubkey}][%d] GetHostdbHostsPubkey default  %+v", o._statusCode, o.Payload)
}
func (o *GetHostdbHostsPubkeyDefault) GetPayload() *models.StandardError {
	return o.Payload
}

func (o *GetHostdbHostsPubkeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StandardError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetHostdbHostsPubkeyOKBody get hostdb hosts pubkey o k body
swagger:model GetHostdbHostsPubkeyOKBody
*/
type GetHostdbHostsPubkeyOKBody struct {

	// entry
	Entry *GetHostdbHostsPubkeyOKBodyEntry `json:"entry,omitempty"`
}

// Validate validates this get hostdb hosts pubkey o k body
func (o *GetHostdbHostsPubkeyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEntry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHostdbHostsPubkeyOKBody) validateEntry(formats strfmt.Registry) error {
	if swag.IsZero(o.Entry) { // not required
		return nil
	}

	if o.Entry != nil {
		if err := o.Entry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHostdbHostsPubkeyOK" + "." + "entry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get hostdb hosts pubkey o k body based on the context it is used
func (o *GetHostdbHostsPubkeyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEntry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHostdbHostsPubkeyOKBody) contextValidateEntry(ctx context.Context, formats strfmt.Registry) error {

	if o.Entry != nil {
		if err := o.Entry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHostdbHostsPubkeyOK" + "." + "entry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBody) UnmarshalBinary(b []byte) error {
	var res GetHostdbHostsPubkeyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetHostdbHostsPubkeyOKBodyEntry get hostdb hosts pubkey o k body entry
swagger:model GetHostdbHostsPubkeyOKBodyEntry
*/
type GetHostdbHostsPubkeyOKBodyEntry struct {

	// true if the host is accepting new contracts.
	// Example: true
	Acceptingcontracts bool `json:"acceptingcontracts,omitempty"`

	// Maximum number of bytes that the host will allow to be requested by a single download request.
	// Example: 17825792
	Maxdownloadbatchsize int64 `json:"maxdownloadbatchsize,omitempty"`

	// Maximum duration in blocks that a host will allow for a file contract.
	// The host commits to keeping files for the full duration under the
	// threat of facing a large penalty for losing or dropping data before
	// the duration is complete. The storage proof window of an incoming file
	// contract must end before the current height + maxduration.
	// There is a block approximately every 10 minutes.
	// e.g. 1 day = 144 blocks
	//
	// Example: 25920
	Maxduration int64 `json:"maxduration,omitempty"`

	// Maximum size in bytes of a single batch of file contract
	// revisions. Larger batch sizes allow for higher throughput as there is
	// significant communication overhead associated with performing a batch
	// upload.
	//
	// Example: 17825792
	Maxrevisebatchsize int64 `json:"maxrevisebatchsize,omitempty"`

	// Remote address of the host. It can be an IPv4, IPv6, or hostname, along with the port. IPv6 addresses are enclosed in square brackets.
	// Example: 123.456.789.0:9982
	Netaddress string `json:"netaddress,omitempty"`

	// publickey
	Publickey *GetHostdbHostsPubkeyOKBodyEntryPublickey `json:"publickey,omitempty"`

	// The string representation of the full public key, used when calling /hostdb/hosts.
	// Example: ed25519:1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	Publickeystring string `json:"publickeystring,omitempty"`

	// Unused storage capacity the host claims it has, in bytes.
	// Example: 35000000000
	Remainingstorage int64 `json:"remainingstorage,omitempty"`

	// scorebreakdown
	Scorebreakdown *GetHostdbHostsPubkeyOKBodyEntryScorebreakdown `json:"scorebreakdown,omitempty"`

	// Smallest amount of data in bytes that can be uploaded or downloaded to or from the host.
	// Example: 4194304
	Sectorsize int64 `json:"sectorsize,omitempty"`

	// Total amount of storage capacity the host claims it has, in bytes.
	// Example: 35000000000
	Totalstorage int64 `json:"totalstorage,omitempty"`

	// Address at which the host can be paid when forming file contracts.
	// Example: 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789ab
	Unlockhash string `json:"unlockhash,omitempty"`

	// A storage proof window is the number of blocks that the host has to
	// get a storage proof onto the blockchain. The window size is the
	// minimum size of window that the host will accept in a file contract.
	//
	// Example: 144
	Windowsize int64 `json:"windowsize,omitempty"`
}

// Validate validates this get hostdb hosts pubkey o k body entry
func (o *GetHostdbHostsPubkeyOKBodyEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePublickey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScorebreakdown(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHostdbHostsPubkeyOKBodyEntry) validatePublickey(formats strfmt.Registry) error {
	if swag.IsZero(o.Publickey) { // not required
		return nil
	}

	if o.Publickey != nil {
		if err := o.Publickey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHostdbHostsPubkeyOK" + "." + "entry" + "." + "publickey")
			}
			return err
		}
	}

	return nil
}

func (o *GetHostdbHostsPubkeyOKBodyEntry) validateScorebreakdown(formats strfmt.Registry) error {
	if swag.IsZero(o.Scorebreakdown) { // not required
		return nil
	}

	if o.Scorebreakdown != nil {
		if err := o.Scorebreakdown.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHostdbHostsPubkeyOK" + "." + "entry" + "." + "scorebreakdown")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get hostdb hosts pubkey o k body entry based on the context it is used
func (o *GetHostdbHostsPubkeyOKBodyEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePublickey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScorebreakdown(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHostdbHostsPubkeyOKBodyEntry) contextValidatePublickey(ctx context.Context, formats strfmt.Registry) error {

	if o.Publickey != nil {
		if err := o.Publickey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHostdbHostsPubkeyOK" + "." + "entry" + "." + "publickey")
			}
			return err
		}
	}

	return nil
}

func (o *GetHostdbHostsPubkeyOKBodyEntry) contextValidateScorebreakdown(ctx context.Context, formats strfmt.Registry) error {

	if o.Scorebreakdown != nil {
		if err := o.Scorebreakdown.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHostdbHostsPubkeyOK" + "." + "entry" + "." + "scorebreakdown")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBodyEntry) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBodyEntry) UnmarshalBinary(b []byte) error {
	var res GetHostdbHostsPubkeyOKBodyEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetHostdbHostsPubkeyOKBodyEntryPublickey Public key used to identify and verify hosts.
swagger:model GetHostdbHostsPubkeyOKBodyEntryPublickey
*/
type GetHostdbHostsPubkeyOKBodyEntryPublickey struct {

	// Algorithm used for signing and verification. Typically "ed25519".
	// Example: ed25519
	Algorithm string `json:"algorithm,omitempty"`

	// Key used to verify signed host messages.
	// Example: RW50cm9weSBpc24ndCB3aGF0IGl0IHVzZWQgdG8gYmU=
	Key string `json:"key,omitempty"`
}

// Validate validates this get hostdb hosts pubkey o k body entry publickey
func (o *GetHostdbHostsPubkeyOKBodyEntryPublickey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get hostdb hosts pubkey o k body entry publickey based on context it is used
func (o *GetHostdbHostsPubkeyOKBodyEntryPublickey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBodyEntryPublickey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBodyEntryPublickey) UnmarshalBinary(b []byte) error {
	var res GetHostdbHostsPubkeyOKBodyEntryPublickey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetHostdbHostsPubkeyOKBodyEntryScorebreakdown A set of scores as determined by the renter. Generally, the host's final
// final score is all of the values multiplied together. Modified renters may
// have additional criteria that they use to judge a host, or may ignore
// certin criteia. In general, these fields should only be used as a loose
// guide for the score of a host, as every renter sees the world differently
// and uses different metrics to evaluate hosts.
//
swagger:model GetHostdbHostsPubkeyOKBodyEntryScorebreakdown
*/
type GetHostdbHostsPubkeyOKBodyEntryScorebreakdown struct {

	// The multiplier that gets applied to the host based on how long it has been a host. Older hosts typically have a lower penalty.
	// Example: 0.1234
	Ageadjustment int64 `json:"ageadjustment,omitempty"`

	// The multiplier that gets applied to the host based on how much proof-of-burn the host has performed. More burn causes a linear increase in score.
	// Example: 23.456
	Burnadjustment int64 `json:"burnadjustment,omitempty"`

	// The multiplier that gets applied to a host based on how much collateral the host is offering. More collateral is typically better, though above a point it can be detrimental.
	// Example: 23.456
	Collateraladjustment int64 `json:"collateraladjustment,omitempty"`

	// The multiplier that gets applied to a host based on the host's price. Lower prices are almost always better. Below a certain, very low price, there is no advantage.
	// Example: 0.1234
	Priceadjustment int64 `json:"priceadjustment,omitempty"`

	// The overall score for the host. Scores are entriely relative, and are
	// consistent only within the current hostdb. Between different machines,
	// different configurations, and different versions the absolute scores for
	// a given host can be off by many orders of magnitude. When displaying to a
	// human, some form of normalization with respect to the other hosts (for
	// example, divide all scores by the median score of the hosts) is
	// recommended.
	//
	// Example: 123456
	Score int64 `json:"score,omitempty"`

	// The multiplier that gets applied to a host based on how much storage is remaining for the host. More storage remaining is better, to a point.
	// Example: 0.1234
	Storageremainingadjustment int64 `json:"storageremainingadjustment,omitempty"`

	// The multiplier that gets applied to a host based on the uptime percentage of the host. The penalty increases extremely quickly as uptime drops below 90%.
	// Example: 0.1234
	Uptimeadjustment int64 `json:"uptimeadjustment,omitempty"`

	// The multiplier that gets applied to a host based on the version of Sia
	// that they are running. Versions get penalties if there are known bugs,
	// scaling limitations, performance limitations, etc. Generally, the most
	// recent version is always the one with the highest score.
	//
	// Example: 0.1234
	Versionadjustment int64 `json:"versionadjustment,omitempty"`
}

// Validate validates this get hostdb hosts pubkey o k body entry scorebreakdown
func (o *GetHostdbHostsPubkeyOKBodyEntryScorebreakdown) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get hostdb hosts pubkey o k body entry scorebreakdown based on context it is used
func (o *GetHostdbHostsPubkeyOKBodyEntryScorebreakdown) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBodyEntryScorebreakdown) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHostdbHostsPubkeyOKBodyEntryScorebreakdown) UnmarshalBinary(b []byte) error {
	var res GetHostdbHostsPubkeyOKBodyEntryScorebreakdown
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
