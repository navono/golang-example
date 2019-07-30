package gops

import (
	"fmt"
	"github.com/google/gops/agent"
	"log"
	"time"
)

func init() {
	fmt.Println("Running app with gops agent.")
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Hour)
}
