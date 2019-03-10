package concurrency

import "fmt"

func sendData(ch chan<- int) {
	ch <- 10
}

func testUnidirectionalChan() {
	// send := make(chan<- int)
	// go sendData(send)

	// 错误！
	// 因为 send 在声明时为 只写 属性
	// 因此不可读
	// fmt.Println(<-send)

	send := make(chan int)
	go sendData(send)
	fmt.Println(<-send)
}
