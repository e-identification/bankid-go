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
	decode(subject Response, response *http.Response, bankID *BankID) (*Response, error)
}

type jsonDecoder struct{}

func newJSONDecoder() Decoder {
	return &jsonDecoder{}
}

func (j jsonDecoder) decode(subject Response, response *http.Response, bankID *BankID) (*Response, error) {
	if !isValidHTTPResponse(response.StatusCode, httpStatusCodes) {
		return nil, fmt.Errorf("invalid http response. Http Code: %d. Body: %s", response.StatusCode, readCloserToString(response.Body))
	}

	if isHttpStatusCodeWithinRange(response.StatusCode, successRange) {
		decoded, err := subject.Decode(response.Body, bankID)
		if err != nil {
			return nil, fmt.Errorf("unable to decode response %w", err)
		}

		return &decoded, nil
	}

	if isHttpStatusCodeWithinRange(response.StatusCode, errorRange) {
		return nil, j.decodeError(response, bankID)
	}

	decoded, err := subject.Decode(response.Body, bankID)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response %w", err)
	}

	return &decoded, nil
}

func (j jsonDecoder) decodeError(response *http.Response, bankID *BankID) error {
	errorResponse := ErrorResponse{}

	_, err := errorResponse.Decode(response.Body, bankID)
	if err != nil {
		return err
	}

	return &errorResponse
}

type statusCodeRange struct {
	start int
	end   int
}
