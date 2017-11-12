package main

import "fmt"
import "time"

// select 语句使得一个 goroutine 在多个通讯操作上等待。
// select 会阻塞，直到条件分支中的某个可以继续执行，这时就会执行那个条件分支。
// 当多个都准备好的时候，会随机选择一个。
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func customTimeout(c1, c2 <-chan string) {
	select {
	case msg1 := <-c1:
		fmt.Println("Message 1", msg1)
	case msg2 := <-c2:
		fmt.Println("Message 2", msg2)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	default:
		fmt.Println("nothing ready")
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		c1 <- "Yo"
	}()
	customTimeout(c1, c2)

	// just wait for code execution complete
	var input string
	fmt.Scanln(&input)
}
