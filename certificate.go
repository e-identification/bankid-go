package bankid

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/NicklasWallgren/bankid/configuration"
	"io/ioutil"
)

func newTLSClientConfig(configuration *configuration.Configuration) (*tls.Config, error) {
	caPool, err := createCertPool(configuration.Environment.CertificationFilePath)

	if err != nil {
		return nil, err
	}

	rpCert, err := createCertLeaf(configuration)

	if err != nil {
		return nil, err
	}

	clientCfg := &tls.Config{
		Certificates: []tls.Certificate{*rpCert},
		ClientCAs:    caPool,
		RootCAs:      caPool,
	}

	return clientCfg, nil
}

func createCertPool(certificatePath string) (*x509.CertPool, error) {
	ca, err := ioutil.ReadFile(certificatePath)

	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(ca) {
		return nil, fmt.Errorf("could not append CA Certificate to pool. Invalid certificate")
	}

	return caPool, nil
}

func createCertLeaf(configuration *configuration.Configuration) (*tls.Certificate, error) {
	rpCert, err := tls.LoadX509KeyPair(configuration.CertFile, configuration.KeyFile)

	if err != nil {
		return nil, err
	}

	return &rpCert, nil
}
