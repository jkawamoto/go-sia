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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostWalletSiacoinsParams creates a new PostWalletSiacoinsParams object
// with the default values initialized.
func NewPostWalletSiacoinsParams() *PostWalletSiacoinsParams {
	var ()
	return &PostWalletSiacoinsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWalletSiacoinsParamsWithTimeout creates a new PostWalletSiacoinsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWalletSiacoinsParamsWithTimeout(timeout time.Duration) *PostWalletSiacoinsParams {
	var ()
	return &PostWalletSiacoinsParams{

		timeout: timeout,
	}
}

// NewPostWalletSiacoinsParamsWithContext creates a new PostWalletSiacoinsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWalletSiacoinsParamsWithContext(ctx context.Context) *PostWalletSiacoinsParams {
	var ()
	return &PostWalletSiacoinsParams{

		Context: ctx,
	}
}

// NewPostWalletSiacoinsParamsWithHTTPClient creates a new PostWalletSiacoinsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWalletSiacoinsParamsWithHTTPClient(client *http.Client) *PostWalletSiacoinsParams {
	var ()
	return &PostWalletSiacoinsParams{
		HTTPClient: client,
	}
}

/*PostWalletSiacoinsParams contains all the parameters to send to the API endpoint
for the post wallet siacoins operation typically these are written to a http.Request
*/
type PostWalletSiacoinsParams struct {

	/*Amount
	  Number of hastings being sent. A hasting is the smallest unit in Sia. There are 10^24 hastings in a siacoin.

	*/
	Amount *int64
	/*Destination
	  Address that is receiving the coins.

	*/
	Destination *string
	/*Outputs
	  JSON array of outputs. The structure of each output is: {"unlockhash": "<destination>", "value": "<amount>"}

	*/
	Outputs *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) WithTimeout(timeout time.Duration) *PostWalletSiacoinsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) WithContext(ctx context.Context) *PostWalletSiacoinsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) WithHTTPClient(client *http.Client) *PostWalletSiacoinsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmount adds the amount to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) WithAmount(amount *int64) *PostWalletSiacoinsParams {
	o.SetAmount(amount)
	return o
}

// SetAmount adds the amount to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) SetAmount(amount *int64) {
	o.Amount = amount
}

// WithDestination adds the destination to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) WithDestination(destination *string) *PostWalletSiacoinsParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) SetDestination(destination *string) {
	o.Destination = destination
}

// WithOutputs adds the outputs to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) WithOutputs(outputs *string) *PostWalletSiacoinsParams {
	o.SetOutputs(outputs)
	return o
}

// SetOutputs adds the outputs to the post wallet siacoins params
func (o *PostWalletSiacoinsParams) SetOutputs(outputs *string) {
	o.Outputs = outputs
}

// WriteToRequest writes these params to a swagger request
func (o *PostWalletSiacoinsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Amount != nil {

		// query param amount
		var qrAmount int64
		if o.Amount != nil {
			qrAmount = *o.Amount
		}
		qAmount := swag.FormatInt64(qrAmount)
		if qAmount != "" {
			if err := r.SetQueryParam("amount", qAmount); err != nil {
				return err
			}
		}

	}

	if o.Destination != nil {

		// query param destination
		var qrDestination string
		if o.Destination != nil {
			qrDestination = *o.Destination
		}
		qDestination := qrDestination
		if qDestination != "" {
			if err := r.SetQueryParam("destination", qDestination); err != nil {
				return err
			}
		}

	}

	if o.Outputs != nil {

		// query param outputs
		var qrOutputs string
		if o.Outputs != nil {
			qrOutputs = *o.Outputs
		}
		qOutputs := qrOutputs
		if qOutputs != "" {
			if err := r.SetQueryParam("outputs", qOutputs); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}