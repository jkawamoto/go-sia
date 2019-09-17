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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRenterDirSiapathParams creates a new GetRenterDirSiapathParams object
// with the default values initialized.
func NewGetRenterDirSiapathParams() *GetRenterDirSiapathParams {
	var ()
	return &GetRenterDirSiapathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRenterDirSiapathParamsWithTimeout creates a new GetRenterDirSiapathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRenterDirSiapathParamsWithTimeout(timeout time.Duration) *GetRenterDirSiapathParams {
	var ()
	return &GetRenterDirSiapathParams{

		timeout: timeout,
	}
}

// NewGetRenterDirSiapathParamsWithContext creates a new GetRenterDirSiapathParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRenterDirSiapathParamsWithContext(ctx context.Context) *GetRenterDirSiapathParams {
	var ()
	return &GetRenterDirSiapathParams{

		Context: ctx,
	}
}

// NewGetRenterDirSiapathParamsWithHTTPClient creates a new GetRenterDirSiapathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRenterDirSiapathParamsWithHTTPClient(client *http.Client) *GetRenterDirSiapathParams {
	var ()
	return &GetRenterDirSiapathParams{
		HTTPClient: client,
	}
}

/*GetRenterDirSiapathParams contains all the parameters to send to the API endpoint
for the get renter dir siapath operation typically these are written to a http.Request
*/
type GetRenterDirSiapathParams struct {

	/*Siapath
	  Location of the file in the renter on the network.

	*/
	Siapath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) WithTimeout(timeout time.Duration) *GetRenterDirSiapathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) WithContext(ctx context.Context) *GetRenterDirSiapathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) WithHTTPClient(client *http.Client) *GetRenterDirSiapathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiapath adds the siapath to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) WithSiapath(siapath string) *GetRenterDirSiapathParams {
	o.SetSiapath(siapath)
	return o
}

// SetSiapath adds the siapath to the get renter dir siapath params
func (o *GetRenterDirSiapathParams) SetSiapath(siapath string) {
	o.Siapath = siapath
}

// WriteToRequest writes these params to a swagger request
func (o *GetRenterDirSiapathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param siapath
	if err := r.SetPathParam("siapath", o.Siapath); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
