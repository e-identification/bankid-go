package bankid

import (
	"errors"
	"fmt"
	"path"
	"runtime"
)

func isValidHttpResponse(statusCode int, httpStatusCodes []int) bool {
	for _, validStatusCode := range httpStatusCodes {
		if statusCode == validStatusCode {
			return true
		}
	}
	return false
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
