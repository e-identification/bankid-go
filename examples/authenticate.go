package main

import (
	"fmt"
	"github.com/NicklasWallgren/bankid"
	"github.com/NicklasWallgren/bankid/configuration"
)

func main() {
	configuration := configuration.New(
		&configuration.TestEnvironment,
		configuration.GetResourcePath("certificates/test.crt"),
		configuration.GetResourcePath("certificates/test.key"))

	bankId := bankid.New(configuration)

	payload := bankid.AuthenticationPayload{PersonalNumber: "<INSERT PERSONAL NUMBER>", EndUserIp: "192.168.1.1"}

	response, err := bankId.Authenticate(&payload)

	if err != nil {
		if response := bankid.UnwrapErrorResponse(err); response != nil {
			fmt.Printf("%s - %s \n", response.Details, response.ErrorCode)
		}

		fmt.Printf("%#v", err)
		return
	}

	fmt.Println(response.Collect())
}
