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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostWalletSeedParams creates a new PostWalletSeedParams object
// with the default values initialized.
func NewPostWalletSeedParams() *PostWalletSeedParams {
	var ()
	return &PostWalletSeedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWalletSeedParamsWithTimeout creates a new PostWalletSeedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWalletSeedParamsWithTimeout(timeout time.Duration) *PostWalletSeedParams {
	var ()
	return &PostWalletSeedParams{

		timeout: timeout,
	}
}

// NewPostWalletSeedParamsWithContext creates a new PostWalletSeedParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWalletSeedParamsWithContext(ctx context.Context) *PostWalletSeedParams {
	var ()
	return &PostWalletSeedParams{

		Context: ctx,
	}
}

// NewPostWalletSeedParamsWithHTTPClient creates a new PostWalletSeedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWalletSeedParamsWithHTTPClient(client *http.Client) *PostWalletSeedParams {
	var ()
	return &PostWalletSeedParams{
		HTTPClient: client,
	}
}

/*PostWalletSeedParams contains all the parameters to send to the API endpoint
for the post wallet seed operation typically these are written to a http.Request
*/
type PostWalletSeedParams struct {

	/*Dictionary
	  Name of the dictionary that should be used when encoding the seed. 'english' is the most common choice when picking a dictionary.

	*/
	Dictionary *string
	/*Encryptionpassword
	  Key used to encrypt the new seed when it is saved to disk.

	*/
	Encryptionpassword string
	/*Seed
	  Dictionary-encoded phrase that corresponds to the seed being added to the wallet.

	*/
	Seed *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post wallet seed params
func (o *PostWalletSeedParams) WithTimeout(timeout time.Duration) *PostWalletSeedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wallet seed params
func (o *PostWalletSeedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wallet seed params
func (o *PostWalletSeedParams) WithContext(ctx context.Context) *PostWalletSeedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wallet seed params
func (o *PostWalletSeedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wallet seed params
func (o *PostWalletSeedParams) WithHTTPClient(client *http.Client) *PostWalletSeedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wallet seed params
func (o *PostWalletSeedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDictionary adds the dictionary to the post wallet seed params
func (o *PostWalletSeedParams) WithDictionary(dictionary *string) *PostWalletSeedParams {
	o.SetDictionary(dictionary)
	return o
}

// SetDictionary adds the dictionary to the post wallet seed params
func (o *PostWalletSeedParams) SetDictionary(dictionary *string) {
	o.Dictionary = dictionary
}

// WithEncryptionpassword adds the encryptionpassword to the post wallet seed params
func (o *PostWalletSeedParams) WithEncryptionpassword(encryptionpassword string) *PostWalletSeedParams {
	o.SetEncryptionpassword(encryptionpassword)
	return o
}

// SetEncryptionpassword adds the encryptionpassword to the post wallet seed params
func (o *PostWalletSeedParams) SetEncryptionpassword(encryptionpassword string) {
	o.Encryptionpassword = encryptionpassword
}

// WithSeed adds the seed to the post wallet seed params
func (o *PostWalletSeedParams) WithSeed(seed *string) *PostWalletSeedParams {
	o.SetSeed(seed)
	return o
}

// SetSeed adds the seed to the post wallet seed params
func (o *PostWalletSeedParams) SetSeed(seed *string) {
	o.Seed = seed
}

// WriteToRequest writes these params to a swagger request
func (o *PostWalletSeedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dictionary != nil {

		// query param dictionary
		var qrDictionary string
		if o.Dictionary != nil {
			qrDictionary = *o.Dictionary
		}
		qDictionary := qrDictionary
		if qDictionary != "" {
			if err := r.SetQueryParam("dictionary", qDictionary); err != nil {
				return err
			}
		}

	}

	// query param encryptionpassword
	qrEncryptionpassword := o.Encryptionpassword
	qEncryptionpassword := qrEncryptionpassword
	if qEncryptionpassword != "" {
		if err := r.SetQueryParam("encryptionpassword", qEncryptionpassword); err != nil {
			return err
		}
	}

	if o.Seed != nil {

		// query param seed
		var qrSeed string
		if o.Seed != nil {
			qrSeed = *o.Seed
		}
		qSeed := qrSeed
		if qSeed != "" {
			if err := r.SetQueryParam("seed", qSeed); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
