package bankid

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestDecodeUsingInvalidHttpResponse(t *testing.T) {
	decoder := newJSONDecoder()
	mockHTTPResponse := &http.Response{StatusCode: 504, Body: ioutil.NopCloser(strings.NewReader("output"))}

	_, err := decoder.decode(nil, mockHTTPResponse, nil)

	if err == nil {
		t.Fail()
		return
	}

	if err.Error() != "invalid http response. Http Code: 504. Body: output" {
		t.Fail()
	}
}

func TestDecodeUsingInvalidSuccessResponseBody(t *testing.T) {
	decoder := newJSONDecoder()
	mockHTTPResponse := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader("[]"))}

	_, err := decoder.decode(&AuthenticateResponse{}, mockHTTPResponse, nil)

	if err == nil {
		t.Fail()
		return
	}

	if !strings.Contains(err.Error(), "unable to decode response json") {
		t.Fail()
	}
}
