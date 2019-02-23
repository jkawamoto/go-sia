// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostHostStorageSectorsDeleteMerklerootParams creates a new PostHostStorageSectorsDeleteMerklerootParams object
// with the default values initialized.
func NewPostHostStorageSectorsDeleteMerklerootParams() *PostHostStorageSectorsDeleteMerklerootParams {
	var ()
	return &PostHostStorageSectorsDeleteMerklerootParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHostStorageSectorsDeleteMerklerootParamsWithTimeout creates a new PostHostStorageSectorsDeleteMerklerootParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHostStorageSectorsDeleteMerklerootParamsWithTimeout(timeout time.Duration) *PostHostStorageSectorsDeleteMerklerootParams {
	var ()
	return &PostHostStorageSectorsDeleteMerklerootParams{

		timeout: timeout,
	}
}

// NewPostHostStorageSectorsDeleteMerklerootParamsWithContext creates a new PostHostStorageSectorsDeleteMerklerootParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHostStorageSectorsDeleteMerklerootParamsWithContext(ctx context.Context) *PostHostStorageSectorsDeleteMerklerootParams {
	var ()
	return &PostHostStorageSectorsDeleteMerklerootParams{

		Context: ctx,
	}
}

// NewPostHostStorageSectorsDeleteMerklerootParamsWithHTTPClient creates a new PostHostStorageSectorsDeleteMerklerootParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHostStorageSectorsDeleteMerklerootParamsWithHTTPClient(client *http.Client) *PostHostStorageSectorsDeleteMerklerootParams {
	var ()
	return &PostHostStorageSectorsDeleteMerklerootParams{
		HTTPClient: client,
	}
}

/*PostHostStorageSectorsDeleteMerklerootParams contains all the parameters to send to the API endpoint
for the post host storage sectors delete merkleroot operation typically these are written to a http.Request
*/
type PostHostStorageSectorsDeleteMerklerootParams struct {

	/*Merkleroot
	  Merkleroot of the sector to delete.

	*/
	Merkleroot string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) WithTimeout(timeout time.Duration) *PostHostStorageSectorsDeleteMerklerootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) WithContext(ctx context.Context) *PostHostStorageSectorsDeleteMerklerootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) WithHTTPClient(client *http.Client) *PostHostStorageSectorsDeleteMerklerootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMerkleroot adds the merkleroot to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) WithMerkleroot(merkleroot string) *PostHostStorageSectorsDeleteMerklerootParams {
	o.SetMerkleroot(merkleroot)
	return o
}

// SetMerkleroot adds the merkleroot to the post host storage sectors delete merkleroot params
func (o *PostHostStorageSectorsDeleteMerklerootParams) SetMerkleroot(merkleroot string) {
	o.Merkleroot = merkleroot
}

// WriteToRequest writes these params to a swagger request
func (o *PostHostStorageSectorsDeleteMerklerootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param merkleroot
	if err := r.SetPathParam("merkleroot", o.Merkleroot); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
