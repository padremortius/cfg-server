package config

import (
	"cfg-server/internal/git"
	"cfg-server/internal/httpserver"
	"cfg-server/internal/svclogger"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App     `yaml:"app" json:"app" validate:"required"`
	BaseApp `yaml:"baseApp" json:"baseApp" validate:"required"`
	Git     git.GitOpts     `yaml:"git" json:"git" validate:"required"`
	HTTP    httpserver.HTTP `yaml:"http" json:"http" validate:"required"`
	Log     svclogger.Log   `yaml:"logger" json:"logger" validate:"required"`
	Version `json:"version"`
}

var (
	Cfg Config
	pwd pwdData
)

// NewConfig initializes the configuration by reading environment variables
// and a YAML configuration file.
//
// It returns an error if there is an issue reading the environment variables
// or the configuration file.
func NewConfig() error {
	err := cleanenv.ReadEnv(&Cfg)
	if err != nil {
		return err
	}

	appConfigName := fmt.Sprint(Cfg.BaseApp.Name, "-", Cfg.BaseApp.ProfileName, ".yml")

	if Cfg.BaseApp.ProfileName == "dev" {
		if err := cleanenv.ReadConfig(appConfigName, &Cfg); err != nil {
			return err
		}
	}

	if err = ReadPwd(); err != nil {
		return errors.New("Read password error: " + err.Error())
	}

	// if err := Cfg.validateConfig(); err != nil {
	// 	return err
	// }
	return nil
}

func (c *Config) ValidateConfig() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
