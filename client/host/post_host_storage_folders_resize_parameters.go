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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostHostStorageFoldersResizeParams creates a new PostHostStorageFoldersResizeParams object
// with the default values initialized.
func NewPostHostStorageFoldersResizeParams() *PostHostStorageFoldersResizeParams {
	var ()
	return &PostHostStorageFoldersResizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHostStorageFoldersResizeParamsWithTimeout creates a new PostHostStorageFoldersResizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHostStorageFoldersResizeParamsWithTimeout(timeout time.Duration) *PostHostStorageFoldersResizeParams {
	var ()
	return &PostHostStorageFoldersResizeParams{

		timeout: timeout,
	}
}

// NewPostHostStorageFoldersResizeParamsWithContext creates a new PostHostStorageFoldersResizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHostStorageFoldersResizeParamsWithContext(ctx context.Context) *PostHostStorageFoldersResizeParams {
	var ()
	return &PostHostStorageFoldersResizeParams{

		Context: ctx,
	}
}

// NewPostHostStorageFoldersResizeParamsWithHTTPClient creates a new PostHostStorageFoldersResizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHostStorageFoldersResizeParamsWithHTTPClient(client *http.Client) *PostHostStorageFoldersResizeParams {
	var ()
	return &PostHostStorageFoldersResizeParams{
		HTTPClient: client,
	}
}

/*PostHostStorageFoldersResizeParams contains all the parameters to send to the API endpoint
for the post host storage folders resize operation typically these are written to a http.Request
*/
type PostHostStorageFoldersResizeParams struct {

	/*Newsize
	  Desired new size of the storage folder. This will be the new capacity of the storage folder.

	*/
	Newsize int64
	/*Path
	  Local path on disk to the storage folder to resize.

	*/
	Path string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) WithTimeout(timeout time.Duration) *PostHostStorageFoldersResizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) WithContext(ctx context.Context) *PostHostStorageFoldersResizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) WithHTTPClient(client *http.Client) *PostHostStorageFoldersResizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewsize adds the newsize to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) WithNewsize(newsize int64) *PostHostStorageFoldersResizeParams {
	o.SetNewsize(newsize)
	return o
}

// SetNewsize adds the newsize to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) SetNewsize(newsize int64) {
	o.Newsize = newsize
}

// WithPath adds the path to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) WithPath(path string) *PostHostStorageFoldersResizeParams {
	o.SetPath(path)
	return o
}

// SetPath adds the path to the post host storage folders resize params
func (o *PostHostStorageFoldersResizeParams) SetPath(path string) {
	o.Path = path
}

// WriteToRequest writes these params to a swagger request
func (o *PostHostStorageFoldersResizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param newsize
	qrNewsize := o.Newsize
	qNewsize := swag.FormatInt64(qrNewsize)
	if qNewsize != "" {
		if err := r.SetQueryParam("newsize", qNewsize); err != nil {
			return err
		}
	}

	// query param path
	qrPath := o.Path
	qPath := qrPath
	if qPath != "" {
		if err := r.SetQueryParam("path", qPath); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
