# BankID library

A library for providing BankID services as a RP (Relying party).
Supports the latest v5 features.

[![Build Status](https://github.com/NicklasWallgren/bankid/workflows/Test/badge.svg)](https://github.com/NicklasWallgren/bankid/actions?query=workflow%3ATest)
[![Reviewdog](https://github.com/NicklasWallgren/bankid/workflows/reviewdog/badge.svg)](https://github.com/NicklasWallgren/bankid/actions?query=workflow%3Areviewdog)
[![Go Report Card](https://goreportcard.com/badge/github.com/stretchr/testify)](https://goreportcard.com/report/github.com/NicklasWallgren/bankid)
[![GoDoc](https://godoc.org/github.com/NicklasWallgren/bankid?status.svg)](https://godoc.org/github.com/NicklasWallgren/bankid) 

Check out the API Documentation http://godoc.org/github.com/NicklasWallgren/bankid

# Installation
The library can be installed through `go get` 
```bash
go get github.com/NicklasWallgren/bankid
```

# Supported versions
We support the two major Go versions, which are 1.14 and 1.15 at the moment.

# Features
- Supports all v5 features

# Examples 

## Initiate sign request
```go
import (
    "fmt"
    "github.com/NicklasWallgren/bankid"
    "github.com/NicklasWallgren/bankid/configuration"
)

configuration := configuration.New(
    &configuration.TestEnvironment,
    configuration.GetResourcePath("certificates/test.crt"),
    configuration.GetResourcePath("certificates/test.key"))

bankId := bankid.New(configuration)

payload := bankid.SignPayload{PersonalNumber: "<INSERT PERSONAL NUMBER>", EndUserIp: "192.168.1.1", UserVisibleData: "Test"}

response, err := bankId.Sign(&payload)

if err != nil {
    if response := bankid.UnwrapErrorResponse(err); response != nil {
        fmt.Printf("%s - %s \n", response.Details, response.ErrorCode)
    }

    fmt.Printf("%#v", err)
    return
}

fmt.Println(response.Collect())
```

## Unit tests
```bash
go test -v -race $(go list ./... | grep -v vendor)
```

## Contributing
  - Fork it!
  - Create your feature branch: `git checkout -b my-new-feature`
  - Commit your changes: `git commit -am 'Useful information about your new features'`
  - Push to the branch: `git push origin my-new-feature`
  - Submit a pull request

## Contributors
  - [Nicklas Wallgren](https://github.com/NicklasWallgren)
  - [All Contributors][link-contributors]

[link-contributors]: ../../contributors