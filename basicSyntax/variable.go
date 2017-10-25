package main

import (
	"fmt"
)

// var 定义了一个变量列表，类型在后。
// 可以定义在函数或包级别
var c, python, java bool

func main() {
	var i int

	fmt.Println(i, c, python, java)

	// 表达式初始化
	var l, m = true, "yse"
	fmt.Println(j, k, l, m)

	// 短声明 :=, 不能在函数外使用。函数外的每个语句都必须以关键字开始（var、func等等）
	n := 5
	fmt.Println(n)
}

// 变量初始化
var j, k int = 1, 3
