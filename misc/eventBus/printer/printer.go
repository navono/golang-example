package printer

import (
	"fmt"

	"github.com/mustafaturan/bus"
)

func init() {
	h := bus.Handler{Handle: print, Matcher: ".*"}
	bus.RegisterHandler("printer", &h)
	fmt.Printf("Registered printer handler...\n")
}

func print(e *bus.Event) {
	fmt.Printf("\nEvent for %s: %+v\n", e.Topic.Name, e)
}

// Load will auto load init
func Load() {
}
