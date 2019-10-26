package counter

import (
	"fmt"

	"github.com/mustafaturan/bus"

	"golang-example/misc/EventBus/models"
)

var topics map[string]uint

var c chan *bus.Event

func init() {
	h := bus.Handler{Handle: count, Matcher: ".*"}
	bus.RegisterHandler("counter", &h)
	fmt.Printf("Registered counter handler...\n")

	topics = make(map[string]uint, 0)
	c = make(chan *bus.Event)

	go increment()
}

func count(e *bus.Event) {
	o := e.Data.(*models.Order)
	o.SyncGroup.Add(1)
	c <- e
}

func increment() {
	for {
		e := <-c
		o := e.Data.(*models.Order)
		n := e.Topic.Name
		if count, ok := topics[n]; ok {
			topics[n] = count + 1
		} else {
			topics[n] = 1
		}

		o.SyncGroup.Done()
	}
}

// FetchEventCount returns total event count for the topic
func FetchEventCount(topicName string) uint {
	return topics[topicName]
}
