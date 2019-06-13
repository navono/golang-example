package dump

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	//"runtime/pprof"
	"syscall"
	"time"
)

// run in root with:
// go run main.go

// show all goroutine info:
// GOTRACEBACK=1 go run main.go
// none、all、system、single、crash，历史原因， 可以设置数字0、1、2，分别代表none、all、system

func init() {
	setupSigusr1Trap()
	go a()
	m1()
}

func m1() {
	m2()
}

func m2() {
	m3()
}

func m3() {
	//debug.PrintStack()
	// or
	//pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

	//panic("panic from m3")
	time.Sleep(time.Hour)
}

func a() {
	time.Sleep(time.Hour)
}

func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		for range c {
			dumpStacks()
		}
	}()
}

func dumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
