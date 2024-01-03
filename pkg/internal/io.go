package internal

import (
	"bytes"
	"io"
)

// TryReadCloserToString reads the content of an io.ReadCloser into a string.
func TryReadCloserToString(readCloser io.ReadCloser) string {
	buf := new(bytes.Buffer)

	if _, err := buf.ReadFrom(readCloser); err != nil {
		return ""
	}

	return buf.String()
}
