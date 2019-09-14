package bankid

import (
	"fmt"
	http "net/http"
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

	decoded, err := subject.Decode(response.Body, bankId)

	return &decoded, err
}
