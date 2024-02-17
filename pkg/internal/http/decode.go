package http

import (
	"fmt"
	"net/http"

	"github.com/e-identification/bankid-go/pkg/internal"
)

type statusCodeRange struct {
	start int
	end   int
}

var (
	successRange = statusCodeRange{200, 300 - 1}
	errorRange   = statusCodeRange{400, 600 - 1}
)

type decoder interface {
	decode(request *Request, response *http.Response) (Response, error)
}

type jsonDecoder struct{}

func newJSONDecoder() decoder {
	return &jsonDecoder{}
}

func (j jsonDecoder) decode(request *Request, response *http.Response) (Response, error) {
	if !isValidHTTPResponse(response.StatusCode, expectedHTTPStatusCodes) {
		return nil, fmt.Errorf("invalid http Response. Http Code: %d. Body: %s",
			response.StatusCode, internal.TryReadCloserToString(response.Body))
	}

	if isHTTPStatusCodeWithinRange(response.StatusCode, successRange) {
		err := Decode(response.Body, request.Response)
		if err != nil {
			return nil, fmt.Errorf("unable to decode response %w", err)
		}

		request.Response.OnDecode()

		return request.Response, nil
	}

	if isHTTPStatusCodeWithinRange(response.StatusCode, errorRange) {
		return nil, j.decodeError(request, response)
	}

	return nil, fmt.Errorf("unable to decode Response %s", internal.TryReadCloserToString(response.Body))
}

func (j jsonDecoder) decodeError(request *Request, response *http.Response) error {
	err := Decode(response.Body, request.ErrorResponse)
	if err != nil {
		return err
	}

	return request.ErrorResponse
}

func isHTTPStatusCodeWithinRange(statusCode int, statusCodeRange statusCodeRange) bool {
	return statusCode >= statusCodeRange.start && statusCode <= statusCodeRange.end
}

func isValidHTTPResponse(statusCode int, httpStatusCodes []int) bool {
	for _, validStatusCode := range httpStatusCodes {
		if statusCode == validStatusCode {
			return true
		}
	}

	return false
}
