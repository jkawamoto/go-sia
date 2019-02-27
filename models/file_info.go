// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FileInfo file info
// swagger:model FileInfo
type FileInfo struct {

	// true if the file is available for download. Files may be available before they are completely uploaded.
	Available bool `json:"available,omitempty"`

	// Block height at which the file ceases availability.
	Expiration int64 `json:"expiration,omitempty"`

	// Size of the file in bytes.
	Filesize int64 `json:"filesize,omitempty"`

	// Path to the local file on disk.
	Localpath string `json:"localpath,omitempty"`

	// indicates if the source file is found on disk
	Ondisk bool `json:"ondisk,omitempty"`

	// indicates if the siafile is recoverable
	Recoverable bool `json:"recoverable,omitempty"`

	// Average redundancy of the file on the network. Redundancy is calculated by dividing the amount of data uploaded in the file's open contracts by the size of the file. Redundancy does not necessarily correspond to availability. Specifically, a redundancy >= 1 does not indicate the file is available as there could be a chunk of the file with 0 redundancy.
	Redundancy float64 `json:"redundancy,omitempty"`

	// true if the file's contracts will be automatically renewed by the renter.
	Renewing bool `json:"renewing,omitempty"`

	// Path to the file in the renter on the network.
	Siapath string `json:"siapath,omitempty"`

	// Total number of bytes successfully uploaded via current file contracts. This number includes padding and rendundancy, so a file with a size of 8192 bytes might be padded to 40 MiB and, with a redundancy of 5, encoded to 200 MiB for upload.
	Uploadedbytes int64 `json:"uploadedbytes,omitempty"`

	// Percentage of the file uploaded, including redundancy. Uploading has completed when uploadprogress is 100. Files may be available for download before upload progress is 100.
	Uploadprogress float64 `json:"uploadprogress,omitempty"`
}

// Validate validates this file info
func (m *FileInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileInfo) UnmarshalBinary(b []byte) error {
	var res FileInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
