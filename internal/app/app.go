package app

import (
	"cfg-server/internal/config"
	"cfg-server/internal/controller/baserouting"
	v1 "cfg-server/internal/controller/v1"
	"cfg-server/internal/git"
	"cfg-server/internal/httpserver"
	"cfg-server/internal/svclogger"
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	repoUrl    = flag.String("repoUrl", "", "Git repo path")
	repoBranch = flag.String("repoBranch", "", "Git branch")
	searchPath = flag.String("searchPath", "", "Git search path")
)

func Run(ver config.Version) {
	flag.Parse()
	ctxmain := context.Background()

	log := svclogger.New("")

	appCfg, err := config.NewConfig()
	if err != nil {
		log.Logger.Error().Msgf("Config error: %v", err)
		os.Exit(-1)
	}

	appCfg.Version = ver

	if *repoUrl != "" {
		appCfg.Git.RepoUrl = *repoUrl
	}

	if *repoBranch != "" {
		appCfg.Git.RepoBranch = *repoBranch
	}
	if *searchPath != "" {
		appCfg.Git.SearchPath = *searchPath
	}

	if appCfg.Git.IgnoreKnownHosts == nil {
		appCfg.Git.IgnoreKnownHosts = new(bool)
		*appCfg.Git.IgnoreKnownHosts = true
	}

	if appCfg.Git.Depth == 0 {
		appCfg.Git.Depth = 5
	}

	log.Logger.Info().Msgf("Start application. Version: %v", appCfg.Version.Version)

	ctx, cancel := context.WithTimeout(ctxmain, appCfg.HTTP.Timeouts.Shutdown)
	defer cancel()

	//init gitRepo
	log.Logger.Info().Msgf("Start clone repo. Repo url: %v, branch: %v", appCfg.Git.RepoUrl, appCfg.Git.RepoBranch)
	git.GitRepo = git.New(appCfg.Git)
	if err := git.InitDir(git.GitRepo.LocalPath); err != nil {
		log.Logger.Error().Msgf("Error init dir: %v", err)
		os.Exit(-1)
	}
	if err := git.GitRepo.CloneRepo(); err != nil {
		log.Logger.Error().Msgf("Error clone repo: %v", err)
		os.Exit(-1)
	}
	log.Logger.Info().Msg("End clone repo.")

	log.ChangeLogLevel(appCfg.Log.Level)

	// HTTP Server
	log.Logger.Info().Msgf("Start web-server on port %v", appCfg.HTTP.Port)

	httpServer := httpserver.New(ctx, log, &appCfg.HTTP)
	baserouting.InitBaseRouter(httpServer.Handler, *appCfg)
	v1.InitAppRouter(httpServer.Handler)
	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGTERM,
		syscall.SIGUSR1,
		syscall.SIGUSR2)

	select {
	case s := <-interrupt:
		log.Logger.Info().Msgf("app - Run - signal: %v", s.String())
	case err := <-httpServer.Notify():
		log.Logger.Error().Msgf("app - Run - httpServer.Notify: %v", err.Error())
	}

	// Shutdown
	if err := httpServer.Shutdown(appCfg.HTTP.Timeouts.Shutdown); err != nil {
		log.Logger.Error().Msgf("app - Run - httpServer.Shutdown: %v", err.Error())
	}
}
