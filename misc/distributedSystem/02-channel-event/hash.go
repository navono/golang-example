package _2_channel_event

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/serialx/hashring"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type MyEventDelegate struct {
	consistent *hashring.HashRing
}

func (d *MyEventDelegate) NotifyJoin(node *memberlist.Node) {
	hostPort := fmt.Sprintf("%s:%d", node.Addr.To4().String(), node.Port)
	log.Printf("join %s", hostPort)

	if d.consistent == nil {
		d.consistent = hashring.New([]string{hostPort})
	} else {
		d.consistent = d.consistent.AddNode(hostPort)
	}
}

func (d *MyEventDelegate) NotifyLeave(node *memberlist.Node) {
	hostPort := fmt.Sprintf("%s:%d", node.Addr.To4().String(), node.Port)
	log.Printf("leave %s", hostPort)

	if d.consistent != nil {
		d.consistent = d.consistent.RemoveNode(hostPort)
	}
}

func (d *MyEventDelegate) NotifyUpdate(node *memberlist.Node) {
	// skip
}

func waitSignal(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	for {
		select {
		case s := <-signalChan:
			log.Printf("signal %s happen", s.String())
			cancel()
		}
	}
}
