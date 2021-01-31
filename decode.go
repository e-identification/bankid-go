package bankid

import (
	"fmt"
	"net/http"
)

var (
	successRange = statusCodeRange{200, 300 - 1}
	errorRange   = statusCodeRange{400, 600 - 1}
)

type decoder interface {
	decode(subject Response, response *http.Response, bankID *BankID) (Response, error)
}

type jsonDecoder struct{}

func newJSONDecoder() decoder {
	return &jsonDecoder{}
}

func (j jsonDecoder) decode(subject Response, response *http.Response, bankID *BankID) (Response, error) {
	if !isValidHTTPResponse(response.StatusCode, expectedHTTPStatusCodes) {
		return nil, fmt.Errorf("invalid http response. Http Code: %d. Body: %s", response.StatusCode, tryReadCloserToString(response.Body))
	}

	if isHTTPStatusCodeWithinRange(response.StatusCode, successRange) {
		decoded, err := subject.Decode(response.Body, bankID)
		if err != nil {
			return nil, fmt.Errorf("unable to decode response %w", err)
		}

		return decoded, nil
	}

	if isHTTPStatusCodeWithinRange(response.StatusCode, errorRange) {
		return nil, j.decodeError(response, bankID)
	}

	return nil, fmt.Errorf("unable to decode response %s", tryReadCloserToString(response.Body))
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
