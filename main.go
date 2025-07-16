package main

import (
	"cfg-server/internal/app"
	"cfg-server/internal/config"
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
