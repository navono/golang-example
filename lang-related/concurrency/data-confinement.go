package concurrency

import (
	"fmt"
)

func adHocConfine() {
	data := make([]int, 4)

	// 约定只在 `loopData` 对 data 进行修改
	loopData := func(handleData chan<- int) {
		defer close(handleData)

		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}

func lexicalConfine() {
	// 返回的 channel 是只读的，用来限制
	chanOwner := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)
}
