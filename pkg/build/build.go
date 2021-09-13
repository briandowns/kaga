package build

import (
	"fmt"
	"os/exec"

	"github.com/briandowns/kaga/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Build(clx *cli.Context, config *config.Config) error {
	var args []string

	if config.Image.Name != "" && config.Image.Tag != "" {
		image := "GOLANG=" + config.Image.Name + ":" + config.Image.Tag
		args = append(args, image)
	}

	if config.SkipValidate {
		args = append(args, "SKIP_VALIDATE=true")
	}

	logrus.Debugf("Building K3s with: make %v", args)

	cmd := exec.Command("make", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout

	if err := cmd.Start(); err != nil {
		return err
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	logrus.Info("Secure K3s build complete")

	return nil
}
