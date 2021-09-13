package cmds

import (
	"errors"

	"github.com/briandowns/kaga/pkg/build"
	"github.com/briandowns/kaga/pkg/config"
	"github.com/briandowns/kaga/pkg/version"
	"github.com/urfave/cli"
)

var buildFlags = []cli.Flag{
	DebugFlag,
	cli.StringFlag{
		Name:        "c",
		Usage:       "config file",
		EnvVar:      version.ProgramUpper + "_CONFIG_FILE",
		Destination: nil,
	},
}

// NewBuildCommand
func NewBuildCommand() cli.Command {
	return cli.Command{
		Name:            "build",
		Usage:           "Build k3s based on the provided config",
		SkipFlagParsing: false,
		SkipArgReorder:  true,
		Action:          run,
		Flags:           buildFlags,
	}
}

// run
func run(clx *cli.Context) error {
	if clx.String("c") == "" {
		return errors.New("config file required")
	}

	config, err := config.Load(clx.String("c"))
	if err != nil {
		return err
	}

	return build.Build(clx, config)
}
