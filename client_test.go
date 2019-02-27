/*
 * client_test.go
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
	"testing"

	"github.com/jkawamoto/go-sia/client"

	httptransport "github.com/go-openapi/runtime/client"
)

func TestNewClient(t *testing.T) {

	cli := NewClient("")

	runtime, ok := cli.Transport.(*httptransport.Runtime)
	if !ok {
		t.Fatal("doesn't match the type of runtime")
	}

	_, ok = runtime.Transport.(*transporter)
	if !ok {
		t.Errorf("doesn't match the type of transport")
	}
	if runtime.Host != client.DefaultHost {
		t.Errorf("%v was set to the host but %v is expected", runtime.Host, client.DefaultHost)
	}

}

func TestNewClientWithHost(t *testing.T) {

	// Set the environment variable
	host := "192.168.100.1:8080"

	// Create a client
	cli := NewClient(host)

	runtime, ok := cli.Transport.(*httptransport.Runtime)
	if !ok {
		t.Fatal("doesn't match the type of runtime")
	}
	if runtime.Host != host {
		t.Errorf("%v was set to the host but %v is expected", runtime.Host, host)
	}

}
