package config

import (
	"fmt"
	"os"

	gojson "github.com/goccy/go-json"
)

type (
	App struct {
	}

	pwdData map[string]string
)

func ReadPwd() error {
	fname := fmt.Sprint("./", Cfg.BaseApp.Name, ".json")
	if _, err := os.Stat(fname); err == nil {
		pwdFile, err := os.ReadFile(fname)
		if err != nil {
			return err
		}

		if err = gojson.Unmarshal(pwdFile, &pwd); err != nil {
			return err
		}

		Cfg.Git.PrivateKey = pwd["git.pKey"]
		Cfg.Git.Password = pwd["git.password"]
	}

	return nil
}
