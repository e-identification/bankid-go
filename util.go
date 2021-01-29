package bankid

import (
	"bytes"
	"io"
)

func readCloserToString(readCloser io.ReadCloser) string {
	buf := new(bytes.Buffer)

	_, _ = buf.ReadFrom(readCloser)

	return buf.String()
}
