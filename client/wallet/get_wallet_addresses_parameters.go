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

// NewGetWalletAddressesParams creates a new GetWalletAddressesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWalletAddressesParams() *GetWalletAddressesParams {
	return &GetWalletAddressesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWalletAddressesParamsWithTimeout creates a new GetWalletAddressesParams object
// with the ability to set a timeout on a request.
func NewGetWalletAddressesParamsWithTimeout(timeout time.Duration) *GetWalletAddressesParams {
	return &GetWalletAddressesParams{
		timeout: timeout,
	}
}

// NewGetWalletAddressesParamsWithContext creates a new GetWalletAddressesParams object
// with the ability to set a context for a request.
func NewGetWalletAddressesParamsWithContext(ctx context.Context) *GetWalletAddressesParams {
	return &GetWalletAddressesParams{
		Context: ctx,
	}
}

// NewGetWalletAddressesParamsWithHTTPClient creates a new GetWalletAddressesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWalletAddressesParamsWithHTTPClient(client *http.Client) *GetWalletAddressesParams {
	return &GetWalletAddressesParams{
		HTTPClient: client,
	}
}

/* GetWalletAddressesParams contains all the parameters to send to the API endpoint
   for the get wallet addresses operation.

   Typically these are written to a http.Request.
*/
type GetWalletAddressesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wallet addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWalletAddressesParams) WithDefaults() *GetWalletAddressesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wallet addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWalletAddressesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wallet addresses params
func (o *GetWalletAddressesParams) WithTimeout(timeout time.Duration) *GetWalletAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wallet addresses params
func (o *GetWalletAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wallet addresses params
func (o *GetWalletAddressesParams) WithContext(ctx context.Context) *GetWalletAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wallet addresses params
func (o *GetWalletAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wallet addresses params
func (o *GetWalletAddressesParams) WithHTTPClient(client *http.Client) *GetWalletAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wallet addresses params
func (o *GetWalletAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetWalletAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
