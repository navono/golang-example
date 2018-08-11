package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func primeFinder(
	done <-chan interface{},
	intStream <-chan int,
) <-chan interface{} {
	primeStream := make(chan interface{})
	go func() {
		defer close(primeStream)
		for integer := range intStream {
			integer--
			prime := true
			for divisor := integer - 1; divisor > 1; divisor-- {
				if integer%divisor == 0 {
					prime = false
					break
				}
			}

			if prime {
				select {
				case <-done:
					return
				case primeStream <- integer:
				}
			}
		}
	}()
	return primeStream
}

func main() {
	rand := func() interface{} {
		return rand.Intn(50000000)
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := ToInt(done, RepeatFn(done, rand))

	// fan-out with CPU numbers
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan interface{}, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	// fan-in the channels
	for prime := range Take(done, FanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
