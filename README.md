# BankID library

[![Build Status](https://github.com/e-identification/bankid-go/workflows/Test/badge.svg)](https://github.com/e-identification/bankid-go/actions?query=workflow%3ATest)
[![Reviewdog](https://github.com/e-identification/bankid-go/workflows/reviewdog/badge.svg)](https://github.com/e-identification/bankid-go/actions?query=workflow%3Areviewdog)
[![Go Report Card](https://goreportcard.com/badge/github.com/e-identification/bankid-go)](https://goreportcard.com/report/github.com/e-identification/bankid-go)
[![GoDoc](https://godoc.org/github.com/e-identification/bankid-go?status.svg)](https://godoc.org/github.com/e-identification/bankid-go)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/564889b413574c82a13b2ff82cade53e)](https://app.codacy.com/gh/e-identification/bankid-go/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)

A library for providing BankID services as an RP (Relying party).  
Supports the latest v6 features.

To learn how to use the library, please refer to the [documentation](https://godoc.org/github.com/e-identification/bankid-go). There are some [examples](./examples) that may be useful as well.

# Installation
The library can be installed through `go get` 
```bash
go get github.com/e-identification/bankid-go
```

# Supported versions
We support the two major Go versions, which are 1.21 and 1.22 at the moment.

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
