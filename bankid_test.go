package bankid

import (
	"context"
	"crypto/tls"
	"github.com/NicklasWallgren/bankid/configuration"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
	"net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"unsafe"
)

const (
	signOkResponse = `{
		"orderRef":"131daac9-16c6-4618-beb0-365768f37288",
		"autoStartToken":"7c40b5c9-fa74-49cf-b98c-bfe651f9a7c6" 
	}`
	collectOkResponse = `{
		"orderRef": "131daac9-16c6-4618-beb0-365768f37288",
  		"status": "complete",
  		"completionData": {
			"user": {
      			"personalNumber": "190000000000",
      			"name": "Karl Karlsson",
      			"givenName": "Karl",
      			"surname": "Karlsson"
    		},
    		"device": {
      			"ipAddress": "192.168.0.1"
			},
    		"cert": {
      			"notBefore": "1502983274000",
      			"notAfter": "1563549674000"
    		},
    		"signature": "<base64-encoded data>",
    		"ocspResponse": "<base64-encoded data>"
		}
	}
	`
)

func TestSignRequestWithValidPayload(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.Write([]byte(signOkResponse))
	})

	bankId, teardown := createMockedClient(&handler)

	defer teardown()

	payload := &SignPayload{PersonalNumber: "123456789123", EndUserIp: "192.168.1.1", UserVisibleData: "Test", Requirement: &Requirement{CardReader: ""}}

	bankId.Sign(context.Background(), payload)
}

func TestSignRequestWithInvalidPayload(t *testing.T) {
	bankId := New(&configuration.Configuration{})

	payload := &SignPayload{PersonalNumber: "INVALID-PERSONAL-NUMBER", EndUserIp: "192.168.1.1", UserVisibleData: "Test", Requirement: &Requirement{CardReader: ""}}
	_, err := bankId.Sign(context.Background(), payload)

	validationErrors := err.(validator.ValidationErrors)
	fieldError := validationErrors[0]
	assert.Equal(t, "PersonalNumber", fieldError.Field())
	assert.Equal(t, "numeric", fieldError.Tag())
}

func TestCollectRequestWithValidPayload(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		w.Write([]byte(collectOkResponse))
	})

	bankId, teardown := createMockedClient(&handler)

	defer teardown()

	payload := &CollectPayload{OrderRef: ""}

	bankId.Collect(context.Background(), payload)
}

func createMockedClient(handler *http.HandlerFunc) (*BankId, func()) {
	configuration := configuration.New(&configuration.TestEnvironment, getResourcePath("certificates/test.crt"), getResourcePath("certificates/test.key"))

	bankId := New(configuration)

	httpClient, teardown := testingHTTPClient(handler)
	client, _ := newClient(configuration, withHttpClient(httpClient))
	var test *Client = &client
	rs := reflect.ValueOf(bankId).Elem()
	rf := rs.Field(2)
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()

	field := reflect.New(reflect.TypeOf(test))
	field.Elem().Set(reflect.ValueOf(test))

	rf.Set(field.Elem())

	return bankId, teardown
}

func testingHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}
