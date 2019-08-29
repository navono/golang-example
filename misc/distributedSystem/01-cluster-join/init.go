package _1_cluster_join

import (
	"github.com/urfave/cli"
	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "jNode1",
		Aliases: []string{"nj1"},

		Usage:    "Start node1",
		Action:   node1,
		Category: "Cluster-nodeJoin",
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "jNode2",
		Aliases: []string{"nj2"},

		Usage:    "Join node2 to cluster",
		Category: "Cluster-nodeJoin",
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
		Aliases: []string{"nj3"},

		Usage:    "Join node3 to cluster",
		Category: "Cluster-nodeJoin",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node3Join,
			},
		},
	})
}
