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

func Run() {
	flag.Parse()
	ctxmain := context.Background()

	log := svclogger.New("")

	if err := config.NewConfig(); err != nil {
		log.Logger.Error().Msgf("Config error: %v", err)
		os.Exit(-1)
	}
	if *repoUrl != "" {
		config.Cfg.Git.RepoUrl = *repoUrl
	}

	if *repoBranch != "" {
		config.Cfg.Git.RepoBranch = *repoBranch
	}
	if *searchPath != "" {
		config.Cfg.Git.SearchPath = *searchPath
	}

	if config.Cfg.Git.IgnoreKnownHosts == nil {
		config.Cfg.Git.IgnoreKnownHosts = new(bool)
		*config.Cfg.Git.IgnoreKnownHosts = true
	}

	if config.Cfg.Git.Depth == 0 {
		config.Cfg.Git.Depth = 5
	}

	err := config.Cfg.ValidateConfig()
	if err != nil {
		log.Logger.Error().Msgf("Config error: %v", err)
		os.Exit(-1)
	}

	log.Logger.Info().Msgf("Start application. Version: %v", config.Cfg.Version.Version)

	ctx, cancel := context.WithTimeout(ctxmain, config.Cfg.HTTP.Timeouts.Shutdown)
	defer cancel()

	//init gitRepo
	log.Logger.Info().Msgf("Start clone repo. Repo url: %v, branch: %v", config.Cfg.Git.RepoUrl, config.Cfg.Git.RepoBranch)
	git.GitRepo = git.New(config.Cfg.Git)
	if err := git.InitDir(git.GitRepo.LocalPath); err != nil {
		log.Logger.Error().Msgf("Error init dir: %v", err)
		os.Exit(-1)
	}
	if err := git.GitRepo.CloneRepo(); err != nil {
		log.Logger.Error().Msgf("Error clone repo: %v", err)
		os.Exit(-1)
	}
	log.Logger.Info().Msg("End clone repo.")

	log.ChangeLogLevel(config.Cfg.Log.Level)

	// HTTP Server
	log.Logger.Info().Msgf("Start web-server on port %v", config.Cfg.HTTP.Port)

	httpServer := httpserver.New(ctx, log, &config.Cfg.HTTP)
	baserouting.InitBaseRouter(httpServer.Handler)
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
	if err := httpServer.Shutdown(config.Cfg.HTTP.Timeouts.Shutdown); err != nil {
		log.Logger.Error().Msgf("app - Run - httpServer.Shutdown: %v", err.Error())
	}
}
