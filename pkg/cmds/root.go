package cmds

import (
	"fmt"
	"runtime"

	"github.com/briandowns/kaga/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	Debug     bool
	DebugFlag = cli.BoolFlag{
		Name:        "debug",
		Usage:       "Turn on debug logging",
		Destination: &Debug,
		EnvVar:      version.ProgramUpper + "_DEBUG",
	}
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = "kaga"
	app.Usage = "Secure K3s Build and Install"
	app.Version = fmt.Sprintf("%s (%s)", version.Version, version.GitCommit)
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
		fmt.Printf("go version %s\n", runtime.Version())
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Turn on debug logging",
			Destination: &Debug,
			EnvVar:      "KAGA_DEBUG",
		},
	}

	app.Before = func(clx *cli.Context) error {
		if Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	return app
}
