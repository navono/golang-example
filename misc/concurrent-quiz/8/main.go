package main

import "fmt"

func main() {
	var ch chan int
	var count int

	go func() {
		ch <- 1
	}()

	go func() {
		count++
		// ch 未初始化
		close(ch)
	}()

	<-ch

	fmt.Println(count)
}
