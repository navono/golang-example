package main

import "fmt"

// channel 是有类型的管道，可以用 channel 操作符 <- 对其发送或者接收值。

// 函数 sum 中的 c 是一个双通道的 channel，也就是能发能收。
// 可以通过在声明时，指定参数的 channel的接受方式，
// 比如 c chan<- string：只允许往 c 发送数据，而不能接收；
// c <-chan string：只允许 c 接收数据

func sum(a []int, c chan<- int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	// “箭头”就是数据流的方向。
	c <- sum // 将和送入 c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	// 和 map 与 slice 一样，channel 使用前必须创建：
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	// 默认情况下，在另一端准备好之前，发送和接收都会阻塞。
	// 这使得 goroutine 可以在没有明确的锁或竞态变量的情况下进行同步。
	x, y := <-c, <-c // 从 c 中获取

	fmt.Println(x, y, x+y)
}
