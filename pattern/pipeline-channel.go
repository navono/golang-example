package main

import (
	"fmt"
)

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func multiply(done <-chan interface{},
	intStream <-chan int,
	multiplier int) <-chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}

func add(done <-chan interface{},
	intStream <-chan int,
	additive int,
) <-chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addedStream <- i + additive:
			}
		}
	}()
	return addedStream
}

func testPipeline() {
	done := make(chan interface{})
	defer close(done)

	// 将值转换到 channel 中
	// 每个 stage 都是抢占式的
	intStream := generator(done, 1, 2, 3, 4)
	pipline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)
	for v := range pipline {
		fmt.Println(v)
	}
}

func main() {
	testPipeline()
}
