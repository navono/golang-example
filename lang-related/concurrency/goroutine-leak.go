package concurrency

import (
	"fmt"
)

func memLeak() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)

			for s := range strings {
				// Do something
				fmt.Println(s)
			}
		}()
		return completed
	}

	// 传入 nil 导致 doWork 内部内存泄漏
	doWork(nil)
	fmt.Println("Done.")
}

func goroutineLeak() {
	memLeak()
}
