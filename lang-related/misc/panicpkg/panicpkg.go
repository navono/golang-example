package panicpkg

import (
	"fmt"
	"runtime/debug"
)

// 使用 panic 的场景
// 1. 程序发生了不可恢复的故障，同时又不能简单地继续执行下去
//    比如说 web server 需要绑定的端口被占用了。
// 2. 程序级别的故障
//    比如一个指针类型的参数接收了一个 nil 的实参。

func fullName(firstName *string, lastName *string) {
	// defer fmt.Println("deferred call in fullName")
	defer recoverName()

	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

// Recover works only when it is called from the same goroutine.
// recover 也可以使用在运行时错误中，比如越界索引
func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)

		// 增加输出栈信息
		debug.PrintStack()
	}
}

func init() {
	defer fmt.Println("deferred call in panicpkg init")
	fmt.Println()
	fmt.Println("===> enter panic package")

	firstName := "Elon"
	fullName(&firstName, nil)

	fmt.Println("<=== exit panic package")
	fmt.Println()
}
