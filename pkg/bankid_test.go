package pkg

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/e-identification/bankid/pkg/configuration"
	bankIdHttp "github.com/e-identification/bankid/pkg/internal/http"
	"github.com/e-identification/bankid/pkg/payload"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"
)

func TestAuthenticate(t *testing.T) {
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/sign_response.json"))
	defer teardown()

	requestPayload := &payload.AuthenticationPayload{
		EndUserIP: "192.168.1.1", Requirement: &payload.Requirement{PersonalNumber: "123456789123"},
	}

	response, err := bankID.Authenticate(context.Background(), requestPayload)
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("Got nil response")
	}
}

func TestPhoneAuthentication(t *testing.T) {
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/phone_sign_response.json"))
	defer teardown()

	requestPayload := &payload.PhoneAuthenticationPayload{
		PersonalNumber: "123456789123", CallInitiator: "RP", UserVisibleData: "Test",
	}

	response, err := bankID.PhoneAuthenticate(context.Background(), requestPayload)
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("Got nil response")
	}

	if response.OrderRef != "131daac9-16c6-4618-beb0-365768f37288" {
		t.Error("Got wrong auto start token")
	}
}

func TestSign(t *testing.T) {
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/sign_response.json"))
	defer teardown()

	requestPayload := &payload.SignPayload{
		EndUserIP:       "192.168.1.1",
		UserVisibleData: "Test",
		Requirement:     &payload.Requirement{CardReader: "", PersonalNumber: "123456789123"},
	}

	response, err := bankID.Sign(context.Background(), requestPayload)
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

func TestPhoneSign(t *testing.T) {
	bankID, teardown := testBankID(fileToResponseHandler(t, "resource/test_data/phone_sign_response.json"))
	defer teardown()

	requestPayload := &payload.PhoneSignPayload{
		PersonalNumber: "123456789123", CallInitiator: "RP", UserVisibleData: "Test",
	}

	response, err := bankID.PhoneSign(context.Background(), requestPayload)
	if err != nil {
		t.Fatal(err)
	}

	if response == nil {
		t.Fatal("Got nil response")
	}

	if response.OrderRef != "131daac9-16c6-4618-beb0-365768f37288" {
		t.Error("Got wrong auto start token")
	}
}

func TestSignWithInvalidPayload(t *testing.T) {
	bankID, _ := NewBankIDClient(configuration.NewConfiguration(configuration.TestEnvironment,
		&configuration.Pkcs12{Content: loadFile(getResourcePath("certificates/test.p12")), Password: "qwerty123"}))

	requestPayload := &payload.SignPayload{
		EndUserIP:       "192.168.1.1",
		UserVisibleData: "Test",
		Requirement:     &payload.Requirement{CardReader: "", PersonalNumber: "INVALID-PERSONAL-NUMBER"},
	}
	_, err := bankID.Sign(context.Background(), requestPayload)

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

	requestPayload := &payload.CollectPayload{OrderRef: ""}

	response, err := bankID.Collect(context.Background(), requestPayload)
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

	requestPayload := &payload.CancelPayload{OrderRef: ""}

	response, err := bankID.Cancel(context.Background(), requestPayload)
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

	qrCodeContent, err := bankID.QRCodeContent(
		"67df3917-fa0d-44e5-b327-edcc928297f8",
		"d28db9a7-4cde-429e-a983-359be676944c",
		0)
	if err != nil {
		t.Fatal(err)
	}
	// nolint: lll
	assert.Equal(t, "bankid.67df3917-fa0d-44e5-b327-edcc928297f8.0.dc69358e712458a66a7525beef148ae8526b1c71610eff2c16cdffb4cdac9bf8", qrCodeContent)
}

// Returns a bankID whose requests will always return
// a response configured by the handler.
func testBankID(handler http.HandlerFunc) (*BankIDClient, func()) {
	clientConfiguration := configuration.NewConfiguration(configuration.TestEnvironment,
		&configuration.Pkcs12{Content: loadFile(getResourcePath("certificates/test.p12")), Password: "qwerty123"})

	bankID, _ := NewBankIDClient(clientConfiguration)

	httpClient, teardown := testHTTPClient(handler)

	client, _ := bankIdHttp.NewClient(clientConfiguration, bankIdHttp.WithHTTPClient(httpClient))
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
				InsecureSkipVerify: true, // nolint:gosec
			},
		},
	}

	return cli, s.Close
}

func fileToResponseHandler(t *testing.T, filename string) http.HandlerFunc {
	t.Helper()

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
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		// nolint:errcheck
		// #nosec G104
		io.WriteString(w, body)
	}
}

func loadFile(path string) []byte {
	// #nosec G304
	fileContent, err := os.ReadFile(path)
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
