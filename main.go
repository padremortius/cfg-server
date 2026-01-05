package main

import (
	"github.com/padremortius/cfg-server/internal/app"
	"github.com/padremortius/cfg-server/internal/config"
)

var (
	aBuildNumber    = ""
	aBuildTimeStamp = ""
	aGitBranch      = ""
	aGitHash        = ""
)

func main() {

	ver := *config.InitVersion(aBuildNumber, aBuildTimeStamp, aGitBranch, aGitHash)
	app.Run(ver)
}
