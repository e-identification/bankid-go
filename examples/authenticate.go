package main

import (
	"fmt"
	"github.com/NicklasWallgren/bankid"
	"github.com/NicklasWallgren/bankid/configuration"
)

func main() {
	configuration := configuration.NewConfiguration(
		&configuration.TestEnvironment,
		configuration.GetResourcePath("certificates/test.crt"),
		configuration.GetResourcePath("certificates/test.key"))

	bankId := bankid.New(configuration)

	payload := bankid.AuthenticationPayload{PersonalNumber: "<INSERT PERSONAL NUMBER>", EndUserIp: "192.168.1.1"}

	response, err := bankId.Authenticate(&payload)

	if err != nil {
		fmt.Println(err)

		return
	}

	if response.IsSuccess() {
		response.Collect()
	}

	fmt.Println(response.ErrorCode)
	fmt.Println(response.OrderRef)
	fmt.Println(response.IsSuccess())
}
