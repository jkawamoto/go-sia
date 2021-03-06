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
	"github.com/go-openapi/strfmt"
)

// NewGetHostStorageParams creates a new GetHostStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostStorageParams() *GetHostStorageParams {
	return &GetHostStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostStorageParamsWithTimeout creates a new GetHostStorageParams object
// with the ability to set a timeout on a request.
func NewGetHostStorageParamsWithTimeout(timeout time.Duration) *GetHostStorageParams {
	return &GetHostStorageParams{
		timeout: timeout,
	}
}

// NewGetHostStorageParamsWithContext creates a new GetHostStorageParams object
// with the ability to set a context for a request.
func NewGetHostStorageParamsWithContext(ctx context.Context) *GetHostStorageParams {
	return &GetHostStorageParams{
		Context: ctx,
	}
}

// NewGetHostStorageParamsWithHTTPClient creates a new GetHostStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostStorageParamsWithHTTPClient(client *http.Client) *GetHostStorageParams {
	return &GetHostStorageParams{
		HTTPClient: client,
	}
}

/* GetHostStorageParams contains all the parameters to send to the API endpoint
   for the get host storage operation.

   Typically these are written to a http.Request.
*/
type GetHostStorageParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostStorageParams) WithDefaults() *GetHostStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host storage params
func (o *GetHostStorageParams) WithTimeout(timeout time.Duration) *GetHostStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host storage params
func (o *GetHostStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host storage params
func (o *GetHostStorageParams) WithContext(ctx context.Context) *GetHostStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host storage params
func (o *GetHostStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host storage params
func (o *GetHostStorageParams) WithHTTPClient(client *http.Client) *GetHostStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host storage params
func (o *GetHostStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
