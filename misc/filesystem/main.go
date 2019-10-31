package filesystem

import (
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "fs",
		Aliases: []string{"fs"},

		Usage:    "Demonstration of file system",
		Category: "FS",
		Subcommands: []cli.Command{
			{
				Name:   "afero",
				Usage:  "afero example",
				Action: aferoAction,
			},
		},
	})
}
