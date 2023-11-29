package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	BaseApp struct {
		Name        string `env-required:"true" yaml:"name" env:"APP_NAME" json:"name" env-update:"true"`
		ProfileName string `yaml:"profileName,omitempty" json:"profileName,omitempty"`
	}

	Version struct {
		Version        string `json:"version"`
		BuildTimeStamp string `json:"buildTimeStamp,omitempty"`
		GitBranch      string `json:"gitBranch,omitempty"`
		GitHash        string `json:"gitHash,omitempty"`
	}
)

var (
	binVersion = "0.0.1"
)

func init() {
	err := cleanenv.ReadConfig("application.yml", &Cfg)
	if err != nil {
		panic(err)
	}
}

func InitVersion(aBuildNumber, aBuildTimeStamp, aGitBranch, aGitHash string) *Version {
	return &Version{
		Version:        fmt.Sprint(binVersion, ".", aBuildNumber),
		GitBranch:      aGitBranch,
		GitHash:        aGitHash,
		BuildTimeStamp: aBuildTimeStamp,
	}
}
