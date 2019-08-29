package _1_cluster_join

import (
	"github.com/hashicorp/memberlist"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
	"syscall"
)

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
