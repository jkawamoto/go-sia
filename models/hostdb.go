// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Hostdb hostdb
//
// swagger:model Hostdb
type Hostdb struct {

	// hosts
	Hosts []*HostdbHostsItems0 `json:"hosts"`
}

// Validate validates this hostdb
func (m *Hostdb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHosts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hostdb) validateHosts(formats strfmt.Registry) error {
	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	for i := 0; i < len(m.Hosts); i++ {
		if swag.IsZero(m.Hosts[i]) { // not required
			continue
		}

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hostdb based on the context it is used
func (m *Hostdb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHosts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Hostdb) contextValidateHosts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hosts); i++ {

		if m.Hosts[i] != nil {
			if err := m.Hosts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hosts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Hostdb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Hostdb) UnmarshalBinary(b []byte) error {
	var res Hostdb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HostdbHostsItems0 hostdb hosts items0
//
// swagger:model HostdbHostsItems0
type HostdbHostsItems0 struct {

	// true if the host is accepting new contracts.
	// Example: true
	Acceptingcontracts bool `json:"acceptingcontracts,omitempty"`

	// The maximum amount of money that the host will put up as collateral for storage that is contracted by the renter.
	// Example: 20000000000
	Collateral string `json:"collateral,omitempty"`

	// The price that a renter has to pay to create a contract with the host. The payment is intended to cover transaction fees for the file contract revision and the storage proof that the host will be submitting to the blockchain.
	// Example: 1000000000000000000000000
	Contractprice string `json:"contractprice,omitempty"`

	// The price that a renter has to pay when downloading data from the host.
	// Example: 35000000000000
	Downloadbandwidthprice string `json:"downloadbandwidthprice,omitempty"`

	// Indicates if the host is currently being filtered from the HostDB
	// Example: false
	Filtered bool `json:"filtered,omitempty"`

	// Firstseen is the last block height at which this host was announced.
	// Example: 160000
	Firstseen int64 `json:"firstseen,omitempty"`

	// Total amount of time the host has been offline in nanosecounds.
	// Example: 0
	Historicdowntime int64 `json:"historicdowntime,omitempty"`

	// Number of historic failed interactions with the host.
	// Example: 0
	Historicfailedinteractions float64 `json:"historicfailedinteractions,omitempty"`

	// Number of historic successful interactions with the host.
	// Example: 5
	Historicsuccessfulinteractions float64 `json:"historicsuccessfulinteractions,omitempty"`

	// Total amount of time the host has been online in nanosecounds.
	// Example: 41634520900246580
	Historicuptime int64 `json:"historicuptime,omitempty"`

	// List of IP subnet masks used by the host. For IPv4 the /24 and for IPv6 the /54 subnet mask is used. A host can have either one IPv4 or one IPv6 subnet or one of each. E.g. these lists are valid: [ "IPv4" ], [ "IPv6" ] or [ "IPv4", "IPv6" ]. The following lists are invalid: [ "IPv4", "IPv4" ], [ "IPv4", "IPv6", "IPv6" ]. Hosts with an invalid list are ignored.
	Ipnets []string `json:"ipnets"`

	// The last time that the interactions within scanhistory have been compressed into the historic ones.
	// Example: 174900
	Lasthistoricupdate int64 `json:"lasthistoricupdate,omitempty"`

	// The last time the list of IP subnet masks was updated. When equal subnet masks are found for different hosts, the host that occupies the subnet mask for a longer time is preferred.
	// Example: 2015-01-01T08:00:00.000000000+04:00
	// Format: date-time
	Lastipnetchange strfmt.DateTime `json:"lastipnetchange,omitempty"`

	// The maximum amount of collateral that the host will put into a single file contract.
	// Example: 1000000000000000000000000000
	Maxcollateral string `json:"maxcollateral,omitempty"`

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
	Publickey *HostdbHostsItems0Publickey `json:"publickey,omitempty"`

	// The string representation of the full public key, used when calling /hostdb/hosts.
	// Example: ed25519:1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
	Publickeystring string `json:"publickeystring,omitempty"`

	// Number of recent failed interactions with the host.
	// Example: 0
	Recentfailedinteractions int64 `json:"recentfailedinteractions,omitempty"`

	// Number of recent successful interactions with the host.
	// Example: 0
	Recentsuccessfulinteractions int64 `json:"recentsuccessfulinteractions,omitempty"`

	// Unused storage capacity the host claims it has, in bytes.
	// Example: 35000000000
	Remainingstorage int64 `json:"remainingstorage,omitempty"`

	// The revision number indicates to the renter what iteration of settings the host is currently at. Settings are generally signed. If the renter has multiple conflicting copies of settings from the host, the renter can expect the one with the higher revision number to be more recent.
	// Example: 12733798
	Revisionnumber int64 `json:"revisionnumber,omitempty"`

	// Measurements that have been taken on the host. The most recent measurements are kept in full detail.
	Scanhistory []*HostdbHostsItems0ScanhistoryItems0 `json:"scanhistory"`

	// Smallest amount of data in bytes that can be uploaded or downloaded to or from the host.
	// Example: 4194304
	Sectorsize int64 `json:"sectorsize,omitempty"`

	// The price that a renter has to pay to store files with the host.
	// Example: 14000000000
	Storageprice string `json:"storageprice,omitempty"`

	// Total amount of storage capacity the host claims it has, in bytes.
	// Example: 35000000000
	Totalstorage int64 `json:"totalstorage,omitempty"`

	// Address at which the host can be paid when forming file contracts.
	// Example: 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789ab
	Unlockhash string `json:"unlockhash,omitempty"`

	// The price that a renter has to pay when uploading data to the host.
	Uploadbandwidthprice string `json:"uploadbandwidthprice,omitempty"`

	// The version of the host.
	// Example: 1.3.4
	Version string `json:"version,omitempty"`

	// A storage proof window is the number of blocks that the host has to get a storage proof onto the blockchain. The window size is the minimum size of window that the host will accept in a file contract.
	// Example: 144
	Windowsize int64 `json:"windowsize,omitempty"`
}

// Validate validates this hostdb hosts items0
func (m *HostdbHostsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastipnetchange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublickey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanhistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostdbHostsItems0) validateLastipnetchange(formats strfmt.Registry) error {
	if swag.IsZero(m.Lastipnetchange) { // not required
		return nil
	}

	if err := validate.FormatOf("lastipnetchange", "body", "date-time", m.Lastipnetchange.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HostdbHostsItems0) validatePublickey(formats strfmt.Registry) error {
	if swag.IsZero(m.Publickey) { // not required
		return nil
	}

	if m.Publickey != nil {
		if err := m.Publickey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publickey")
			}
			return err
		}
	}

	return nil
}

func (m *HostdbHostsItems0) validateScanhistory(formats strfmt.Registry) error {
	if swag.IsZero(m.Scanhistory) { // not required
		return nil
	}

	for i := 0; i < len(m.Scanhistory); i++ {
		if swag.IsZero(m.Scanhistory[i]) { // not required
			continue
		}

		if m.Scanhistory[i] != nil {
			if err := m.Scanhistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scanhistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hostdb hosts items0 based on the context it is used
func (m *HostdbHostsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublickey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScanhistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostdbHostsItems0) contextValidatePublickey(ctx context.Context, formats strfmt.Registry) error {

	if m.Publickey != nil {
		if err := m.Publickey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publickey")
			}
			return err
		}
	}

	return nil
}

func (m *HostdbHostsItems0) contextValidateScanhistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Scanhistory); i++ {

		if m.Scanhistory[i] != nil {
			if err := m.Scanhistory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scanhistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostdbHostsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostdbHostsItems0) UnmarshalBinary(b []byte) error {
	var res HostdbHostsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HostdbHostsItems0Publickey Public key used to identify and verify hosts.
//
// swagger:model HostdbHostsItems0Publickey
type HostdbHostsItems0Publickey struct {

	// Algorithm used for signing and verification. Typically "ed25519".
	// Example: ed25519
	Algorithm string `json:"algorithm,omitempty"`

	// Key used to verify signed host messages.
	// Example: RW50cm9weSBpc24ndCB3aGF0IGl0IHVzZWQgdG8gYmU=
	Key string `json:"key,omitempty"`
}

// Validate validates this hostdb hosts items0 publickey
func (m *HostdbHostsItems0Publickey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hostdb hosts items0 publickey based on context it is used
func (m *HostdbHostsItems0Publickey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostdbHostsItems0Publickey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostdbHostsItems0Publickey) UnmarshalBinary(b []byte) error {
	var res HostdbHostsItems0Publickey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HostdbHostsItems0ScanhistoryItems0 hostdb hosts items0 scanhistory items0
//
// swagger:model HostdbHostsItems0ScanhistoryItems0
type HostdbHostsItems0ScanhistoryItems0 struct {

	// success
	// Example: true
	Success bool `json:"success,omitempty"`

	// timestamp
	// Example: 2018-09-23T08:00:00.000000000+04:00
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this hostdb hosts items0 scanhistory items0
func (m *HostdbHostsItems0ScanhistoryItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostdbHostsItems0ScanhistoryItems0) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hostdb hosts items0 scanhistory items0 based on context it is used
func (m *HostdbHostsItems0ScanhistoryItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostdbHostsItems0ScanhistoryItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostdbHostsItems0ScanhistoryItems0) UnmarshalBinary(b []byte) error {
	var res HostdbHostsItems0ScanhistoryItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
