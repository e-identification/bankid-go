package bankid

import (
	"encoding/base64"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

const (
	base64Length = "base64Length"
)

// newValidator returns a new instance of 'Validator' prepared with custom validator methods.
func newValidator() *validator.Validate {
	instance := validator.New()

	if err := instance.RegisterValidation(base64Length, validateBase64Length); err != nil {
		panic(err)
	}

	return instance
}

// validateBase64Length validates the length of a encoded string.
func validateBase64Length(fl validator.FieldLevel) bool {
	length, err := strconv.Atoi(fl.Param())

	if err != nil {
		return false
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(fl.Field().String()))

	return len(encoded) <= length
}
