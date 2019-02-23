/*
 * transporter.go
 *
 * Copyright (c) 2019 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package sia

import (
	"net/http"
	"strings"
)

// transporter is a transporter which adds authentication information  to each request before transporting it.
type transporter struct {
	http.RoundTripper
}

// newTransporter creates a transporter which wraps a given transporter.
func newTransporter(transport http.RoundTripper) *transporter {

	return &transporter{
		RoundTripper: transport,
	}

}

// RoundTrip adds User-Agent header to the given request and processes it.
func (t *transporter) RoundTrip(req *http.Request) (*http.Response, error) {

	// Unescape slashes
	req.URL.RawPath = strings.Replace(req.URL.RawPath, "%2F", "/", -1)

	// Append user agent header
	req.Header.Add("User-Agent", "Sia-Agent")

	return t.RoundTripper.RoundTrip(req)

}
