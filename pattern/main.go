package main

import (
	"fmt"
)

func testTee() {
	done := make(chan interface{})
	defer close(done)

	out1, out2 := Tee(done, Take(done, Repeat(done, 1, 2), 4))
	for val1 := range out1 {
		fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	}
}

func testBridge() {
	genVals := func() <-chan <-chan interface{} {
		chanStream := make(chan (<-chan interface{}))
		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()
		return chanStream
	}

	for v := range Bridge(nil, genVals()) {
		fmt.Printf("%v ", v)
	}
}

func main() {
	// testTee()
	testBridge()
}
