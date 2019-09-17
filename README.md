# Sia client library for Go
[![GoDoc](https://godoc.org/github.com/jkawamoto/go-sia?status.svg)](http://godoc.org/github.com/jkawamoto/go-sia)
[![Build Status](https://travis-ci.org/jkawamoto/go-sia.svg?branch=master)](https://travis-ci.org/jkawamoto/go-sia)
[![Release](https://img.shields.io/badge/release-0.5.4-brightgreen.svg)](https://github.com/jkawamoto/go-sia/releases/tag/v0.5.4)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)


```go
import (
    "context"
    
    "github.com/jkawamoto/go-sia"
    "github.com/jkawamoto/go-sia/client/renter"
    httptransport "github.com/go-openapi/runtime/client"
)

var (
    DefaultDataPieces   int64 = 10
    DefaultParityPieces int64 = 20
)

client := sia.NewClient("")
client.Renter.PostRenterUploadSiapath(
    renter.NewPostRenterUploadSiapathParamsWithContext(context.Background()).
        WithSiapath("sia-path").
        WithSource("path/to/local/file").
        WithDatapieces(&DefaultDataPieces).
        WithParitypieces(&DefaultParityPieces),
    httptransport.BasicAuth("", "Sia Password"))
```