// Code generated by go-swagger; DO NOT EDIT.

package daemon

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

// NewGetDaemonVersionParams creates a new GetDaemonVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDaemonVersionParams() *GetDaemonVersionParams {
	return &GetDaemonVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDaemonVersionParamsWithTimeout creates a new GetDaemonVersionParams object
// with the ability to set a timeout on a request.
func NewGetDaemonVersionParamsWithTimeout(timeout time.Duration) *GetDaemonVersionParams {
	return &GetDaemonVersionParams{
		timeout: timeout,
	}
}

// NewGetDaemonVersionParamsWithContext creates a new GetDaemonVersionParams object
// with the ability to set a context for a request.
func NewGetDaemonVersionParamsWithContext(ctx context.Context) *GetDaemonVersionParams {
	return &GetDaemonVersionParams{
		Context: ctx,
	}
}

// NewGetDaemonVersionParamsWithHTTPClient creates a new GetDaemonVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDaemonVersionParamsWithHTTPClient(client *http.Client) *GetDaemonVersionParams {
	return &GetDaemonVersionParams{
		HTTPClient: client,
	}
}

/* GetDaemonVersionParams contains all the parameters to send to the API endpoint
   for the get daemon version operation.

   Typically these are written to a http.Request.
*/
type GetDaemonVersionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get daemon version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDaemonVersionParams) WithDefaults() *GetDaemonVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get daemon version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDaemonVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get daemon version params
func (o *GetDaemonVersionParams) WithTimeout(timeout time.Duration) *GetDaemonVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get daemon version params
func (o *GetDaemonVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get daemon version params
func (o *GetDaemonVersionParams) WithContext(ctx context.Context) *GetDaemonVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get daemon version params
func (o *GetDaemonVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get daemon version params
func (o *GetDaemonVersionParams) WithHTTPClient(client *http.Client) *GetDaemonVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get daemon version params
func (o *GetDaemonVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDaemonVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
