package main

import (
	"fmt"
	"time"
)

// 向带缓冲的 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。
// 而当缓冲区为空的时候接收操作会阻塞。

func main() {
	// channel 可以是 带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲
	bufferedChan := make(chan string, 3)
	doneChan := make(chan string)
	// var w sync.WaitGroup
	// w.Add(1)

	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()

	// main 函数也是个 goroutine
	<-time.After(time.Second * 1)

	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)

		// w.Done()
		doneChan <- "done!"
	}()

	// w.Wait()
	fmt.Println(<-doneChan)
}
