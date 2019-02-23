// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new daemon API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for daemon API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDaemonConstants Returns the set of constants in use.
*/
func (a *Client) GetDaemonConstants(params *GetDaemonConstantsParams, authInfo runtime.ClientAuthInfoWriter) (*GetDaemonConstantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDaemonConstantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDaemonConstants",
		Method:             "GET",
		PathPattern:        "/daemon/constants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDaemonConstantsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDaemonConstantsOK), nil

}

/*
GetDaemonStop cleanly shuts down the daemon. May take a few seconds.
*/
func (a *Client) GetDaemonStop(params *GetDaemonStopParams, authInfo runtime.ClientAuthInfoWriter) (*GetDaemonStopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDaemonStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDaemonStop",
		Method:             "GET",
		PathPattern:        "/daemon/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDaemonStopReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDaemonStopOK), nil

}

/*
GetDaemonVersion returns the version of the Sia daemon currently running.
*/
func (a *Client) GetDaemonVersion(params *GetDaemonVersionParams, authInfo runtime.ClientAuthInfoWriter) (*GetDaemonVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDaemonVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDaemonVersion",
		Method:             "GET",
		PathPattern:        "/daemon/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDaemonVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDaemonVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
