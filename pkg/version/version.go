package version

import "strings"

var (
	Program      = "kaga"
	ProgramUpper = strings.ToUpper(Program)
	Version      = "dev"
	GitCommit    = "HEAD"
)
