package bankid

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/stimtech/go-bankid/configuration"
	"golang.org/x/crypto/pkcs12"
)

func newTLSClientConfig(configuration *configuration.Configuration) (*tls.Config, error) {
	caPool, err := createCertPool(configuration.Environment.Certificate)
	if err != nil {
		return nil, err
	}

	rpCert, err := createCertLeaf(configuration)
	if err != nil {
		return nil, err
	}

	// #nosec G402
	clientCfg := &tls.Config{
		Certificates: []tls.Certificate{*rpCert},
		ClientCAs:    caPool,
		RootCAs:      caPool,
	}

	return clientCfg, nil
}

func createCertPool(base64EncodedCertificate string) (*x509.CertPool, error) {
	certificate, err := base64.StdEncoding.DecodeString(base64EncodedCertificate)
	if err != nil {
		return nil, fmt.Errorf("could not decode the certificate. %w", err)
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(certificate) {
		return nil, fmt.Errorf("could not append CA Certificate to pool. Invalid base64EncodedCertificate")
	}

	return caPool, nil
}

func createCertLeaf(configuration *configuration.Configuration) (*tls.Certificate, error) {
	key, leaf, err := pkcs12.Decode(configuration.Pkcs12.Content, configuration.Pkcs12.Password)
	if err != nil {
		return nil, fmt.Errorf("unable to load pkcs12 %w", err)
	}

	cert := &tls.Certificate{
		Certificate: [][]byte{leaf.Raw},
		PrivateKey:  key.(crypto.PrivateKey),
		Leaf:        leaf,
	}

	return cert, nil
}
