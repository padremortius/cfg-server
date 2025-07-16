package gitdata

import (
	"cfg-server/internal/git"
	"errors"
)

func GetDataFromGit(env, appName, profileName string) (interface{}, error) {
	r := git.GitRepo

	dirExists, err := git.DirExists(r.LocalPath)
	if err != nil {
		return nil, errors.New("Error check repo dir: " + err.Error())
	}

	if !dirExists {
		git.InitDir(r.LocalPath)
		if err = r.CloneRepo(); err != nil {
			return nil, errors.New("Error clone repo: " + err.Error())
		}
	} else {
		if err = r.PullRepo(); err != nil {
			return nil, errors.New("Error pull repo: " + err.Error())
		}
	}

	data, err := r.GetCfgByAppName(env, appName, profileName)
	if err != nil {
		return nil, errors.New("Error get data from git: " + err.Error())
	}
	return data, nil
}
