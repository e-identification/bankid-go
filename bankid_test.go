package bankid

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/NicklasWallgren/bankid/configuration"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestAuthenticate(t *testing.T) {
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/sign_response.json"))
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
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/sign_response.json"))
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

	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		t.Error("Invalid error type")
	}

	fieldError := validationErrors[0]
	assert.Equal(t, "PersonalNumber", fieldError.Field())
	assert.Equal(t, "numeric", fieldError.Tag())
}

func TestCollect(t *testing.T) {
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/collect_response.json"))
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
	bankID, teardown := testBankID(stringToResponseHandler(t, "{}"))
	defer teardown()

	payload := &CancelPayload{OrderRef: ""}

	response, err := bankID.Cancel(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("Got nil response")
	}
}

func TestQRCodeContent(t *testing.T) {
	bankID, teardown := testBankID(stringToResponseHandler(t, "{}"))
	defer teardown()

	qrCodeContent, err := bankID.QRCodeContent("67df3917-fa0d-44e5-b327-edcc928297f8", "d28db9a7-4cde-429e-a983-359be676944c", 0)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "bankid.67df3917-fa0d-44e5-b327-edcc928297f8.0.dc69358e712458a66a7525beef148ae8526b1c71610eff2c16cdffb4cdac9bf8", qrCodeContent)
}

// Returns a bankID whose requests will always return
// a response configured by the handler.
func testBankID(handler http.HandlerFunc) (*BankID, func()) {
	configuration := configuration.New(configuration.TestEnvironment,
		&configuration.Pkcs12{Content: loadFile(getResourcePath("certificates/test.p12")), Password: "qwerty123"})

	bankID := New(configuration)

	httpClient, teardown := testHTTPClient(handler)

	client, _ := newClient(configuration, withHTTPClient(httpClient))
	bankID.client = client

	return bankID, teardown
}

func testHTTPClient(handler http.Handler) (*http.Client, func()) {
	s := httptest.NewTLSServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String()) // nolint:wrapcheck
			},
			// #nosec G402
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return cli, s.Close
}

func fileToResponseHandler(t *testing.T, filename string) http.HandlerFunc {
	file, err := os.Open(filename) // #nosec G304
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		// nolint:errcheck
		// #nosec G104
		io.Copy(w, file)
		// nolint:errcheck
		// #nosec G104
		file.Close()
	}
}

func stringToResponseHandler(t *testing.T, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		// nolint:errcheck
		// #nosec G104
		io.WriteString(w, body)
	}
}

func loadFile(path string) []byte {
	// #nosec G304
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return fileContent
}

func getResourcePath(path string) (directory string) {
	dir, err := getResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}

func getResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "./resource"), nil
}
