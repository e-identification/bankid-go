package bankid

import (
	"context"
	"fmt"
	"github.com/NicklasWallgren/bankid/configuration"
	"io"
	"net/http"
	"strings"
)

var (
	// The known API endpoints status codes
	httpStatusCodes = []int{200, 400, 401, 403, 404, 408, 415, 500, 503}
)

// Client is the interface implemented by types that can invoke the BankID REST API
type Client interface {
	// call is responsible for making the HTTP call against BankId REST API
	call(request Request, context context.Context, bankId *BankId) (*Response, error)
}

type client struct {
	client        *http.Client
	configuration *configuration.Configuration
	encoder       Encoder
	decoder       Decoder
}

// Option definition
type Option func(*client)

// newClient returns a new instance of 'newClient'
func newClient(configuration *configuration.Configuration, options ...Option) (Client, error) {
	clientCfg, err := newTLSClientConfig(configuration)

	if err != nil {
		return nil, fmt.Errorf("error reading and/or parsing the certification files. Cause: %s", err)
	}

	netClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: clientCfg,
		},
	}

	instance := &client{client: &netClient, configuration: configuration, encoder: newJsonEncoder(), decoder: newJsonDecoder()}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(instance)
	}

	return instance, nil
}

// Function to create Option func to set net/http client
func withHttpClient(target *http.Client) Option {
	return func(subject *client) {
		subject.client = target
	}
}

// call is responsible for making the HTTP call against BankId REST API
func (c client) call(request Request, context context.Context, bankId *BankId) (*Response, error) {
	encoded, err := c.encoder.encode(request.Payload())

	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(c.composeUrl(request), strings.NewReader(string(encoded)))
	if err != nil {
		return nil, err
	}

	resp, err := c.request(req.WithContext(context))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return c.decoder.decode(request.Response(), resp, bankId)
}

// newRequest creates and prepares a instance of http request
func (c client) newRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, body)

	req.Header.Add("Content-Type", "application/json")

	return req, err
}

func (c client) composeUrl(request Request) string {
	return c.configuration.Environment.BaseUrl + "/" + request.Uri()
}

func (c client) request(request *http.Request) (*http.Response, error) {
	return c.client.Do(request)
}
