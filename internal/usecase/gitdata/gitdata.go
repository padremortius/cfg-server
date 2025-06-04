package gitdata

import (
	"cfg-server/internal/git"
	"cfg-server/internal/svclogger"
)

func GetDataFromGit(env, appName, profileName string) (interface{}, error) {
	r := git.GitRepo

	dirExists, err := git.DirExists(r.LocalPath)
	if err != nil {
		svclogger.Logger.Logger.Error().Msgf("Error check repo dir: %v", err)
		return nil, err
	}

	if !dirExists {
		git.InitDir(r.LocalPath)
		if err = r.CloneRepo(); err != nil {
			svclogger.Logger.Logger.Error().Msgf("Error clone repo: %v", err)
			return nil, err
		}
	} else {
		if err = r.PullRepo(); err != nil {
			svclogger.Logger.Logger.Error().Msgf("Error pull repo: %v", err)
			return nil, err
		}
	}

	data, err := r.GetCfgByAppName(env, appName, profileName)
	if err != nil {
		svclogger.Logger.Logger.Error().Msgf("Error get data from git: %v", err)
		return nil, err
	}
	return data, nil
}
