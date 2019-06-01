package concurrency

import (
	"fmt"
	"time"
)

func randomServer1(ch chan string) {
	ch <- "from random server1"
}

func randomServer2(ch chan string) {
	ch <- "from random server2"
}

func testRandomSelect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go randomServer1(output1)
	go randomServer2(output2)
	time.Sleep(1 * time.Second)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
