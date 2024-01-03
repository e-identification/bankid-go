package internal

import (
	"encoding/base64"
	"strconv"

	"gopkg.in/go-playground/validator.v9"
)

const (
	base64Length = "base64Length"
)

// NewValidator returns a new instance of 'Validator' prepared with custom validator methods.
func NewValidator() (*validator.Validate, error) {
	instance := validator.New()

	if err := instance.RegisterValidation(base64Length, validateBase64Length); err != nil {
		return nil, err // nolint:wrapcheck
	}

	return instance, nil
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
