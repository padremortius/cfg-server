package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/padremortius/cfg-server/internal/git"
	"github.com/padremortius/cfg-server/pkgs/baseconfig"
	"github.com/padremortius/cfg-server/pkgs/httpserver"
	"github.com/padremortius/cfg-server/pkgs/svclogger"
)

type (
	App struct {
	}

	Config struct {
		App                `yaml:"app" json:"app" validate:"required"`
		baseconfig.BaseApp `yaml:"baseApp" json:"baseApp" validate:"required"`
		Git                git.GitOpts        `yaml:"git" json:"git" validate:"required"`
		HTTP               httpserver.HTTP    `yaml:"http" json:"http" validate:"required"`
		Log                svclogger.Log      `yaml:"logger" json:"logger" validate:"required"`
		Version            baseconfig.Version `json:"version"`
	}
)

func (c *Config) ReadBaseConfig() error {
	if err := cleanenv.ReadConfig("application.yml", c); err != nil {
		return err
	}
	return nil
}

func (c *Config) ReadPwd() error {
	pwd, err := baseconfig.FillPwdMap(c.SecPath)
	if err != nil {
		return err
	}

	c.Git.PrivateKey = pwd["git_PrivateKey"]
	c.Git.Password = pwd["git_password"]

	return nil
}

// NewConfig initializes the configuration by reading environment variables
// and a YAML configuration file.
//
// It returns an error if there is an issue reading the environment variables
// or the configuration file.
func NewConfig(aBuildNumber, aBuildTimeStamp, aGitBranch, aGitHash string) (*Config, error) {
	var cfg Config

	if err := cfg.ReadBaseConfig(); err != nil {
		return &Config{}, fmt.Errorf("NewConfig: %v", err)
	}

	cfg.Version = *baseconfig.InitVersion(aBuildNumber, aBuildTimeStamp, aGitBranch, aGitHash)

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return &Config{}, err
	}

	if _, err := os.Stat(".env"); err == nil {
		if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
			return &Config{}, err
		}
	}

	appConfigName := fmt.Sprint(cfg.Name, "-", cfg.ProfileName, ".yml")

	if cfg.ProfileName == "dev" {
		if err := cleanenv.ReadConfig(appConfigName, &cfg); err != nil {
			return &Config{}, fmt.Errorf("read config error: %v", err)
		}
	}

	if err = cfg.ReadPwd(); err != nil {
		return &Config{}, fmt.Errorf("read password error: %v", err)
	}

	if err := cfg.validateConfig(); err != nil {
		return &Config{}, fmt.Errorf("validation error: %v", err)
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
