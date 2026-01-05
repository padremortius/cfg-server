package gitdata

import (
	"fmt"

	"github.com/padremortius/cfg-server/internal/git"
	"github.com/padremortius/cfg-server/pkgs/common"
)

func GetDataFromGit(env, appName, profileName string) (interface{}, error) {
	r := git.GitRepo

	dirExists, err := common.DirExists(r.LocalPath)
	if err != nil {
		return nil, fmt.Errorf("Error check repo dir: %v", err)
	}

	if !dirExists {
		common.InitDir(r.LocalPath)
		if err = r.CloneRepo(); err != nil {
			return nil, fmt.Errorf("Error clone repo: %v", err)
		}
	} else {
		if err = r.PullRepo(); err != nil {
			return nil, fmt.Errorf("Error pull repo: %v", err)
		}
	}

	data, err := r.GetCfgByAppName(env, appName, profileName)
	if err != nil {
		return nil, fmt.Errorf("Error get data from git: %v", err)
	}
	return data, nil
}
