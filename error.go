package bankid

import (
	"errors"
	"fmt"
)

var payloadValidationError = errors.New("payload validation error")

func PayloadValidationError() error {
	return fmt.Errorf("%w", payloadValidationError)
}
