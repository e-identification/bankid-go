package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/NicklasWallgren/bankid"
	"github.com/NicklasWallgren/bankid/configuration"
)

func main() {
	configuration := configuration.New(
		configuration.TestEnvironment,
		&configuration.Pkcs12{Content: loadPkcs12(getResourcePath("certificates/test.p12")), Password: "qwerty123"},
	)

	bankID := bankid.New(configuration)

	payload := bankid.AuthenticationPayload{PersonalNumber: "<INSERT PERSONAL NUMBER>", EndUserIP: "192.168.1.1"}

	response, err := bankID.Authenticate(context.Background(), &payload)
	if err != nil {
		if response := unwrapAsErrorResponse(err); response != nil {
			fmt.Printf("%s - %s \n", response.Details, response.ErrorCode)
		}

		fmt.Printf("%#v", err)
		return
	}

	fmt.Println(response.Collect(context.Background()))
}

func loadPkcs12(path string) []byte {
	cert, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return cert
}

func unwrapAsErrorResponse(err error) *bankid.ErrorResponse {
	var response bankid.ErrorResponse

	if errors.Is(err, response) && errors.As(err, &response) {
		return &response
	}

	return nil
}

// getResourceDirectoryPath returns the full path to the resource directory.
func getResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", fmt.Errorf("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "../resource"), nil
}

// getResourcePath returns the full path to the resource.
func getResourcePath(path string) (directory string) {
	dir, err := getResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}
