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

// NewPostHostParams creates a new PostHostParams object
// with the default values initialized.
func NewPostHostParams() *PostHostParams {
	var ()
	return &PostHostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostHostParamsWithTimeout creates a new PostHostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostHostParamsWithTimeout(timeout time.Duration) *PostHostParams {
	var ()
	return &PostHostParams{

		timeout: timeout,
	}
}

// NewPostHostParamsWithContext creates a new PostHostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostHostParamsWithContext(ctx context.Context) *PostHostParams {
	var ()
	return &PostHostParams{

		Context: ctx,
	}
}

// NewPostHostParamsWithHTTPClient creates a new PostHostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostHostParamsWithHTTPClient(client *http.Client) *PostHostParams {
	var ()
	return &PostHostParams{
		HTTPClient: client,
	}
}

/*PostHostParams contains all the parameters to send to the API endpoint
for the post host operation typically these are written to a http.Request
*/
type PostHostParams struct {

	/*Acceptingcontracts
	  When set to true, the host will accept new file contracts if the
	terms are reasonable. When set to false, the host will not accept new
	file contracts at all.


	*/
	Acceptingcontracts *bool
	/*Collateral
	  The maximum amount of money that the host will put up as collateral per byte per block of storage that is contracted by the renter.

	*/
	Collateral *string
	/*Collateralbudget
	  The total amount of money that the host will allocate to collateral across all file contracts.

	*/
	Collateralbudget *string
	/*Maxcollateral
	  The maximum amount of collateral that the host will put into a single file contract.

	*/
	Maxcollateral *string
	/*Maxdownloadbatchsize
	  The maximum size of a single download request from a renter. Each
	download request has multiple round trips of communication that
	exchange money. Larger batch sizes mean fewer round trips, but more
	financial risk for the host - the renter can get a free batch when
	downloading by refusing to provide a signature.


	*/
	Maxdownloadbatchsize *string
	/*Maxduration
	  The maximum duration of a file contract that the host will accept.
	The storage proof window must end before the current height +
	maxduration.


	*/
	Maxduration *string
	/*Maxrevisebatchsize
	  The maximum size of a single batch of file contract revisions. The
	renter can perform DoS attacks on the host by uploading a batch of
	data then refusing to provide a signature to pay for the data. The
	host can reduce this exposure by limiting the batch size. Larger
	batch sizes allow for higher throughput as there is significant
	communication overhead associated with performing a batch upload.


	*/
	Maxrevisebatchsize *string
	/*Mincontractprice
	  The minimum price that the host will demand from a renter when
	forming a contract. Typically this price is to cover transaction
	fees on the file contract revision and storage proof, but can also
	be used if the host has a low amount of collateral. The price is a
	minimum because the host may automatically adjust the price upwards
	in times of high demand.


	*/
	Mincontractprice *string
	/*Mindownloadbandwidthprice
	  The minimum price that the host will demand from a renter when the
	renter is downloading data. If the host is saturated, the host may
	increase the price from the minimum.


	*/
	Mindownloadbandwidthprice *string
	/*Minstorageprice
	  The minimum price that the host will demand when storing data for
	extended periods of time. If the host is low on space, the price of
	storage may be set higher than the minimum.


	*/
	Minstorageprice *string
	/*Minuploadbandwidthprice
	  The minimum price that the host will demand from a renter when the
	renter is uploading data. If the host is saturated, the host may
	increase the price from the minimum.


	*/
	Minuploadbandwidthprice *string
	/*Netaddress
	  The IP address or hostname (including port) that the host should be
	contacted at. If left blank, the host will automatically figure out
	its ip address and use that. If given, the host will use the address
	given.


	*/
	Netaddress *string
	/*Windowsize
	  The storage proof window is the number of blocks that the host has
	to get a storage proof onto the blockchain. The window size is the
	minimum size of window that the host will accept in a file contract.


	*/
	Windowsize *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post host params
func (o *PostHostParams) WithTimeout(timeout time.Duration) *PostHostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post host params
func (o *PostHostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post host params
func (o *PostHostParams) WithContext(ctx context.Context) *PostHostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post host params
func (o *PostHostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post host params
func (o *PostHostParams) WithHTTPClient(client *http.Client) *PostHostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post host params
func (o *PostHostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptingcontracts adds the acceptingcontracts to the post host params
func (o *PostHostParams) WithAcceptingcontracts(acceptingcontracts *bool) *PostHostParams {
	o.SetAcceptingcontracts(acceptingcontracts)
	return o
}

// SetAcceptingcontracts adds the acceptingcontracts to the post host params
func (o *PostHostParams) SetAcceptingcontracts(acceptingcontracts *bool) {
	o.Acceptingcontracts = acceptingcontracts
}

// WithCollateral adds the collateral to the post host params
func (o *PostHostParams) WithCollateral(collateral *string) *PostHostParams {
	o.SetCollateral(collateral)
	return o
}

// SetCollateral adds the collateral to the post host params
func (o *PostHostParams) SetCollateral(collateral *string) {
	o.Collateral = collateral
}

// WithCollateralbudget adds the collateralbudget to the post host params
func (o *PostHostParams) WithCollateralbudget(collateralbudget *string) *PostHostParams {
	o.SetCollateralbudget(collateralbudget)
	return o
}

// SetCollateralbudget adds the collateralbudget to the post host params
func (o *PostHostParams) SetCollateralbudget(collateralbudget *string) {
	o.Collateralbudget = collateralbudget
}

// WithMaxcollateral adds the maxcollateral to the post host params
func (o *PostHostParams) WithMaxcollateral(maxcollateral *string) *PostHostParams {
	o.SetMaxcollateral(maxcollateral)
	return o
}

// SetMaxcollateral adds the maxcollateral to the post host params
func (o *PostHostParams) SetMaxcollateral(maxcollateral *string) {
	o.Maxcollateral = maxcollateral
}

// WithMaxdownloadbatchsize adds the maxdownloadbatchsize to the post host params
func (o *PostHostParams) WithMaxdownloadbatchsize(maxdownloadbatchsize *string) *PostHostParams {
	o.SetMaxdownloadbatchsize(maxdownloadbatchsize)
	return o
}

// SetMaxdownloadbatchsize adds the maxdownloadbatchsize to the post host params
func (o *PostHostParams) SetMaxdownloadbatchsize(maxdownloadbatchsize *string) {
	o.Maxdownloadbatchsize = maxdownloadbatchsize
}

// WithMaxduration adds the maxduration to the post host params
func (o *PostHostParams) WithMaxduration(maxduration *string) *PostHostParams {
	o.SetMaxduration(maxduration)
	return o
}

// SetMaxduration adds the maxduration to the post host params
func (o *PostHostParams) SetMaxduration(maxduration *string) {
	o.Maxduration = maxduration
}

// WithMaxrevisebatchsize adds the maxrevisebatchsize to the post host params
func (o *PostHostParams) WithMaxrevisebatchsize(maxrevisebatchsize *string) *PostHostParams {
	o.SetMaxrevisebatchsize(maxrevisebatchsize)
	return o
}

// SetMaxrevisebatchsize adds the maxrevisebatchsize to the post host params
func (o *PostHostParams) SetMaxrevisebatchsize(maxrevisebatchsize *string) {
	o.Maxrevisebatchsize = maxrevisebatchsize
}

// WithMincontractprice adds the mincontractprice to the post host params
func (o *PostHostParams) WithMincontractprice(mincontractprice *string) *PostHostParams {
	o.SetMincontractprice(mincontractprice)
	return o
}

// SetMincontractprice adds the mincontractprice to the post host params
func (o *PostHostParams) SetMincontractprice(mincontractprice *string) {
	o.Mincontractprice = mincontractprice
}

// WithMindownloadbandwidthprice adds the mindownloadbandwidthprice to the post host params
func (o *PostHostParams) WithMindownloadbandwidthprice(mindownloadbandwidthprice *string) *PostHostParams {
	o.SetMindownloadbandwidthprice(mindownloadbandwidthprice)
	return o
}

// SetMindownloadbandwidthprice adds the mindownloadbandwidthprice to the post host params
func (o *PostHostParams) SetMindownloadbandwidthprice(mindownloadbandwidthprice *string) {
	o.Mindownloadbandwidthprice = mindownloadbandwidthprice
}

// WithMinstorageprice adds the minstorageprice to the post host params
func (o *PostHostParams) WithMinstorageprice(minstorageprice *string) *PostHostParams {
	o.SetMinstorageprice(minstorageprice)
	return o
}

// SetMinstorageprice adds the minstorageprice to the post host params
func (o *PostHostParams) SetMinstorageprice(minstorageprice *string) {
	o.Minstorageprice = minstorageprice
}

// WithMinuploadbandwidthprice adds the minuploadbandwidthprice to the post host params
func (o *PostHostParams) WithMinuploadbandwidthprice(minuploadbandwidthprice *string) *PostHostParams {
	o.SetMinuploadbandwidthprice(minuploadbandwidthprice)
	return o
}

// SetMinuploadbandwidthprice adds the minuploadbandwidthprice to the post host params
func (o *PostHostParams) SetMinuploadbandwidthprice(minuploadbandwidthprice *string) {
	o.Minuploadbandwidthprice = minuploadbandwidthprice
}

// WithNetaddress adds the netaddress to the post host params
func (o *PostHostParams) WithNetaddress(netaddress *string) *PostHostParams {
	o.SetNetaddress(netaddress)
	return o
}

// SetNetaddress adds the netaddress to the post host params
func (o *PostHostParams) SetNetaddress(netaddress *string) {
	o.Netaddress = netaddress
}

// WithWindowsize adds the windowsize to the post host params
func (o *PostHostParams) WithWindowsize(windowsize *string) *PostHostParams {
	o.SetWindowsize(windowsize)
	return o
}

// SetWindowsize adds the windowsize to the post host params
func (o *PostHostParams) SetWindowsize(windowsize *string) {
	o.Windowsize = windowsize
}

// WriteToRequest writes these params to a swagger request
func (o *PostHostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Acceptingcontracts != nil {

		// query param acceptingcontracts
		var qrAcceptingcontracts bool
		if o.Acceptingcontracts != nil {
			qrAcceptingcontracts = *o.Acceptingcontracts
		}
		qAcceptingcontracts := swag.FormatBool(qrAcceptingcontracts)
		if qAcceptingcontracts != "" {
			if err := r.SetQueryParam("acceptingcontracts", qAcceptingcontracts); err != nil {
				return err
			}
		}

	}

	if o.Collateral != nil {

		// query param collateral
		var qrCollateral string
		if o.Collateral != nil {
			qrCollateral = *o.Collateral
		}
		qCollateral := qrCollateral
		if qCollateral != "" {
			if err := r.SetQueryParam("collateral", qCollateral); err != nil {
				return err
			}
		}

	}

	if o.Collateralbudget != nil {

		// query param collateralbudget
		var qrCollateralbudget string
		if o.Collateralbudget != nil {
			qrCollateralbudget = *o.Collateralbudget
		}
		qCollateralbudget := qrCollateralbudget
		if qCollateralbudget != "" {
			if err := r.SetQueryParam("collateralbudget", qCollateralbudget); err != nil {
				return err
			}
		}

	}

	if o.Maxcollateral != nil {

		// query param maxcollateral
		var qrMaxcollateral string
		if o.Maxcollateral != nil {
			qrMaxcollateral = *o.Maxcollateral
		}
		qMaxcollateral := qrMaxcollateral
		if qMaxcollateral != "" {
			if err := r.SetQueryParam("maxcollateral", qMaxcollateral); err != nil {
				return err
			}
		}

	}

	if o.Maxdownloadbatchsize != nil {

		// query param maxdownloadbatchsize
		var qrMaxdownloadbatchsize string
		if o.Maxdownloadbatchsize != nil {
			qrMaxdownloadbatchsize = *o.Maxdownloadbatchsize
		}
		qMaxdownloadbatchsize := qrMaxdownloadbatchsize
		if qMaxdownloadbatchsize != "" {
			if err := r.SetQueryParam("maxdownloadbatchsize", qMaxdownloadbatchsize); err != nil {
				return err
			}
		}

	}

	if o.Maxduration != nil {

		// query param maxduration
		var qrMaxduration string
		if o.Maxduration != nil {
			qrMaxduration = *o.Maxduration
		}
		qMaxduration := qrMaxduration
		if qMaxduration != "" {
			if err := r.SetQueryParam("maxduration", qMaxduration); err != nil {
				return err
			}
		}

	}

	if o.Maxrevisebatchsize != nil {

		// query param maxrevisebatchsize
		var qrMaxrevisebatchsize string
		if o.Maxrevisebatchsize != nil {
			qrMaxrevisebatchsize = *o.Maxrevisebatchsize
		}
		qMaxrevisebatchsize := qrMaxrevisebatchsize
		if qMaxrevisebatchsize != "" {
			if err := r.SetQueryParam("maxrevisebatchsize", qMaxrevisebatchsize); err != nil {
				return err
			}
		}

	}

	if o.Mincontractprice != nil {

		// query param mincontractprice
		var qrMincontractprice string
		if o.Mincontractprice != nil {
			qrMincontractprice = *o.Mincontractprice
		}
		qMincontractprice := qrMincontractprice
		if qMincontractprice != "" {
			if err := r.SetQueryParam("mincontractprice", qMincontractprice); err != nil {
				return err
			}
		}

	}

	if o.Mindownloadbandwidthprice != nil {

		// query param mindownloadbandwidthprice
		var qrMindownloadbandwidthprice string
		if o.Mindownloadbandwidthprice != nil {
			qrMindownloadbandwidthprice = *o.Mindownloadbandwidthprice
		}
		qMindownloadbandwidthprice := qrMindownloadbandwidthprice
		if qMindownloadbandwidthprice != "" {
			if err := r.SetQueryParam("mindownloadbandwidthprice", qMindownloadbandwidthprice); err != nil {
				return err
			}
		}

	}

	if o.Minstorageprice != nil {

		// query param minstorageprice
		var qrMinstorageprice string
		if o.Minstorageprice != nil {
			qrMinstorageprice = *o.Minstorageprice
		}
		qMinstorageprice := qrMinstorageprice
		if qMinstorageprice != "" {
			if err := r.SetQueryParam("minstorageprice", qMinstorageprice); err != nil {
				return err
			}
		}

	}

	if o.Minuploadbandwidthprice != nil {

		// query param minuploadbandwidthprice
		var qrMinuploadbandwidthprice string
		if o.Minuploadbandwidthprice != nil {
			qrMinuploadbandwidthprice = *o.Minuploadbandwidthprice
		}
		qMinuploadbandwidthprice := qrMinuploadbandwidthprice
		if qMinuploadbandwidthprice != "" {
			if err := r.SetQueryParam("minuploadbandwidthprice", qMinuploadbandwidthprice); err != nil {
				return err
			}
		}

	}

	if o.Netaddress != nil {

		// query param netaddress
		var qrNetaddress string
		if o.Netaddress != nil {
			qrNetaddress = *o.Netaddress
		}
		qNetaddress := qrNetaddress
		if qNetaddress != "" {
			if err := r.SetQueryParam("netaddress", qNetaddress); err != nil {
				return err
			}
		}

	}

	if o.Windowsize != nil {

		// query param windowsize
		var qrWindowsize string
		if o.Windowsize != nil {
			qrWindowsize = *o.Windowsize
		}
		qWindowsize := qrWindowsize
		if qWindowsize != "" {
			if err := r.SetQueryParam("windowsize", qWindowsize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}