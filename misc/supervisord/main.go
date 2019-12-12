package supervisord

import (
	sd "github.com/ochinchina/supervisord"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "supervisord",
		Aliases: []string{"sd"},

		Usage:    "Demonstration of supervisord",
		Action:   sdAction,
		Category: "Misc",
	})
}

func sdAction(c *cli.Context) error {
	sd.NewSupervisor()
	return nil
}
