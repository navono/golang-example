package concurrency

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func testSelect() {
	output1 := make(chan string)
	output2 := make(chan string)

	go server1(output1)
	go server2(output2)

	for {
		select {
		case s1 := <-output1:
			fmt.Println(s1)
			return
		case s2 := <-output2:
			fmt.Println(s2)
		// 进入 default 分支，然后直接退出，因此应该加上 for
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("no value received")
		}
	}
}
