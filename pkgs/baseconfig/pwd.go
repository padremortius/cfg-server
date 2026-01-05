package baseconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/padremortius/cfg-server/pkgs/common"
)

func FillPwdMap(path string) (map[string]string, error) {
	pwd := make(map[string]string, 0)
	if len(path) < 1 {
		return pwd, fmt.Errorf("path to file with secrets is empty")
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return pwd, err
	}

	for _, entry := range entries {
		buff, err := common.ReadFile(filepath.Join(path, entry.Name()))
		if err != nil {
			return pwd, fmt.Errorf("error read file %v with error: %v", entry.Name(), err.Error())
		}
		pwd[entry.Name()] = string(buff)
	}
	return pwd, nil
}
