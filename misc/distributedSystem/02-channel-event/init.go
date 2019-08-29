package _2_channel_event

import (
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "cNode1",
		Aliases: []string{"ce1"},

		Usage:    "Start node1",
		Action:   node1,
		Category: "channelEvent",
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "cNode2",
		Aliases: []string{"ce2"},

		Usage:    "Join node2 to cluster",
		Category: "channelEvent",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node2Join,
			},
		},
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "cNode3",
		Aliases: []string{"ce3"},

		Usage:    "Join node3 to cluster",
		Category: "channelEvent",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node3Join,
			},
		},
	})
}
