// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new wallet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for wallet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetWallet returns basic information about the wallet,  such as whether the wallet is locked or unlocked.
*/
func (a *Client) GetWallet(params *GetWalletParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWallet",
		Method:             "GET",
		PathPattern:        "/wallet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletOK), nil

}

/*
GetWalletAddress gets a new address from the wallet generated by the primary seed. An error will be returned if the wallet is locked.
*/
func (a *Client) GetWalletAddress(params *GetWalletAddressParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletAddress",
		Method:             "GET",
		PathPattern:        "/wallet/address",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletAddressOK), nil

}

/*
GetWalletAddresses fetches the list of addresses from the wallet.  If the wallet has not been created or unlocked, no addresses will be returned. After the wallet is unlocked, this call will continue to return its addresses even after the wallet is locked again.
*/
func (a *Client) GetWalletAddresses(params *GetWalletAddressesParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletAddressesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletAddressesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletAddresses",
		Method:             "GET",
		PathPattern:        "/wallet/addresses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletAddressesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletAddressesOK), nil

}

/*
GetWalletBackup creates a backup of the wallet settings file. Though this can easily be done manually, the settings file is often in an unknown or difficult to find location. The /wallet/backup call can spare users the trouble of needing to find their wallet file. The destination file is overwritten if it already exists.
*/
func (a *Client) GetWalletBackup(params *GetWalletBackupParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletBackupNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletBackup",
		Method:             "GET",
		PathPattern:        "/wallet/backup",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletBackupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletBackupNoContent), nil

}

/*
GetWalletSeeds returns a list of seeds in use by the wallet.  The primary seed is the only seed that gets used to generate new addresses. This call is unavailable when the wallet is locked.
A seed is an encoded version of a 128 bit random seed. The output is 15 words chosen from a small dictionary as indicated by the input. The most common choice for the dictionary is going to be 'english'. The underlying seed is the same no matter what dictionary is used for the encoding. The encoding also contains a small checksum of the seed,  to help catch simple mistakes when copying. The library entropy-mnemonics is used when encoding.
*/
func (a *Client) GetWalletSeeds(params *GetWalletSeedsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletSeedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletSeedsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletSeeds",
		Method:             "GET",
		PathPattern:        "/wallet/seeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletSeedsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletSeedsOK), nil

}

/*
GetWalletTransactionID gets the transaction associated with a specific transaction id.
*/
func (a *Client) GetWalletTransactionID(params *GetWalletTransactionIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletTransactionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletTransactionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletTransactionID",
		Method:             "GET",
		PathPattern:        "/wallet/transaction/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletTransactionIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletTransactionIDOK), nil

}

/*
GetWalletTransactions returns a list of transactions related to the wallet.
*/
func (a *Client) GetWalletTransactions(params *GetWalletTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletTransactions",
		Method:             "GET",
		PathPattern:        "/wallet/transactions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletTransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletTransactionsOK), nil

}

/*
GetWalletTransactionsAddr returns all of the transactions related to a specific address.
*/
func (a *Client) GetWalletTransactionsAddr(params *GetWalletTransactionsAddrParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletTransactionsAddrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletTransactionsAddrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletTransactionsAddr",
		Method:             "GET",
		PathPattern:        "/wallet/transactions/{addr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletTransactionsAddrReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletTransactionsAddrOK), nil

}

/*
GetWalletVerifyAddressAddr takes the address specified by addr and returns a JSON response indicating if the address is valid.
*/
func (a *Client) GetWalletVerifyAddressAddr(params *GetWalletVerifyAddressAddrParams, authInfo runtime.ClientAuthInfoWriter) (*GetWalletVerifyAddressAddrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWalletVerifyAddressAddrParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetWalletVerifyAddressAddr",
		Method:             "GET",
		PathPattern:        "/wallet/verify/address/{addr}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWalletVerifyAddressAddrReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWalletVerifyAddressAddrOK), nil

}

/*
PostWallet033x loads a v0.3.3.x wallet into the current wallet, harvesting all of the secret keys. All spendable addresses in the loaded wallet will become spendable from the current wallet. An error will be returned if the given encryptionpassword is incorrect.
*/
func (a *Client) PostWallet033x(params *PostWallet033xParams, authInfo runtime.ClientAuthInfoWriter) (*PostWallet033xNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWallet033xParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWallet033x",
		Method:             "POST",
		PathPattern:        "/wallet/033x",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWallet033xReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWallet033xNoContent), nil

}

/*
PostWalletChangepassword changes the wallet's encryption password.
*/
func (a *Client) PostWalletChangepassword(params *PostWalletChangepasswordParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletChangepasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletChangepasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletChangepassword",
		Method:             "POST",
		PathPattern:        "/wallet/changepassword",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletChangepasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletChangepasswordNoContent), nil

}

/*
PostWalletInit initializes the wallet. After the wallet has been initialized once, it does not need to be initialized again, and future calls to /wallet/init will return an error, unless the force flag is set. The encryption password is provided by the api call. If the password is blank, then the password will be set to the same as the seed.
*/
func (a *Client) PostWalletInit(params *PostWalletInitParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletInitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletInitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletInit",
		Method:             "POST",
		PathPattern:        "/wallet/init",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletInitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletInitOK), nil

}

/*
PostWalletInitSeed initializes the wallet using a preexisting seed.  After the wallet has been initialized once, it does not need to be initialized again, and future calls to /wallet/init/seed will return an error unless the force flag is set. The encryption password is provided by the api call. If the password is blank, then the password will be set to the same as the seed. Note that loading a preexisting seed requires scanning the blockchain to determine how many keys have been generated from the seed. For this reason, /wallet/init/seed can only be called if the blockchain is synced.
*/
func (a *Client) PostWalletInitSeed(params *PostWalletInitSeedParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletInitSeedNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletInitSeedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletInitSeed",
		Method:             "POST",
		PathPattern:        "/wallet/init/seed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletInitSeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletInitSeedNoContent), nil

}

/*
PostWalletLock locks the wallet, wiping all secret keys. After being locked, the keys are encrypted. Queries for the seed, to send siafunds, and related queries become unavailable. Queries concerning transaction history and balance are still available.
*/
func (a *Client) PostWalletLock(params *PostWalletLockParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletLockNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletLockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletLock",
		Method:             "POST",
		PathPattern:        "/wallet/lock",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletLockReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletLockNoContent), nil

}

/*
PostWalletSeed gives the wallet a seed to track when looking for incoming transactions. The wallet will be able to spend outputs related to addresses created by the seed. The seed is added as an auxiliary seed, and does not replace the primary seed. Only the primary seed will be used for generating new addresses.
*/
func (a *Client) PostWalletSeed(params *PostWalletSeedParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletSeedNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletSeedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletSeed",
		Method:             "POST",
		PathPattern:        "/wallet/seed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletSeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletSeedNoContent), nil

}

/*
PostWalletSiacoins Send siacoins to an address or set of addresses. The outputs are arbitrarily selected from addresses in the wallet. If 'outputs' is supplied, 'amount' and 'destination' must be empty. The number of outputs should not exceed 400; this may result in a transaction too large to fit in the transaction pool.
*/
func (a *Client) PostWalletSiacoins(params *PostWalletSiacoinsParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletSiacoinsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletSiacoinsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletSiacoins",
		Method:             "POST",
		PathPattern:        "/wallet/siacoins",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletSiacoinsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletSiacoinsOK), nil

}

/*
PostWalletSiafunds sends siafunds to an address.  The outputs are arbitrarily selected from addresses in the wallet. Any siacoins available in the siafunds being sent (as well as the siacoins  available in any siafunds that end up in a refund address) will become  available to the wallet as siacoins after 144 confirmations.  To access all of the siacoins in the siacoin claim balance, send all of  the siafunds to an address in your control (this will give you all the  siacoins, while still letting you control the siafunds).
*/
func (a *Client) PostWalletSiafunds(params *PostWalletSiafundsParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletSiafundsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletSiafundsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletSiafunds",
		Method:             "POST",
		PathPattern:        "/wallet/siafunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletSiafundsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletSiafundsOK), nil

}

/*
PostWalletSiagkey Load a key into the wallet that was generated by siag.  Most siafunds are currently in addresses created by siag.
*/
func (a *Client) PostWalletSiagkey(params *PostWalletSiagkeyParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletSiagkeyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletSiagkeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletSiagkey",
		Method:             "POST",
		PathPattern:        "/wallet/siagkey",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletSiagkeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletSiagkeyNoContent), nil

}

/*
PostWalletSweepSeed Scan the blockchain for outputs belonging to a seed and send them to an address owned by the wallet.
*/
func (a *Client) PostWalletSweepSeed(params *PostWalletSweepSeedParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletSweepSeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletSweepSeedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletSweepSeed",
		Method:             "POST",
		PathPattern:        "/wallet/sweep/seed",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletSweepSeedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletSweepSeedOK), nil

}

/*
PostWalletUnlock unlocks the wallet. The wallet is capable of knowing whether the correct password was provided.
*/
func (a *Client) PostWalletUnlock(params *PostWalletUnlockParams, authInfo runtime.ClientAuthInfoWriter) (*PostWalletUnlockNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWalletUnlockParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostWalletUnlock",
		Method:             "POST",
		PathPattern:        "/wallet/unlock",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWalletUnlockReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWalletUnlockNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
