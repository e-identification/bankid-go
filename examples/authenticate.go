package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/NicklasWallgren/bankid/v2/pkg"
	"github.com/NicklasWallgren/bankid/v2/pkg/configuration"
	"github.com/NicklasWallgren/bankid/v2/pkg/payload"
)

func main() {
	configuration := configuration.NewConfiguration(
		configuration.TestEnvironment,
		&configuration.Pkcs12{Content: loadPkcs12(getResourcePath("certificates/test.p12")), Password: "qwerty123"},
	)

	bankID, err := pkg.NewBankIDClient(configuration)

	payload := payload.AuthenticationPayload{EndUserIP: "192.168.1.1", UserVisibleData: "blabla", Requirement: &payload.Requirement{PersonalNumber: "201912312392"}}

	httpResponse, err := bankID.Authenticate(context.Background(), &payload)
	if err != nil {
		var apiError *pkg.APIError

		if errors.As(err, &apiError) {
			fmt.Printf("%s - %s \n", apiError.Details, apiError.ErrorCode)
			return
		}

		fmt.Printf("%#v", err)
		return
	}

	fmt.Println(httpResponse)

	// fmt.Println(httpResponse.Collect(context.Background()))
}

func loadPkcs12(path string) []byte {
	cert, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return cert
}

// getResourceDirectoryPath returns the full path to the resource directory.
func getResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", fmt.Errorf("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "../pkg/resource"), nil
}

// getResourcePath returns the full path to the resource.
func getResourcePath(path string) (directory string) {
	dir, err := getResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}
