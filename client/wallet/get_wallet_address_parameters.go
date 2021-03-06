// Code generated by go-swagger; DO NOT EDIT.

package wallet

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

// NewGetWalletAddressParams creates a new GetWalletAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWalletAddressParams() *GetWalletAddressParams {
	return &GetWalletAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWalletAddressParamsWithTimeout creates a new GetWalletAddressParams object
// with the ability to set a timeout on a request.
func NewGetWalletAddressParamsWithTimeout(timeout time.Duration) *GetWalletAddressParams {
	return &GetWalletAddressParams{
		timeout: timeout,
	}
}

// NewGetWalletAddressParamsWithContext creates a new GetWalletAddressParams object
// with the ability to set a context for a request.
func NewGetWalletAddressParamsWithContext(ctx context.Context) *GetWalletAddressParams {
	return &GetWalletAddressParams{
		Context: ctx,
	}
}

// NewGetWalletAddressParamsWithHTTPClient creates a new GetWalletAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWalletAddressParamsWithHTTPClient(client *http.Client) *GetWalletAddressParams {
	return &GetWalletAddressParams{
		HTTPClient: client,
	}
}

/* GetWalletAddressParams contains all the parameters to send to the API endpoint
   for the get wallet address operation.

   Typically these are written to a http.Request.
*/
type GetWalletAddressParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wallet address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWalletAddressParams) WithDefaults() *GetWalletAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wallet address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWalletAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wallet address params
func (o *GetWalletAddressParams) WithTimeout(timeout time.Duration) *GetWalletAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wallet address params
func (o *GetWalletAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wallet address params
func (o *GetWalletAddressParams) WithContext(ctx context.Context) *GetWalletAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wallet address params
func (o *GetWalletAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wallet address params
func (o *GetWalletAddressParams) WithHTTPClient(client *http.Client) *GetWalletAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wallet address params
func (o *GetWalletAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWalletAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
