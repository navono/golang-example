package memberlist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/hashicorp/memberlist"
	"github.com/pborman/uuid"
	"github.com/urfave/cli"

	"golang-example/cmd"
)

func init() {
	cmd.Cmds = append(cmd.Cmds, cli.Command{
		Name:    "memberlist",
		Aliases: []string{"memberlist"},

		Usage:  "Demonstration of memberlist",
		Action: memberlistAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "members",
				Usage:    "comma separated list of members",
				Required: false,
			},
			cli.IntFlag{
				Name:  "port",
				Value: 4001,
				Usage: "http port",
			},
		},
		Category: "Distributed",
	})
}

func memberlistAction(c *cli.Context) error {
	members := c.String("members")
	port := c.Int("port")
	if err := start(members); err != nil {
		fmt.Println(err)
		return err
	}

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/del", delHandler)
	http.HandleFunc("/get", getHandler)
	fmt.Printf("Listening on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type (
	broadcast struct {
		msg    []byte
		notify chan<- struct{}
	}

	delegate      struct{}
	eventDelegate struct{}

	update struct {
		Action string
		Data   map[string]string
	}
)

var (
	mtx        sync.RWMutex
	items      = map[string]string{}
	broadcasts *memberlist.TransmitLimitedQueue
)

func (b *broadcast) Invalidates(others memberlist.Broadcast) bool {
	return false
}

func (b *broadcast) Message() []byte {
	return b.msg
}

func (b *broadcast) Finished() {
	if b.notify != nil {
		close(b.notify)
	}
}

func (d *delegate) NodeMeta(limit int) []byte {
	return []byte{}
}

func (d *delegate) NotifyMsg(b []byte) {
	if len(b) == 0 {
		return
	}

	switch b[0] {
	case 'd':
		var updates []*update
		if err := json.Unmarshal(b[1:], &updates); err != nil {
			return
		}

		mtx.Lock()
		for _, u := range updates {
			for k, v := range u.Data {
				switch u.Action {
				case "add":
					items[k] = v
				case "del":
					delete(items, k)
				}
			}
		}
		mtx.Unlock()
	}
}

func (d *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return broadcasts.GetBroadcasts(overhead, limit)
}

func (d *delegate) LocalState(join bool) []byte {
	mtx.RLock()
	m := items
	mtx.RUnlock()
	b, _ := json.Marshal(m)
	return b
}

func (d *delegate) MergeRemoteState(buf []byte, join bool) {
	if len(buf) == 0 {
		return
	}
	if !join {
		return
	}

	var m map[string]string
	if err := json.Unmarshal(buf, &m); err != nil {
		return
	}

	mtx.Lock()
	for k, v := range m {
		items[k] = v
	}
	mtx.Unlock()
}

func (ed *eventDelegate) NotifyJoin(node *memberlist.Node) {
	fmt.Println("A node has joined: " + node.String())
}

func (ed *eventDelegate) NotifyLeave(node *memberlist.Node) {
	fmt.Println("A node has left: " + node.String())
}

func (ed *eventDelegate) NotifyUpdate(node *memberlist.Node) {
	fmt.Println("A node was updated: " + node.String())
}

func start(members string) error {
	hostname, _ := os.Hostname()
	c := memberlist.DefaultLocalConfig()
	c.Events = &eventDelegate{}
	c.Delegate = &delegate{}
	c.BindPort = 0
	c.Name = hostname + "-" + uuid.NewUUID().String()
	m, err := memberlist.Create(c)
	if err != nil {
		return err
	}
	if len(members) > 0 {
		parts := strings.Split(members, ",")
		_, err := m.Join(parts)
		if err != nil {
			return err
		}
	}
	broadcasts = &memberlist.TransmitLimitedQueue{
		NumNodes: func() int {
			return m.NumMembers()
		},
		RetransmitMult: 3,
	}
	node := m.LocalNode()
	fmt.Printf("Local member %s:%d\n", node.Addr, node.Port)
	return nil
}
