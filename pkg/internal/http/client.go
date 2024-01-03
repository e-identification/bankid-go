package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/NicklasWallgren/bankid/v2/pkg/configuration"
)

// The known API endpoints status codes.
var expectedHTTPStatusCodes = []int{200, 400, 401, 403, 404, 408, 415, 500, 503}

// Client is the interface implemented by types that can invoke the BankID REST API.
type Client interface {
	// Call is responsible for making the HTTP call against BankID REST API
	Call(context context.Context, request *Request) (Response, error)
}

type client struct {
	client        *http.Client
	configuration *configuration.Configuration
	encoder       encoder
	decoder       decoder
}

// Option definition.
type Option func(*client)

// NewClient returns a new instance of 'NewClient'.
func NewClient(configuration *configuration.Configuration, options ...Option) (Client, error) {
	clientCfg, err := NewTLSClientConfig(configuration)
	if err != nil {
		return nil, fmt.Errorf("error reading and/or parsing the certification files. %w", err)
	}

	netClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: clientCfg,
		},
	}

	instance := &client{
		client: &netClient, configuration: configuration,
		encoder: newJSONEncoder(), decoder: newJSONDecoder(),
	}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(instance)
	}

	return instance, nil
}

// WithHTTPClient Function to create Option func to set net/http client.
func WithHTTPClient(target *http.Client) Option {
	return func(subject *client) {
		subject.client = target
	}
}

// Call is responsible for making the HTTP call against BankID REST API.
func (c client) Call(ctx context.Context, request *Request) (Response, error) {
	encoded, err := c.encoder.encode(request.Payload)
	if err != nil {
		return nil, fmt.Errorf("unable to encode payload. %w", err)
	}

	req, err := c.newRequest(ctx, c.urlFrom(request), strings.NewReader(string(encoded)))
	if err != nil {
		return nil, fmt.Errorf("unable to create request. %w", err)
	}

	resp, err := c.request(req)
	if err != nil {
		return nil, fmt.Errorf("unable to execute the http request . %w", err)
	}

	defer resp.Body.Close() // nolint:errcheck

	return c.decoder.decode(request, resp) // nolint:wrapcheck
}

// newRequest creates and prepares a instance of http Request.
func (c client) newRequest(context context.Context, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(context, http.MethodPost, url, body)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err // nolint:wrapcheck
	}

	return req, nil
}

func (c client) urlFrom(request *Request) string {
	return c.configuration.Environment.BaseURL + "/" + request.URI
}

func (c client) request(request *http.Request) (*http.Response, error) {
	return c.client.Do(request) // nolint:wrapcheck
}
