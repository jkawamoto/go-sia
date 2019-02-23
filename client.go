/*
 * client.go
 *
 * Copyright (c) 2019 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package sia

//noinspection SpellCheckingInspection
import (
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/jkawamoto/go-sia/client"
)

const (
	EnvSiaHost = "SIA_HOST"
)

func NewClient() (cli *client.Sia) {

	host := os.Getenv(EnvSiaHost)
	if host == "" {
		host = client.DefaultHost
	}
	cli = client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     host,
		BasePath: client.DefaultBasePath,
		Schemes:  client.DefaultSchemes,
	})

	switch transport := cli.Transport.(type) {
	case *httptransport.Runtime:
		transport.Transport = newTransporter(transport.Transport)
		cli.SetTransport(transport)
	}

	return

}
