package http

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/e-identification/bankid-go/pkg/response"
)

func TestDecodeUsingInvalidHttpResponse(t *testing.T) {
	decoder := newJSONDecoder()
	mockHTTPResponse := &http.Response{
		StatusCode: http.StatusGatewayTimeout, Body: io.NopCloser(strings.NewReader("output")),
	}

	_, err := decoder.decode(nil, mockHTTPResponse)
	if err == nil {
		t.Fail()
		return
	}

	if err.Error() != "invalid http Response. Http Code: 504. Body: output" {
		t.Fail()
	}
}

func TestDecodeUsingHttpResponseWithContentTypeOtherThanApplicationJson(t *testing.T) {
	decoder := newJSONDecoder()
	mockHTTPResponse := &http.Response{
		StatusCode: http.StatusForbidden, Body: io.NopCloser(strings.NewReader("output")),
	}

	_, err := decoder.decode(nil, mockHTTPResponse)
	if err.Error() != "unable to decode error response: output" {
		t.Fail()
	}
}

func TestDecodeUsingInvalidSuccessResponseBody(t *testing.T) {
	decoder := newJSONDecoder()
	mockHTTPResponse := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader("[]"))}

	request := Request{Response: &response.AuthenticateResponse{}}

	_, err := decoder.decode(&request, mockHTTPResponse)
	if err == nil {
		t.Fail()
		return
	}

	if !strings.Contains(err.Error(), "unable to decode response json") {
		t.Fail()
	}
}
