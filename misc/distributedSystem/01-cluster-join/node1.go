package _1_cluster_join

import (
	"github.com/hashicorp/memberlist"
	"github.com/urfave/cli"
	"golang-example/cmd"

	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "node1",
		Aliases: []string{"n1"},

		Usage:    "Start node1",
		Action:   node1,
		Category: "memberlist",
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "node2",
		Aliases: []string{"n2"},

		Usage:    "Join node2 to cluster",
		Category: "memberlist",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node2Join,
			},
		},
	})

	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "node3",
		Aliases: []string{"n3"},

		Usage:    "Join node3 to cluster",
		Category: "memberlist",
		Subcommands: []cli.Command{
			{
				Name:   "join",
				Usage:  "join exist cluster",
				Action: node3Join,
			},
		},
	})
}

func node1(c *cli.Context) error {
	conf := memberlist.DefaultLocalConfig()
	conf.Name = "node1"

	list, err := memberlist.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	local := list.LocalNode()
	log.Printf("node1 at %s:%d", local.Addr.To4().String(), local.Port)

	log.Printf("wait for other member connections")
	waitSignal()

	return nil
}

func waitSignal() {
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGTERM)
	signal.Notify(signalChan, syscall.SIGINT)
	for {
		select {
		case s := <-signalChan:
			log.Printf("signal %s happen", s.String())
			return
		}
	}
}
