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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDaemonConstantsParams creates a new GetDaemonConstantsParams object
// with the default values initialized.
func NewGetDaemonConstantsParams() *GetDaemonConstantsParams {

	return &GetDaemonConstantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDaemonConstantsParamsWithTimeout creates a new GetDaemonConstantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDaemonConstantsParamsWithTimeout(timeout time.Duration) *GetDaemonConstantsParams {

	return &GetDaemonConstantsParams{

		timeout: timeout,
	}
}

// NewGetDaemonConstantsParamsWithContext creates a new GetDaemonConstantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDaemonConstantsParamsWithContext(ctx context.Context) *GetDaemonConstantsParams {

	return &GetDaemonConstantsParams{

		Context: ctx,
	}
}

// NewGetDaemonConstantsParamsWithHTTPClient creates a new GetDaemonConstantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDaemonConstantsParamsWithHTTPClient(client *http.Client) *GetDaemonConstantsParams {

	return &GetDaemonConstantsParams{
		HTTPClient: client,
	}
}

/*GetDaemonConstantsParams contains all the parameters to send to the API endpoint
for the get daemon constants operation typically these are written to a http.Request
*/
type GetDaemonConstantsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get daemon constants params
func (o *GetDaemonConstantsParams) WithTimeout(timeout time.Duration) *GetDaemonConstantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get daemon constants params
func (o *GetDaemonConstantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get daemon constants params
func (o *GetDaemonConstantsParams) WithContext(ctx context.Context) *GetDaemonConstantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get daemon constants params
func (o *GetDaemonConstantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get daemon constants params
func (o *GetDaemonConstantsParams) WithHTTPClient(client *http.Client) *GetDaemonConstantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get daemon constants params
func (o *GetDaemonConstantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDaemonConstantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}