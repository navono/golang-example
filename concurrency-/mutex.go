package concurrency

import (
	"fmt"
	"sync"
)

var x = 0

func unsafeIncrement(wg *sync.WaitGroup) {
	x++
	wg.Done()
}

func testRaceCondition() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go unsafeIncrement(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

var y = 0

func safeIncrement(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	y++
	wg.Done()
}

func testMutext() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go safeIncrement(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of y", y)
}

var z = 0

func safeIncrementWithChan(wg *sync.WaitGroup, ch chan bool) {
	// 因为 ch 是容量为 1 的缓存 chan，所以当其他 goroutine 试图去写时，
	// 会被阻塞。所有也就只一次允许 1 个 goroutine 对 z 进行写。
	ch <- true
	z++

	// 对 ch 进行读，这也就是会释放出 ch 的空间，让其他 goroutine 进行写
	<-ch
	wg.Done()
}

func testRaceConditionWithChan() {
	var w sync.WaitGroup

	// 容量为 1 的 buffered chan
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go safeIncrementWithChan(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", z)
}

// Mutex vs Channels
// In general use channels when Goroutines need to communicate with each other
// and mutexes when only one Goroutine should access the critical section of code.

// 简单说就是需要 goroutine 间通信的，用 Channel；
// 只控制单 goroutine 访问临界区的，用 Mutex。
