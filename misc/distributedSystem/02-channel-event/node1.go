package _2_channel_event

import (
	"github.com/hashicorp/memberlist"
	"github.com/urfave/cli"

	"context"
	"fmt"
	"log"
	"time"
)

func node1(c *cli.Context) error {
	conf := memberlist.DefaultLocalConfig()
	conf.Name = "node1"
	conf.BindPort = 7947 // avoid port confliction
	conf.AdvertisePort = conf.BindPort
	conf.Events = new(MyEventDelegate)

	list, err := memberlist.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	local := list.LocalNode()
	_, _ = list.Join([]string{
		fmt.Sprintf("%s:%d", local.Addr.To4().String(), local.Port),
	})

	stopCtx, cancel := context.WithCancel(context.TODO())
	go waitSignal(cancel)

	tick := time.NewTicker(3 * time.Second)
	run := true
	for run {
		select {
		case <-tick.C:
			evt := conf.Events.(*MyEventDelegate)
			if evt == nil {
				log.Printf("consistent isnt initialized")
				continue
			}
			log.Printf("current node size: %d", evt.consistent.Size())

			keys := []string{"hello", "world"}
			for _, key := range keys {
				node, ok := evt.consistent.GetNode(key)
				if ok == true {
					log.Printf("node1 search %s => %s", key, node)
				} else {
					log.Printf("no node available")
				}
			}
		case <-stopCtx.Done():
			log.Printf("stop called")
			run = false
		}
	}
	tick.Stop()
	log.Printf("bye.")

	return nil
}
