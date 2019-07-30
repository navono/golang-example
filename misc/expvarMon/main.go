// 可以编译，运行时内存可能暴涨

package main

import (
	"bytes"
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var pool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func main() {
	go func() {
		for {
			processRequest(1 << 28) // 256MiB
		}
	}()
	for i := 0; i < 10; i++ {
		go func() {
			for {
				processRequest(1 << 10) // 1KiB
			}
		}()
	}

	// for expVarMon
	go func() {
		log.Fatal(http.ListenAndServe(":1234", nil))
	}()

	var stats runtime.MemStats
	for i := 0; ; i++ {
		runtime.ReadMemStats(&stats)
		fmt.Printf("Cycle %d: %dB\n", i, stats.Alloc)
		time.Sleep(time.Second)
		runtime.GC()
	}
}

func processRequest(size int) {
	b := pool.Get().(*bytes.Buffer)
	time.Sleep(500 * time.Millisecond)
	b.Grow(size)
	pool.Put(b)
	time.Sleep(1 * time.Millisecond)
}
