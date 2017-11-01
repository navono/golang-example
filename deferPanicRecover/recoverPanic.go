package main

import "fmt"

// Panic:

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		// `recover`只在 `defer`中才起作用
		// 对于正常的执行流，调用`recover`不会有任何副作用，`recover`会返回`nil`。
		// 如果当前的`goroutine`是`panic`状态，调用`recover`会接收到`panic`的值，并
		// 恢复到正常的执行流
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		// 停止当前的正常执行流，同时变为`panicking`。
		// 但是`defer`函数会被正常执行。
		// 然后返回到调用方。如果调用方存在`defer`函数，且有`recover`，
		// 那么`panic`则会停止，如果没有则崩溃。
		// 可注释掉函数`f`中的`defer`查看。
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in g ", i)
	fmt.Println("Printing in g ", i)
	g(i + 1)
}
