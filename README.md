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

## Initiate authenticate request
```php
$client = new Client(new Config(<CERTFICATE>));

$authenticationResponse = $client->authenticate(new AuthenticationPayload(<PERSONAL NUMBER>, <IP ADDRESS>));

if (!$authenticationResponse->isSuccess()) {
    var_dump($authenticationResponse->getErrorCode(), $authenticationResponse->getDetails());

    return;
}

$collectResponse = $authenticationResponse->collect(); 
```

# Certificates
The web service API can only be accessed by a RP that has a valid SSL client certificate. The RP certificate is obtained from the
bank that the RP has purchased the BankID service from.

## Generate PEM certificate
```bash
openssl pkcs12 -in <filename>.pfx -out <cert>.pem -nodes
```

## Unit tests
```bash
go test -v -race $(go list ./... | grep -v vendor)
```

## TODO
Add unit tests
Add validator translator
Add all Recommended User Messages

## Contributing
  - Fork it!
  - Create your feature branch: `git checkout -b my-new-feature`
  - Commit your changes: `git commit -am 'Useful information about your new features'`
  - Push to the branch: `git push origin my-new-feature`
  - Submit a pull request

## Contributors
  - [Nicklas Wallgren](https://github.com/NicklasWallgren)
  - [All Contributors][link-contributors]
