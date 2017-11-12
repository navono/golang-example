package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// `goroutine` 是由 Go 运行时环境管理的轻量级线程。并且不会等待函数执行完。
// main 函数本身也是一个`goroutine`。

// `goroutine` 在相同的地址空间中运行，因此访问共享内存必须进行同步。
// sync 提供了这种可能，不过在 Go 中并不经常用到，因为有其他的办法。
func main() {
	// say 和 world 是当前 `goroutine` 中定义的，但是在新的 `goroutine` 中运行 say。
	go say("world")
	say("hello")
}
