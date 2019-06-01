package concurrency

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

func simpleCase() {
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

func block() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		// 此处阻塞，导致 c 的 reader 被阻塞
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	// 被阻塞
	case <-c:
		fmt.Printf("Unblocked %v later. \n", time.Since(start))
	}

	// 引发几个问题：
	// 1. 当有多个 channel 需要读，怎么办
	// 2. 当没有 channel 准备好时，怎么办
	// 3. 在当前时刻没有 channel 准备好，但是又要去执行点什么，怎么办

	// A1. go 的运行时会对多个 case 的读采用一种平均随机的算法来调度，也就是每个读都有平局的执行机会（multiSelect）
	// A2. 可以在最后一个 case 语句加入超时
	// A3. 在 select 中加入 default
}

func multiSelect() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func timeoutSelect() {
	var c <-chan int32

	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}

func defaultSelect() {
	start := time.Now()
	var c1, c2 <-chan int

	select {
	case <-c1:
	case <-c2:
	default:
		fmt.Printf("In default after %v\n", time.Since(start))
	}
}

func forSelect() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

func main() {
	// simpleCase()
	// block()
	// multiSelect()
	// timeoutSelect()
	// defaultSelect()
	forSelect()
}
