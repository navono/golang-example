package main

import "fmt"

func main() {
	a()
	fmt.Println(b())
}

// `defer`函数的参数是在当`defer`语句被评估（evaluated）是计算的
func a() {
	i := 0
	// 输出 0
	defer fmt.Println(i)
	i++
	// 后进先出
	defer fmt.Println(i)

	return
}

func b() (j int) {
	// `defer`函数可能会读取并赋值给调用函数的命名返回值
	defer func() {
		// 因为 `return` 语句已经返回1，
		// 所以传到此的`j`为1
		// 最终函数`b`返回2
		j++
	}()
	return 1
}
