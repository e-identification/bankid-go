package bankid

import (
	"fmt"
	"net/http"
)

var (
	successRange = statusCodeRange{200, 300 - 1}
	errorRange   = statusCodeRange{400, 600 - 1}
)

type Decoder interface {
	decode(subject Response, response *http.Response, bankId *BankId) (*Response, error)
}

type jsonDecoder struct{}

func newJsonDecoder() Decoder {
	return &jsonDecoder{}
}

func (j jsonDecoder) decode(subject Response, response *http.Response, bankId *BankId) (*Response, error) {
	if !isValidHttpResponse(response.StatusCode, httpStatusCodes) {
		return nil, fmt.Errorf("invalid http response. Http Code: %d. Body: %s", response.StatusCode, response.Body)
	}

	if isHttpStatusCodeWithinRange(response.StatusCode, successRange) {
		decoded, err := subject.Decode(response.Body, bankId)

		return &decoded, err
	}

	if isHttpStatusCodeWithinRange(response.StatusCode, errorRange) {
		return nil, j.decodeError(response, bankId)
	}

	decoded, err := subject.Decode(response.Body, bankId)

	return &decoded, err
}

func (j jsonDecoder) decodeError(response *http.Response, bankId *BankId) error {
	errorResponse := ErrorResponse{}

	_, err := errorResponse.Decode(response.Body, bankId)

	if err != nil {
		return err
	}

	return &errorResponse
}

type statusCodeRange struct {
	start int
	end   int
}
