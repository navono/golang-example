package main

import (
	"fmt"
	//"github.com/google/gops/agent"
	//"log"
	"runtime"
	"time"
)

func main() {
	//if err := agent.Listen(agent.Options{}); err != nil {
	//	log.Fatal(err)
	//}

	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()

	go func(ch chan int) {
		// receive nil ch
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
