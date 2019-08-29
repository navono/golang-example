package _1_cluster_join

import (
	"github.com/hashicorp/memberlist"
	"github.com/urfave/cli"
	"log"
	"time"
)

func node3Join(c *cli.Context) error {
	if !c.Args().Present() {
		return cli.NewExitError("no master cluster address", 3)
	}

	conf := memberlist.DefaultLocalConfig()
	conf.Name = "node3"
	conf.BindPort = 7948 // avoid port confliction
	conf.AdvertisePort = conf.BindPort

	list, err := memberlist.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	local := list.LocalNode()
	log.Printf("node3 at %s:%d", local.Addr.To4().String(), local.Port)

	join := c.Args().First()
	log.Printf("cluster join to %s", join)

	if _, err := list.Join([]string{join}); err != nil {
		log.Fatal(err)
	}

	for _, member := range list.Members() {
		log.Printf("Member: %s(%s:%d)", member.Name, member.Addr.To4().String(), member.Port)
	}

	log.Printf("wait 10 sec")
	time.Sleep(10 * time.Second)

	log.Printf("leave cluster")
	_ = list.Leave(5 * time.Second)

	return nil
}
