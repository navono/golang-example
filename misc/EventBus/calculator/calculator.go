package calculator

import (
	"fmt"
	"time"

	"golang-example/misc/EventBus/models"

	"github.com/mustafaturan/bus"
)

var total int64

var c chan *bus.Event

func init() {
	h := bus.Handler{Handle: sum, Matcher: "^order.(created|canceled)$"}
	bus.RegisterHandler("calculator", &h)
	fmt.Printf("Registered calculator handler...\n")

	total = 0
	c = make(chan *bus.Event)

	go calculate()
}

func sum(e *bus.Event) {
	o := e.Data.(*models.Order)
	o.SyncGroup.Add(1)
	c <- e
}

func calculate() {
	for {
		e := <-c
		o := e.Data.(*models.Order)
		amount := int64(o.Amount)
		switch e.Topic.Name {
		case "order.created":
			total += amount
			o.SyncGroup.Done()
		case "order.canceled":
			total -= amount
			time.Sleep(5 * time.Second)
			o.SyncGroup.Done()
		default:
			fmt.Printf("whoops unexpected topic (%s)", e.Topic.Name)
		}
	}
}

// TotalAmount returns total amount
func TotalAmount() int64 {
	return total
}
