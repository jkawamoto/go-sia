/*
 * transporter_test.go
 *
 * Copyright (c) 2019-2021 Junpei Kawamoto
 *
 * This software is released under the MIT License.
 *
 * http://opensource.org/licenses/mit-license.php
 */

package sia

import (
	"net/http"
	"testing"
)

type mockRoundTripper struct {
	Request  *http.Request
	Response *http.Response
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	m.Request = req
	return m.Response, nil
}

func TestSiaTransporter(t *testing.T) {
	sampleResponse := &http.Response{
		Status:     "Test",
		StatusCode: 999,
	}

	mock := &mockRoundTripper{
		Response: sampleResponse,
	}

	req, err := http.NewRequest(http.MethodTrace, "https://example.com/test", nil)
	if err != nil {
		t.Fatalf("cannot create a sample request: %v", err)
	}

	tran := newTransporter(mock)

	res, err := tran.RoundTrip(req)
	if err != nil {
		t.Errorf("RoundTrip failed: %v", err)
	}
	if res != sampleResponse {
		t.Errorf("expected %v but got %v", sampleResponse, res)
	}
	if mock.Request != req {
		t.Errorf("expected %v but got %v", req, mock.Request)
	}
	if agent := mock.Request.Header.Get("User-Agent"); agent != "Sia-Agent" {
		t.Errorf("expected user agent is %v but got %v", "Sia-Agent", agent)
	}
}
