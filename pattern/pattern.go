package pattern

import (
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "pattern",
		Aliases: []string{"pt"},

		Usage: "Start exe with pattern",
		Subcommands: []cli.Command{
			{
				Name:   "or",
				Usage:  "start `or` pattern",
				Action: orAction,
			},
		},
		Category: "Pattern",
	})
}
