package bankid

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestTryReadCloserToString(t *testing.T) {
	stringReader := strings.NewReader("output")
	stringReadCloser := ioutil.NopCloser(stringReader)

	result := tryReadCloserToString(stringReadCloser)

	if result != "output" {
		t.Fail()
	}
}
