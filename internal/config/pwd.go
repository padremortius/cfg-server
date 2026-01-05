package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/padremortius/cfg-server/pkgs/common"
)

type pwdData map[string]string

func fillPwdMap(path string) (pwdData, error) {
	pwd := make(pwdData, 0)
	if len(path) < 1 {
		return pwd, errors.New("SEC_PATH is empty")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return pwd, err
	}

	for _, entry := range entries {
		buff, err := common.ReadFile(filepath.Join(path, entry.Name()))
		if err != nil {
			return pwd, errors.New("Error read file: " + entry.Name() + " Error: " + err.Error())
		}
		pwd[entry.Name()] = string(buff)
	}
	return pwd, nil
}
