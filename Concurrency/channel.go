package concurrency

import (
	"fmt"
	"time"
)

func helloGoroutine() {
	fmt.Println("Hello goroutine")
}

func helloGoroutine2(done chan bool) {
	fmt.Println("Hello goroutine")
	done <- true
}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func testHello() {
	done := make(chan bool)
	go hello(done)

	// 阻塞等待
	<-done
}

func hello2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func testHello2() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello2(done)
	<-done
	fmt.Println("Main received data")
}

// deadlock
// 如果一个 goroutine 通过 channel 发送了数据，
// 那么它就会假设其他的 goroutine 会通过 channel 接收数据，
// 否则就会死锁
func deadLock() {
	ch := make(chan int)
	ch <- 5
}

func init() {
	fmt.Println()
	fmt.Println("===> enter concurrency package")

	// 没有任何输出，因为 goroutine 是异步执行的，创建 goroutine 后
	// 会立马返回，执行后面的代码
	// go helloGoroutine()

	// done := make(chan bool)
	// go helloGoroutine2(done)
	// // channel 的读与写默认都是阻塞的，所以此处会等待 goroutine 执行完
	// <-done

	// testHello2()
	// deadLock()

	// testIterateChan()
	// testRangeChan()

	// testCalc()

	// testBufferdChan()
	// testWaitGroup()
	// testWorkerPool()

	// testSelect()
	// testRandomSelect()

	testRaceCondition()
	testMutext()
	testRaceConditionWithChan()

	fmt.Println("<=== exit concurrency package")
	fmt.Println()
}
