package bankid

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"path"
	"runtime"
)

func isValidHTTPResponse(statusCode int, httpStatusCodes []int) bool {
	for _, validStatusCode := range httpStatusCodes {
		if statusCode == validStatusCode {
			return true
		}
	}
	return false
}

func isHTTPStatusCodeWithinRange(statusCode int, statusCodeRange statusCodeRange) bool {
	return statusCode >= statusCodeRange.start && statusCode <= statusCodeRange.end
}

func getResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "./resource"), nil
}

func getResourcePath(path string) (directory string) {
	dir, err := getResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}

func readCloserToString(readCloser io.ReadCloser) string {
	buf := new(bytes.Buffer)

	_, _ = buf.ReadFrom(readCloser)

	return buf.String()
}
