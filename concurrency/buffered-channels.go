package main

import "fmt"

// 向带缓冲的 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。
// 而当缓冲区为空的时候接收操作会阻塞。

func main() {
	// channel 可以是 带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
