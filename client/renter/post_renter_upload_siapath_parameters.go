// Code generated by go-swagger; DO NOT EDIT.

package renter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRenterUploadSiapathParams creates a new PostRenterUploadSiapathParams object
// with the default values initialized.
func NewPostRenterUploadSiapathParams() *PostRenterUploadSiapathParams {
	var ()
	return &PostRenterUploadSiapathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRenterUploadSiapathParamsWithTimeout creates a new PostRenterUploadSiapathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRenterUploadSiapathParamsWithTimeout(timeout time.Duration) *PostRenterUploadSiapathParams {
	var ()
	return &PostRenterUploadSiapathParams{

		timeout: timeout,
	}
}

// NewPostRenterUploadSiapathParamsWithContext creates a new PostRenterUploadSiapathParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRenterUploadSiapathParamsWithContext(ctx context.Context) *PostRenterUploadSiapathParams {
	var ()
	return &PostRenterUploadSiapathParams{

		Context: ctx,
	}
}

// NewPostRenterUploadSiapathParamsWithHTTPClient creates a new PostRenterUploadSiapathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRenterUploadSiapathParamsWithHTTPClient(client *http.Client) *PostRenterUploadSiapathParams {
	var ()
	return &PostRenterUploadSiapathParams{
		HTTPClient: client,
	}
}

/*PostRenterUploadSiapathParams contains all the parameters to send to the API endpoint
for the post renter upload siapath operation typically these are written to a http.Request
*/
type PostRenterUploadSiapathParams struct {

	/*Datapieces
	  The number of data pieces to use when erasure coding the file.

	*/
	Datapieces *int64
	/*Paritypieces
	  The number of parity pieces to use when erasure coding the file. Total redundancy of the file is (datapieces+paritypieces)/datapieces.

	*/
	Paritypieces *int64
	/*Siapath
	  Location where the file will reside in the renter on the network. The path must be non-empty, may not include any path traversal strings ("./", "../"), and may not begin with a forward-slash character.

	*/
	Siapath string
	/*Source
	  Location on disk of the file being uploaded.

	*/
	Source string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithTimeout(timeout time.Duration) *PostRenterUploadSiapathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithContext(ctx context.Context) *PostRenterUploadSiapathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithHTTPClient(client *http.Client) *PostRenterUploadSiapathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatapieces adds the datapieces to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithDatapieces(datapieces *int64) *PostRenterUploadSiapathParams {
	o.SetDatapieces(datapieces)
	return o
}

// SetDatapieces adds the datapieces to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetDatapieces(datapieces *int64) {
	o.Datapieces = datapieces
}

// WithParitypieces adds the paritypieces to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithParitypieces(paritypieces *int64) *PostRenterUploadSiapathParams {
	o.SetParitypieces(paritypieces)
	return o
}

// SetParitypieces adds the paritypieces to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetParitypieces(paritypieces *int64) {
	o.Paritypieces = paritypieces
}

// WithSiapath adds the siapath to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithSiapath(siapath string) *PostRenterUploadSiapathParams {
	o.SetSiapath(siapath)
	return o
}

// SetSiapath adds the siapath to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetSiapath(siapath string) {
	o.Siapath = siapath
}

// WithSource adds the source to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) WithSource(source string) *PostRenterUploadSiapathParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the post renter upload siapath params
func (o *PostRenterUploadSiapathParams) SetSource(source string) {
	o.Source = source
}

// WriteToRequest writes these params to a swagger request
func (o *PostRenterUploadSiapathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Datapieces != nil {

		// query param datapieces
		var qrDatapieces int64
		if o.Datapieces != nil {
			qrDatapieces = *o.Datapieces
		}
		qDatapieces := swag.FormatInt64(qrDatapieces)
		if qDatapieces != "" {
			if err := r.SetQueryParam("datapieces", qDatapieces); err != nil {
				return err
			}
		}

	}

	if o.Paritypieces != nil {

		// query param paritypieces
		var qrParitypieces int64
		if o.Paritypieces != nil {
			qrParitypieces = *o.Paritypieces
		}
		qParitypieces := swag.FormatInt64(qrParitypieces)
		if qParitypieces != "" {
			if err := r.SetQueryParam("paritypieces", qParitypieces); err != nil {
				return err
			}
		}

	}

	// path param siapath
	if err := r.SetPathParam("siapath", o.Siapath); err != nil {
		return err
	}

	// query param source
	qrSource := o.Source
	qSource := qrSource
	if qSource != "" {
		if err := r.SetQueryParam("source", qSource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
