// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new host API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for host API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHost fetches status information about the host
*/
func (a *Client) GetHost(params *GetHostParams, authInfo runtime.ClientAuthInfoWriter) (*GetHostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHost",
		Method:             "GET",
		PathPattern:        "/host",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetHost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetHostStorage gets a list of folders tracked by the host's storage manager.
*/
func (a *Client) GetHostStorage(params *GetHostStorageParams, authInfo runtime.ClientAuthInfoWriter) (*GetHostStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHostStorageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetHostStorage",
		Method:             "GET",
		PathPattern:        "/host/storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHostStorageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHostStorageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetHostStorageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHost configures hosting parameters. All parameters are optional; unspecified parameters will be left unchanged.
*/
func (a *Client) PostHost(params *PostHostParams, authInfo runtime.ClientAuthInfoWriter) (*PostHostNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHost",
		Method:             "POST",
		PathPattern:        "/host",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHostDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHostAnnounce Announce the host to the network as a source of storage. Generally only needs to be called once.
*/
func (a *Client) PostHostAnnounce(params *PostHostAnnounceParams, authInfo runtime.ClientAuthInfoWriter) (*PostHostAnnounceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostAnnounceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHostAnnounce",
		Method:             "POST",
		PathPattern:        "/host/announce",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostAnnounceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostAnnounceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHostAnnounceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHostStorageFoldersAdd adds a storage folder to the manager. The manager may not check that there is enough space available on-disk to support as much storage as requested
*/
func (a *Client) PostHostStorageFoldersAdd(params *PostHostStorageFoldersAddParams, authInfo runtime.ClientAuthInfoWriter) (*PostHostStorageFoldersAddNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostStorageFoldersAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHostStorageFoldersAdd",
		Method:             "POST",
		PathPattern:        "/host/storage/folders/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostStorageFoldersAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostStorageFoldersAddNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHostStorageFoldersAddDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHostStorageFoldersRemove Remove a storage folder from the manager. All storage on the folder will be moved to other storage folders, meaning that no data will be lost. If the manager is unable to save data, an error will be returned and the operation will be stopped.

*/
func (a *Client) PostHostStorageFoldersRemove(params *PostHostStorageFoldersRemoveParams, authInfo runtime.ClientAuthInfoWriter) (*PostHostStorageFoldersRemoveNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostStorageFoldersRemoveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHostStorageFoldersRemove",
		Method:             "POST",
		PathPattern:        "/host/storage/folders/remove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostStorageFoldersRemoveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostStorageFoldersRemoveNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHostStorageFoldersRemoveDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHostStorageFoldersResize grows or shrink a storage folder in the manager. The manager may not check that there is enough space on-disk to support growing the storage folder, but should gracefully handle running out of space unexpectedly. When shrinking a storage folder, any data in the folder that needs to be moved will be placed into other storage folders, meaning that no data will be lost. If the manager is unable to migrate the data, an error will be returned and the operation will be stopped.

*/
func (a *Client) PostHostStorageFoldersResize(params *PostHostStorageFoldersResizeParams, authInfo runtime.ClientAuthInfoWriter) (*PostHostStorageFoldersResizeNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostStorageFoldersResizeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHostStorageFoldersResize",
		Method:             "POST",
		PathPattern:        "/host/storage/folders/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostStorageFoldersResizeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostStorageFoldersResizeNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHostStorageFoldersResizeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostHostStorageSectorsDeleteMerkleroot deletes a sector, meaning that the manager will be unable to upload that sector and be unable to provide a storage proof on that sector.
This endpoint is for removing the data entirely, and will remove instances of the sector appearing at all heights.
The primary purpose is to comply with legal requests to remove data.

*/
func (a *Client) PostHostStorageSectorsDeleteMerkleroot(params *PostHostStorageSectorsDeleteMerklerootParams, authInfo runtime.ClientAuthInfoWriter) (*PostHostStorageSectorsDeleteMerklerootNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostHostStorageSectorsDeleteMerklerootParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostHostStorageSectorsDeleteMerkleroot",
		Method:             "POST",
		PathPattern:        "/host/storage/sectors/delete/{merkleroot}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostHostStorageSectorsDeleteMerklerootReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostHostStorageSectorsDeleteMerklerootNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostHostStorageSectorsDeleteMerklerootDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
