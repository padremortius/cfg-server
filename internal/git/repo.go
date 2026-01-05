package git

import (
	"fmt"
	"slices"
	"time"

	"github.com/padremortius/cfg-server/pkgs/common"

	"dario.cat/mergo"
	gogit "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	gitssh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

var GitRepo *Repo = nil

func New(opts GitOpts) (*Repo, error) {
	var (
		publicKey *gitssh.PublicKeys
		keyError  error
		locRepo   *Repo = nil
	)

	if publicKey, keyError = gitssh.NewPublicKeys("git", []byte(opts.PrivateKey), opts.Password); keyError != nil {
		return locRepo, fmt.Errorf("Error parsing ssh private key: %v", keyError)
	}
	if *opts.IgnoreKnownHosts {
		publicKey.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	}

	locRepo = &Repo{
		RepoUrl:    opts.RepoUrl,
		RepoBranch: opts.RepoBranch,
		Auth:       publicKey,
		Depth:      opts.Depth,
		SearchPath: opts.SearchPath,
		LocalPath:  opts.LocalPath,
		SyncTime:   opts.SyncTime,
	}
	return locRepo, nil
}

func (r *Repo) CloneRepo() error {
	currTime := time.Now()
	if currTime.Sub(r.UpdateTime) > r.SyncTime {
		r.RWMutex.Lock()
		defer func() {
			r.RWMutex.Unlock()
			r.UpdateTime = time.Now()
		}()
		_, err := gogit.PlainClone(r.LocalPath, false, &gogit.CloneOptions{
			RemoteName:    "origin",
			URL:           r.RepoUrl,
			Auth:          r.Auth,
			ReferenceName: plumbing.ReferenceName(fmt.Sprint("refs/heads/", r.RepoBranch)),
			Depth:         r.Depth,
			SingleBranch:  true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) PullRepo() error {
	currTime := time.Now()
	if currTime.Sub(r.UpdateTime) > r.SyncTime {
		r.RWMutex.Lock()
		defer func() {
			r.RWMutex.Unlock()
			r.UpdateTime = time.Now()
		}()

		repo, err := gogit.PlainOpen(r.LocalPath)
		if err != nil {
			return fmt.Errorf("Error opening repo: %v", err)
		}

		w, err := repo.Worktree()
		if err != nil {
			return fmt.Errorf("Error getting worktree: %v", err)
		}

		err = w.Pull(&gogit.PullOptions{
			RemoteName:   "origin",
			Auth:         r.Auth,
			Force:        true,
			Depth:        r.Depth,
			SingleBranch: true,
		})
		if (err != nil) && err != gogit.NoErrAlreadyUpToDate {
			return fmt.Errorf("Error pulling repo: %v", err)
		}
	}
	return nil
}

func ItemExists(list []string, item string) bool {
	return slices.Contains(list, item)
}

func initFileList(localPath, env, appName, profileName string) []string {
	var res []string

	listFName := []string{
		fmt.Sprint(localPath, "/application.yaml"),
		fmt.Sprint(localPath, "/application.yml"),
		fmt.Sprint(localPath, "/application", "-", profileName, ".yaml"),
		fmt.Sprint(localPath, "/application", "-", profileName, ".yml"),
		fmt.Sprint(localPath, "/", appName, ".yaml"),
		fmt.Sprint(localPath, "/", appName, ".yml"),
		fmt.Sprint(localPath, "/", appName, "-", profileName, ".yaml"),
		fmt.Sprint(localPath, "/", appName, "-", profileName, ".yaml"),
		fmt.Sprint(localPath, "/", env, "/application.yaml"),
		fmt.Sprint(localPath, "/", env, "/application.yml"),
		fmt.Sprint(localPath, "/", env, "/application", "-", profileName, ".yaml"),
		fmt.Sprint(localPath, "/", env, "/application", "-", profileName, ".yml"),
		fmt.Sprint(localPath, "/", env, "/", appName, ".yaml"),
		fmt.Sprint(localPath, "/", env, "/", appName, ".yml"),
		fmt.Sprint(localPath, "/", env, "/", appName, "-", profileName, ".yaml"),
		fmt.Sprint(localPath, "/", env, "/", appName, "-", profileName, ".yaml"),
	}

	for _, fname := range listFName {
		if common.FileExists(fname) && !ItemExists(res, fname) {
			res = append(res, fname)
		}
	}

	return res
}

func (r *Repo) GetCfgByAppName(envName, appName, profileName string) (any, error) {
	var (
		data, res map[string]any
		err       error
	)

	res = make(map[string]any)

	if envName == "" {
		envName = r.SearchPath
	}

	listFName := initFileList(r.LocalPath, envName, appName, profileName)

	for _, fName := range listFName {
		data, err = common.GetDataFromFile(fName)
		if err != nil {
			return res, err
		}
		err = mergo.Merge(&res, data, mergo.WithOverride)
		if err != nil {
			return res, fmt.Errorf("Error merging file: %v", err)
		}
	}

	return res, nil
}
