package main

import (
	"github.com/padremortius/cfg-server/internal/app"
)

var (
	aBuildNumber    = ""
	aBuildTimeStamp = ""
	aGitBranch      = ""
	aGitHash        = ""
)

func main() {
	app.Run(aBuildNumber, aBuildTimeStamp, aGitBranch, aGitHash)
}
