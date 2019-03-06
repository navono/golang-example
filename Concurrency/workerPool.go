package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 核心功能：
// 1. 创建一个 goroutine 池，都监听一个输入类型的 buffered Channel，等待任务
// 2. 往输入类型的 buffered Channel 增加任务
// 3. 在任务完成后，等待输出类型的 buffered Channel 的结果

type job struct {
	id       int
	randomNO int
}

type result struct {
	job         job
	sumOfDigits int
}

var jobs = make(chan job, 10)
var results = make(chan result, 10)

func digits2(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := result{job, digits2(job.randomNO)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNO := rand.Intn(999)
		job := job{i, randomNO}
		jobs <- job
	}
	close(jobs)
}

func resultPrint(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n",
			result.job.id, result.job.randomNO, result.sumOfDigits)
	}
	done <- true
}

func testWorkerPool() {
	startTime := time.Now()

	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go resultPrint(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done

	endTime := time.Now()

	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
