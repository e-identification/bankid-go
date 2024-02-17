# BankID library

[![Build Status](https://github.com/e-identification/bankid/workflows/Test/badge.svg)](https://github.com/e-identification/bankid/actions?query=workflow%3ATest)
[![Reviewdog](https://github.com/e-identification/bankid/workflows/reviewdog/badge.svg)](https://github.com/e-identification/bankid/actions?query=workflow%3Areviewdog)
[![Go Report Card](https://goreportcard.com/badge/github.com/e-identification/bankid)](https://goreportcard.com/report/github.com/e-identification/bankid)
[![GoDoc](https://godoc.org/github.com/e-identification/bankid?status.svg)](https://godoc.org/github.com/e-identification/bankid)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/cabd5fbbcde543ec959fb4a3581600ed)](https://app.codacy.com/gh/NicklasWallgren/bankid?utm_source=github.com&utm_medium=referral&utm_content=NicklasWallgren/bankid&utm_campaign=Badge_Grade)

A library for providing BankID services as an RP (Relying party).  
Supports the latest v6 features.

To learn how to use the library, please refer to the [documentation](https://godoc.org/github.com/e-identification/bankid). There are some [examples](./examples) that may be useful as well.

# Installation
The library can be installed through `go get` 
```bash
go get github.com/e-identification/bankid
```

# Supported versions
We support the two major Go versions, which are 1.20 and 1.21 at the moment.

# Features
- Supports all v6.0 features

# SDK
```go
// Creates new BankIDClient instance
NewBankIDClient(configuration *configuration.Configuration) (*BankIDClient)

// Initiates an authentication order 
(b BankIDClient) Authenticate(context context.Context, payload *AuthenticationPayload) (*AuthenticateResponse, error)

// Initiates a phone authentication order 
(b BankIDClient) PhoneAuthenticate(context context.Context, payload *PhoneAuthenticationPayload) (*PhoneAuthenticateResponse, error)

// Initiates a sign order
(b BankIDClient) Sign(context context.Context, payload *SignPayload) (*SignResponse, error)

// Initiates a phone sign order
(b BankIDClient) PhoneSign(context context.Context, payload *PhoneSignPayload) (*PhoneSignResponse, error)

// Collects the result of a sign or auth order using the orderRef as reference
(b BankIDClient) Collect(context context.Context, payload *CollectPayload) (*CollectResponse, error)

// Cancels an ongoing sign or auth order
(b BankIDClient) Cancel(context context.Context, payload *CancelPayload) (*CancelResponse, error)
```

## Unit tests
```bash
go test -v -race $(go list ./...)
```

### Code Guide

We use GitHub Actions to make sure the codebase is consistent (`golangci-lint run`) and continuously tested (`go test -v -race $(go list ./...)`). We try to keep comments at a maximum of 120 characters of length and code at 120.

## Contributing

If you find any problems or have suggestions about this library, please submit an issue. Moreover, any pull request, code review and feedback are welcome.

## License

[MIT](./LICENSE)
