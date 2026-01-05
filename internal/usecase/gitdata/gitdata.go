package gitdata

import (
	"fmt"

	"github.com/padremortius/cfg-server/internal/git"
	"github.com/padremortius/cfg-server/pkgs/common"
)

func GetDataFromGit(env, appName, profileName string) (any, error) {
	r := git.GitRepo

	dirExists, err := common.DirExists(r.LocalPath)
	if err != nil {
		return nil, fmt.Errorf("error check repo dir: %v", err)
	}

	if !dirExists {
		if err = common.InitDir(r.LocalPath); err != nil {
			return nil, fmt.Errorf("error init dir %v with error: %v", r.LocalPath, err)
		}
		if err = r.CloneRepo(); err != nil {
			return nil, fmt.Errorf("error clone repo: %v", err)
		}
	} else {
		if err = r.PullRepo(); err != nil {
			return nil, fmt.Errorf("error pull repo: %v", err)
		}
	}

	data, err := r.GetCfgByAppName(env, appName, profileName)
	if err != nil {
		return nil, fmt.Errorf("error get data from git: %v", err)
	}
	return data, nil
}
