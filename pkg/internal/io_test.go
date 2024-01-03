package internal

import (
	"io"
	"strings"
	"testing"
)

func TestTryReadCloserToString(t *testing.T) {
	stringReader := strings.NewReader("output")
	stringReadCloser := io.NopCloser(stringReader)

	result := TryReadCloserToString(stringReadCloser)

	if result != "output" {
		t.Fail()
	}
}
