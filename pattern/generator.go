package main

import (
	"fmt"
	"math/rand"
)

// Repeat generator
func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

// RepeatFn generator
func RepeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()
	return valueStream
}

// Take generator
func Take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// ToString generator
func ToString(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
	stringStream := make(chan string)
	go func() {
		defer close(stringStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case stringStream <- v.(string):
			}
		}
	}()
	return stringStream
}

// ToInt generator
func ToInt(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for v := range valueStream {
			select {
			case <-done:
				return
			case intStream <- v.(int):
			}
		}
	}()
	return intStream
}

func testRepeat() {
	done := make(chan interface{})
	defer close(done)

	for num := range Take(done, Repeat(done, 1, 2), 10) {
		fmt.Printf("%v ", num)
	}
}

func testRepeatFn() {
	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} { return rand.Int() }
	for num := range Take(done, RepeatFn(done, rand), 10) {
		fmt.Println(num)
	}
}

func testTypeContert() {
	done := make(chan interface{})
	defer close(done)

	var message string
	for token := range ToString(done, Take(done, Repeat(done, "I", "am."), 5)) {
		message += token
	}
	fmt.Printf("message: %s...", message)
}

// func main() {
// 	testTypeContert()
// }
