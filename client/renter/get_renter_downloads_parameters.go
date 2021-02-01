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

// NewGetRenterDownloadsParams creates a new GetRenterDownloadsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRenterDownloadsParams() *GetRenterDownloadsParams {
	return &GetRenterDownloadsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRenterDownloadsParamsWithTimeout creates a new GetRenterDownloadsParams object
// with the ability to set a timeout on a request.
func NewGetRenterDownloadsParamsWithTimeout(timeout time.Duration) *GetRenterDownloadsParams {
	return &GetRenterDownloadsParams{
		timeout: timeout,
	}
}

// NewGetRenterDownloadsParamsWithContext creates a new GetRenterDownloadsParams object
// with the ability to set a context for a request.
func NewGetRenterDownloadsParamsWithContext(ctx context.Context) *GetRenterDownloadsParams {
	return &GetRenterDownloadsParams{
		Context: ctx,
	}
}

// NewGetRenterDownloadsParamsWithHTTPClient creates a new GetRenterDownloadsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRenterDownloadsParamsWithHTTPClient(client *http.Client) *GetRenterDownloadsParams {
	return &GetRenterDownloadsParams{
		HTTPClient: client,
	}
}

/* GetRenterDownloadsParams contains all the parameters to send to the API endpoint
   for the get renter downloads operation.

   Typically these are written to a http.Request.
*/
type GetRenterDownloadsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get renter downloads params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRenterDownloadsParams) WithDefaults() *GetRenterDownloadsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get renter downloads params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRenterDownloadsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get renter downloads params
func (o *GetRenterDownloadsParams) WithTimeout(timeout time.Duration) *GetRenterDownloadsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get renter downloads params
func (o *GetRenterDownloadsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get renter downloads params
func (o *GetRenterDownloadsParams) WithContext(ctx context.Context) *GetRenterDownloadsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get renter downloads params
func (o *GetRenterDownloadsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get renter downloads params
func (o *GetRenterDownloadsParams) WithHTTPClient(client *http.Client) *GetRenterDownloadsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get renter downloads params
func (o *GetRenterDownloadsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRenterDownloadsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
