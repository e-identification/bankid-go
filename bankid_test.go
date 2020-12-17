package bankid

import (
	"context"
	"crypto/tls"
	"github.com/NicklasWallgren/bankid/configuration"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	bankID, teardown := testClient(fileToResponseHandler(t, "resource/test_data/sign_response.json"))
	defer teardown()

	payload := &AuthenticationPayload{PersonalNumber: "123456789123", EndUserIP: "192.168.1.1"}

	response, err := bankID.Authenticate(context.Background(), payload)

	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("Got nil response")
	}
}

func TestSign(t *testing.T) {
	bankID, teardown := testClient(fileToResponseHandler(t, "resource/test_data/sign_response.json"))
	defer teardown()

	payload := &SignPayload{PersonalNumber: "123456789123", EndUserIP: "192.168.1.1", UserVisibleData: "Test", Requirement: &Requirement{CardReader: ""}}

	response, err := bankID.Sign(context.Background(), payload)

	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("Got nil response")
	}
	if response.AutoStartToken != "7c40b5c9-fa74-49cf-b98c-bfe651f9a7c6" {
		t.Error("Got wrong auto start token")
	}
}

func TestSignWithInvalidPayload(t *testing.T) {
	bankID := New(&configuration.Configuration{})

	payload := &SignPayload{PersonalNumber: "INVALID-PERSONAL-NUMBER", EndUserIP: "192.168.1.1", UserVisibleData: "Test", Requirement: &Requirement{CardReader: ""}}
	_, err := bankID.Sign(context.Background(), payload)

	validationErrors := err.(validator.ValidationErrors)
	fieldError := validationErrors[0]
	assert.Equal(t, "PersonalNumber", fieldError.Field())
	assert.Equal(t, "numeric", fieldError.Tag())
}

func TestCollect(t *testing.T) {
	bankID, teardown := testClient(fileToResponseHandler(t, "resource/test_data/collect_response.json"))
	defer teardown()

	payload := &CollectPayload{OrderRef: ""}

	response, err := bankID.Collect(context.Background(), payload)

	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("Got nil response")
	}
}

func TestCancel(t *testing.T) {

}

func testClient(handler http.HandlerFunc) (*BankID, func()) {
	configuration := configuration.New(&configuration.TestEnvironment, getResourcePath("certificates/test.crt"), getResourcePath("certificates/test.key"))

	bankID := New(configuration)

	httpClient, teardown := testHTTPClient(handler)

	client, _ := newClient(configuration, withHTTPClient(httpClient))
	bankID.client = &client

	return bankID, teardown
}

func testHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			// #nosec G402
			TLSClientConfig: &tls.Config{},
		},
	}

	return cli, s.Close
}

func fileToResponseHandler(t *testing.T, filename string) http.HandlerFunc {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		io.Copy(w, file) // nolint:errcheck
		file.Close()
	}
}
