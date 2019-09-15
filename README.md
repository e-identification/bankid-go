# BankID library

A library for providing BankID services as a RP (Relying party).
Supports the latest v5 features.

[![Build Status](https://travis-ci.org/NicklasWallgren/bankid.svg?branch=master)](https://travis-ci.org/NicklasWallgren/bankid)
[![Go Report Card](https://goreportcard.com/badge/github.com/stretchr/testify)](https://goreportcard.com/report/github.com/NicklasWallgren/bankid)
[![GoDoc](https://godoc.org/github.com/NicklasWallgren/bankid?status.svg)](https://godoc.org/github.com/NicklasWallgren/bankid) 

Check out the API Documentation http://godoc.org/github.com/NicklasWallgren/bankid

# Installation
The library can be installed through `go get` 
```bash
go get github.com/NicklasWallgren/bankid
```

# Supported versions
We support the two major Go versions, which are 1.12 and 1.13 at the moment.

# Features
- Supports all v5 features

# Examples 

## Initiate sign request
```go
configuration := configuration.NewConfiguration(
    &configuration.TestEnvironment,
    configuration.GetResourcePath("certificates/test.crt"),
    configuration.GetResourcePath("certificates/test.key"))

bankId := bankid.New(configuration)
payload := bankid.SignPayload{PersonalNumber: "<INSERT PERSONAL NUMBER>", EndUserIp: "192.168.1.1", UserVisibleData: "Test"}
response, err := bankId.Sign(&payload)

if err != nil {
    fmt.Println(err)

    return
}

if response.IsSuccess() {
    response.Collect()
} 
```

## Unit tests
```bash
go test -v -race $(go list ./... | grep -v vendor)
```

## TODO
 - Add unit tests
 - Add validator translator

## Contributing
  - Fork it!
  - Create your feature branch: `git checkout -b my-new-feature`
  - Commit your changes: `git commit -am 'Useful information about your new features'`
  - Push to the branch: `git push origin my-new-feature`
  - Submit a pull request

## Contributors
  - [Nicklas Wallgren](https://github.com/NicklasWallgren)
  - [All Contributors][link-contributors]
