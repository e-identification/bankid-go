package configuration

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// Environment contains the environment specific fields.
type Environment struct {
	BaseURL               string
	CertificationFilePath string
}

// Configuration contains the configuration specific fields.
type Configuration struct {
	Environment *Environment
	CertFile    string
	KeyFile     string
	Timeout     time.Duration
}

// New creates a new configuration.
func New(environment *Environment, certFile string, keyFile string, options ...Option) *Configuration {
	instance := &Configuration{Environment: environment, CertFile: certFile, KeyFile: keyFile, Timeout: 60}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(instance)
	}

	return instance
}

// Option definition.
type Option func(*Configuration)

// Function to create Option func to set the timeout limit.
// nolint:deadcode, unused
func setTimeout(timeout time.Duration) Option {
	return func(subject *Configuration) {
		subject.Timeout = timeout
	}
}

var (
	// TestEnvironment contains the environment specific fields for the test environment.
	TestEnvironment = Environment{BaseURL: "https://appapi2.test.bankid.com/rp/v5", CertificationFilePath: GetResourcePath("certificates/ca.test.crt")}
	// ProductionEnvironment contains the environment specific fields for the production environment.
	ProductionEnvironment = Environment{BaseURL: "https://appapi2.bankid.com/rp/v5", CertificationFilePath: GetResourcePath("certificates/ca.prod.crt")}
)

// GetResourceDirectoryPath returns the full path to the resource directory.
func GetResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", fmt.Errorf("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "../resource"), nil
}

// GetResourcePath returns the full path to the resource.
func GetResourcePath(path string) (directory string) {
	dir, err := GetResourceDirectoryPath()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}
