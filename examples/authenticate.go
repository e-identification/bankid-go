package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/NicklasWallgren/bankid"
	"github.com/NicklasWallgren/bankid/configuration"
)

func main() {
	configuration := configuration.New(
		&configuration.TestEnvironment,
		configuration.GetResourcePath("certificates/test.crt"),
		configuration.GetResourcePath("certificates/test.key"))

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

func unwrapAsErrorResponse(err error) *bankid.ErrorResponse {
	var response bankid.ErrorResponse

	if errors.Is(err, response) && errors.As(err, &response) {
		return &response
	}

	return nil
}
