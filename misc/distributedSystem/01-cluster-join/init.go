package _1_cluster_join

import (
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "jNode1",
		Aliases: []string{"n1"},

		Usage:    "Start node1",
		Action:   node1,
		Category: "nodeJoin",
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "jNode2",
		Aliases: []string{"n2"},

		Usage:    "Join node2 to cluster",
		Category: "nodeJoin",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node2Join,
			},
		},
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "jNode3",
		Aliases: []string{"n3"},

		Usage:    "Join node3 to cluster",
		Category: "nodeJoin",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node3Join,
			},
		},
	})
}
