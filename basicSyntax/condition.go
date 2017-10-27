package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// 后置，也可用作 while 使用
	sum2 := 1
	for sum2 < 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Println("%s.", os)
	}

	// defer 会延迟函数的执行，直到上层函数返回
	defer fmt.Println("Yo")
	fmt.Println("What's up.")

	inverse()
}

func pow(x, n, lim float64) float64 {
	// v 是局部作用域（if语句内）的变量
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func inverse() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
