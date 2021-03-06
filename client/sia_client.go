// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/jkawamoto/go-sia/client/consensus"
	"github.com/jkawamoto/go-sia/client/daemon"
	"github.com/jkawamoto/go-sia/client/gateway"
	"github.com/jkawamoto/go-sia/client/host"
	"github.com/jkawamoto/go-sia/client/host_d_b"
	"github.com/jkawamoto/go-sia/client/miner"
	"github.com/jkawamoto/go-sia/client/renter"
	"github.com/jkawamoto/go-sia/client/wallet"
)

// Default sia HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost:9980"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new sia HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Sia {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new sia HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Sia {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new sia client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Sia {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Sia)
	cli.Transport = transport
	cli.Consensus = consensus.New(transport, formats)
	cli.Daemon = daemon.New(transport, formats)
	cli.Gateway = gateway.New(transport, formats)
	cli.Host = host.New(transport, formats)
	cli.Hostdb = host_d_b.New(transport, formats)
	cli.Miner = miner.New(transport, formats)
	cli.Renter = renter.New(transport, formats)
	cli.Wallet = wallet.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Sia is a client for sia
type Sia struct {
	Consensus consensus.ClientService

	Daemon daemon.ClientService

	Gateway gateway.ClientService

	Host host.ClientService

	Hostdb host_d_b.ClientService

	Miner miner.ClientService

	Renter renter.ClientService

	Wallet wallet.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Sia) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Consensus.SetTransport(transport)
	c.Daemon.SetTransport(transport)
	c.Gateway.SetTransport(transport)
	c.Host.SetTransport(transport)
	c.Hostdb.SetTransport(transport)
	c.Miner.SetTransport(transport)
	c.Renter.SetTransport(transport)
	c.Wallet.SetTransport(transport)
}
