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
	"github.com/go-openapi/strfmt"
)

// NewGetRenterFileSiapathParams creates a new GetRenterFileSiapathParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRenterFileSiapathParams() *GetRenterFileSiapathParams {
	return &GetRenterFileSiapathParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRenterFileSiapathParamsWithTimeout creates a new GetRenterFileSiapathParams object
// with the ability to set a timeout on a request.
func NewGetRenterFileSiapathParamsWithTimeout(timeout time.Duration) *GetRenterFileSiapathParams {
	return &GetRenterFileSiapathParams{
		timeout: timeout,
	}
}

// NewGetRenterFileSiapathParamsWithContext creates a new GetRenterFileSiapathParams object
// with the ability to set a context for a request.
func NewGetRenterFileSiapathParamsWithContext(ctx context.Context) *GetRenterFileSiapathParams {
	return &GetRenterFileSiapathParams{
		Context: ctx,
	}
}

// NewGetRenterFileSiapathParamsWithHTTPClient creates a new GetRenterFileSiapathParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRenterFileSiapathParamsWithHTTPClient(client *http.Client) *GetRenterFileSiapathParams {
	return &GetRenterFileSiapathParams{
		HTTPClient: client,
	}
}

/* GetRenterFileSiapathParams contains all the parameters to send to the API endpoint
   for the get renter file siapath operation.

   Typically these are written to a http.Request.
*/
type GetRenterFileSiapathParams struct {

	/* Siapath.

	   Location of the file in the renter on the network.
	*/
	Siapath string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get renter file siapath params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRenterFileSiapathParams) WithDefaults() *GetRenterFileSiapathParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get renter file siapath params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRenterFileSiapathParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get renter file siapath params
func (o *GetRenterFileSiapathParams) WithTimeout(timeout time.Duration) *GetRenterFileSiapathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get renter file siapath params
func (o *GetRenterFileSiapathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get renter file siapath params
func (o *GetRenterFileSiapathParams) WithContext(ctx context.Context) *GetRenterFileSiapathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get renter file siapath params
func (o *GetRenterFileSiapathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get renter file siapath params
func (o *GetRenterFileSiapathParams) WithHTTPClient(client *http.Client) *GetRenterFileSiapathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get renter file siapath params
func (o *GetRenterFileSiapathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiapath adds the siapath to the get renter file siapath params
func (o *GetRenterFileSiapathParams) WithSiapath(siapath string) *GetRenterFileSiapathParams {
	o.SetSiapath(siapath)
	return o
}

// SetSiapath adds the siapath to the get renter file siapath params
func (o *GetRenterFileSiapathParams) SetSiapath(siapath string) {
	o.Siapath = siapath
}

// WriteToRequest writes these params to a swagger request
func (o *GetRenterFileSiapathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
