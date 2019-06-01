package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func signalChild() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					// Do something
					fmt.Println(s)
				case <-done:
					fmt.Println("receive Done")
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	// terminated := doWork(done, nil)
	doWork(done, nil)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// Cancel the operation after 1 sec
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
		wg.Done()
	}()

	// join the doWork goroutine
	// <-terminated

	wg.Wait()

	// 此处是增加 doWork 内部 terminated 退出的机会，和上述的 <-terminated 目的一样
	// 区别在于 <-terminated 是确定性的
	time.Sleep(1 * time.Second)

	fmt.Println("Done")
}

func blockRead() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStream exited.")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random int: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}

func goroutineCancel() {
	// signalChild()
	blockRead()
}
