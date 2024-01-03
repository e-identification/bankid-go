package pkg

import "fmt"

// A ValidationError is returned when the payload is found to be invalid.
type ValidationError struct {
	Field         string
	RejectedValue any
	wrapped       error
}

// NewValidationError initialize a new ValidationError.
func NewValidationError(field string, rejectedValue any, wrapped error) *ValidationError {
	return &ValidationError{Field: field, RejectedValue: rejectedValue, wrapped: wrapped}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation error. %v", v.wrapped)
}

// Unwrap unwraps the wrapped error.
func (v ValidationError) Unwrap() error {
	return v.wrapped
}

// A APIError is returned when the BankID RP API returns an error.
type APIError struct {
	ErrorCode string `json:"errorCode"`
	Details   string `json:"details"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s. %s", e.ErrorCode, e.Details)
}
