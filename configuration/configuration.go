package configuration

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"time"
)

type Environment struct {
	BaseUrl               string
	CertificationFilePath string
}

type Configuration struct {
	Environment *Environment
	CertFile    string
	KeyFile     string
	Timeout     time.Duration
}

func New(environment *Environment, certFile string, keyFile string, options ...Option) *Configuration {
	instance := &Configuration{Environment: environment, CertFile: certFile, KeyFile: keyFile, Timeout: 60}

	// Apply options if there are any, can overwrite default
	for _, option := range options {
		option(instance)
	}

	return instance
}

// Option definition
type Option func(*Configuration)

// Function to create Option func to set the timeout limit
// nolint:deadcode
func setTimeout(timeout time.Duration) Option {
	return func(subject *Configuration) {
		subject.Timeout = timeout
	}
}

var (
	TestEnvironment       = Environment{BaseUrl: "https://appapi2.test.bankid.com/rp/v5", CertificationFilePath: GetResourcePath("certificates/ca.test.crt")}
	ProductionEnvironment = Environment{BaseUrl: "https://appapi2.bankid.com/rp/v5", CertificationFilePath: GetResourcePath("certificates/ca.prod.crt")}
)

func GetResourceDirectoryPath() (directory string, err error) {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("could not derive directory path")
	}

	return fmt.Sprintf("%s/%s", path.Dir(filename), "../resource"), nil
}

func GetResourcePath(path string) (directory string) {
	dir, err := GetResourceDirectoryPath()

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", dir, path)
}
