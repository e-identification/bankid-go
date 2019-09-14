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

func NewConfiguration(environment *Environment, certFile string, keyFile string, timeout time.Duration) *Configuration {
	if timeout == -1 {
		timeout = 5
	}

	return &Configuration{Environment: environment, CertFile: certFile, KeyFile: keyFile, Timeout: timeout}
}

var (
	TestEnvironment       = Environment{BaseUrl: "https://appapi2.test.bankid.com/rp/v5", CertificationFilePath: GetResourcePath("certificates/ca.test.crt")}
	ProductionEnvironment = Environment{BaseUrl: "https://appapi2.bankid.com/rp/v5", CertificationFilePath: GetResourcePath("certificates/ca.test.crt")}
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
