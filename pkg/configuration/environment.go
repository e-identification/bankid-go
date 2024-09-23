package configuration

// Environment contains the environment specific fields.
type Environment struct {
	BaseURL     string
	Certificate string
}

// NewEnvironment creates a new environment.
func NewEnvironment(baseURL string, certificate string) *Environment {
	return &Environment{baseURL, certificate}
}
