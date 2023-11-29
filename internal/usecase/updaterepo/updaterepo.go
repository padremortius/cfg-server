package updaterepo

import (
	"cfg-server/internal/git"
	"cfg-server/internal/svclogger"
	"context"
)

func RunTask(ctxParent context.Context) {
	svclogger.Logger.Logger.Debug().Msgf("Start task 'Update repo'")
	_, cancel := context.WithCancel(ctxParent)
	defer cancel()

	locRepo := git.GitRepo

	dirExists, err := git.DirExists("repo")
	if err != nil {
		svclogger.Logger.Logger.Error().Msgf("Error check repo dir: %v", err)
		return
	}

	if !dirExists {
		git.InitDir("repo")
		if err = locRepo.CloneRepo(); err != nil {
			svclogger.Logger.Logger.Error().Msgf("Error clone repo: %v", err)
			return
		}
	} else {
		if err = locRepo.PullRepo(); err != nil {
			svclogger.Logger.Logger.Error().Msgf("Error pull repo: %v", err)
			return
		}
	}

	svclogger.Logger.Logger.Debug().Msgf("End task 'Update repo'")
}
