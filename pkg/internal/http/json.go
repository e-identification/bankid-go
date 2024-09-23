package http

import (
	"encoding/json"
	"io"
)

// Decode decodes a json string into the target type.
func Decode(readCloser io.ReadCloser, target any) error {
	decoder := json.NewDecoder(readCloser)

	return decoder.Decode(&target) // nolint:wrapcheck
}
