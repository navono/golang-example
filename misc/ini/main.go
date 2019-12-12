package ini

import (
	"fmt"

	"github.com/urfave/cli"
	"gopkg.in/ini.v1"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "ini",
		Aliases: []string{"ini"},

		Usage:    "Demonstration of load ini file",
		Action:   iniAction,
		Category: "File",
	})
}

func iniAction(c *cli.Context) error {
	cfg, err := ini.Load("./misc/ini/supervisor.conf")
	if err != nil {
		return err
	}

	for _, sec := range cfg.Sections() {
		fmt.Println(sec)
	}

	return nil
}
