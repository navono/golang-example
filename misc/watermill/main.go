package watermill

import (
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "watermill",
		Aliases: []string{"wm"},

		Usage:    "Demonstration of watermill",
		Category: "event driven",
		Subcommands: []cli.Command{
			{
				Name:   "ch",
				Usage:  "channel example",
				Action: channelAction,
			},
			{
				Name:   "http",
				Usage:  "http example",
				Action: httpAction,
			},
			{
				Name:   "tx",
				Usage:  "transaction example",
				Action: txAction,
			},
		},
	})
}
