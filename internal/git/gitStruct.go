package git

import (
	"sync"
	"time"

	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type (
	GitOpts struct {
		RepoUrl          string `yaml:"repoUrl" json:"repoUrl" validate:"required"`
		RepoBranch       string `yaml:"repoBranch" json:"repoBranch" validate:"required"`
		Depth            int    `yaml:"depth" json:"depth" validate:"required"`
		SearchPath       string `yaml:"searchPath" json:"searchPath" validate:"required"`
		IgnoreKnownHosts *bool  `yaml:"ignoreKnownHosts" json:"ignoreKnownHosts" validate:"required"`
		PrivateKey       string `yaml:"-" json:"-" validate:"required"`
		Password         string `yaml:"-,omitempty" json:"-,omitempty"`
	}

	Repo struct {
		RepoUrl    string    `yaml:"repoUrl" json:"repoUrl" validate:"required"`
		LocalPath  string    `yaml:"localPath" json:"localPath" validate:"required"`
		RepoBranch string    `yaml:"repoBranch" json:"repoBranch" validate:"required"`
		Depth      int       `yaml:"depth" json:"depth" validate:"required"`
		SearchPath string    `yaml:"searchPath" json:"searchPath" validate:"required"`
		UpdateTime time.Time `yaml:"-" json:"-"`
		RWMutex    sync.RWMutex
		Auth       *gitssh.PublicKeys `yaml:"-" json:"-"`
	}
)
