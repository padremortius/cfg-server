package config

import (
	"cfg-server/internal/git"
	"cfg-server/internal/httpserver"
	"cfg-server/internal/svclogger"
	"errors"
	"fmt"
	"os"

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
func NewConfig() (*Config, error) {
	var cfg Config

	if err := cfg.ReadBaseConfig(); err != nil {
		return &Config{}, errors.New("NewConfig: " + err.Error())
	}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &Config{}, err
	}

	if _, err := os.Stat(".env"); err == nil {
		if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
			return &Config{}, err
		}
	}

	appConfigName := fmt.Sprint(cfg.BaseApp.Name, "-", cfg.BaseApp.ProfileName, ".yml")

	if cfg.BaseApp.ProfileName == "dev" {
		if err := cleanenv.ReadConfig(appConfigName, &cfg); err != nil {
			return &Config{}, errors.New("Read config error: " + err.Error())
		}
	}

	if err = cfg.ReadPwd(); err != nil {
		return &Config{}, errors.New("Read password error: " + err.Error())
	}

	if err := cfg.validateConfig(); err != nil {
		return &Config{}, errors.New("Validation error: " + err.Error())
	}
	return &cfg, nil
}

func (c *Config) validateConfig() error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}
