package bankid

import (
	"bytes"
	"io"
)

func tryReadCloserToString(readCloser io.ReadCloser) string {
	buf := new(bytes.Buffer)

	if _, err := buf.ReadFrom(readCloser); err != nil {
		return ""
	}

	return buf.String()
}
