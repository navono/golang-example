package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// 向带缓冲的 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。
// 而当缓冲区为空的时候接收操作会阻塞。

func simpleCase() {
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

func bufCase() {
	var stdoutBuf bytes.Buffer
	defer stdoutBuf.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuf, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuf, "Sending: %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuf, "Received %v. \n", integer)
	}
}

func main() {
	// bufCase()
	ownership()
}

// channel 的所有权，即初始化、写、关闭channel的 goroutine。一般可以从单向的 channel 来区分。
// 拥有 channle 所有权的 goroutine 一般拥有`写`权限，也就是（chan or chan<-）。

// 因此，一个拥有 channle 所有权的 goroutine 应该：
// 1. 初始化这个 channel
// 2. 执行写，或者将所有权转移给其他的 goroutine
// 3. 关闭这个 channel
// 4. 将上述的3条压缩，且通过 reader channel暴露

// 执行以上的约束，也就意味着以下的事实：
// 1. 因为初始化了 channel，所以就不会向一个 nil 的 channel执行写操作而导致死锁
// 2. 因为初始化了 channel，所以就不会关闭一个 nil 的channel而导致 panic
// 3. 因为拥有了 channel 的所有权，所以也就决定了 channel 的 close 操作，所以也就不会
//    对一个 channel 进行多次 close 而导致的 panic
// 4. 在编译期可以进行类型检测，用来防止不当的写操作

// 此时，对于 channel 的 reader 来说，就需要关心两件事：
// 1. 得知道 channel 什么时候被关闭
// 2. 处理 blocking

// 第一点可以通过读操作的返回值来确定。第二点则要根据具体的算法和场景决定。

func ownership() {
	chanOwner := func() <-chan int {
		// 初始化
		resultStream := make(chan int, 5)
		// 将写操作封装到一个 goroutine 中
		go func() {
			// 确保 channel 被正确 close
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		// 返回一个只读的 channel
		return resultStream
	}

	// resultStream的生命周期都被封装到了 chanOwner 函数中
	
	resultStream := chanOwner()
	// 对于 channel 的reader，只需要关心 channel 的 blocking 和 closed
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}

	fmt.Println("Done receiving!")
}
