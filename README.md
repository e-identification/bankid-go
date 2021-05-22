# BankID library

A library for providing BankID services as a RP (Relying party).
Supports the latest v5 features.

[![Build Status](https://github.com/NicklasWallgren/bankid/workflows/Test/badge.svg)](https://github.com/NicklasWallgren/bankid/actions?query=workflow%3ATest)
[![Reviewdog](https://github.com/NicklasWallgren/bankid/workflows/reviewdog/badge.svg)](https://github.com/NicklasWallgren/bankid/actions?query=workflow%3Areviewdog)
[![Go Report Card](https://goreportcard.com/badge/github.com/NicklasWallgren/bankid)](https://goreportcard.com/report/github.com/NicklasWallgren/bankid)
[![GoDoc](https://godoc.org/github.com/NicklasWallgren/bankid?status.svg)](https://godoc.org/github.com/NicklasWallgren/bankid)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cabd5fbbcde543ec959fb4a3581600ed)](https://app.codacy.com/gh/NicklasWallgren/bankid?utm_source=github.com&utm_medium=referral&utm_content=NicklasWallgren/bankid&utm_campaign=Badge_Grade)

Check out the API Documentation http://godoc.org/github.com/NicklasWallgren/bankid

# Installation
The library can be installed through `go get` 
```bash
go get github.com/NicklasWallgren/bankid
```

# Supported versions
We support the two major Go versions, which are 1.14 and 1.15 at the moment.

# Features
- Supports all v5.1 features

# Examples 

## Initiate sign request
```go
import (
    "context"
    "fmt"
    "io/ioutil"
    "github.com/NicklasWallgren/bankid"
    "github.com/NicklasWallgren/bankid/configuration"
)

certificate, err := ioutil.ReadFile("path/to/environment.p12")
if err != nil {
    panic(err)
}

configuration := configuration.New(
    configuration.TestEnvironment,
    &configuration.Pkcs12{Content: certificate), Password: "p12 password"},
)

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

### Code Guide

We use GitHub Actions to make sure the codebase is consistent (`golangci-lint run`) and continuously tested (`go test -v -race $(go list ./... | grep -v vendor)`). We try to keep comments at a maximum of 120 characters of length and code at 120.


## Contributing

If you find any problems or have suggestions about this library, please submit an issue. Moreover, any pull request, code review and feedback are welcome.

## Contributors
  - [Nicklas Wallgren](https://github.com/NicklasWallgren)
  - [All Contributors][link-contributors]

[link-contributors]: ../../contributors

## License

[MIT](./LICENSE)
