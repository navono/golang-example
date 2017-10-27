package main

import (
	"fmt"
)

// 指针保存了变量的内存地址
// Go 没有指针运算

// & 符号会生成一个指向其作用对象的指针。
// * 符号表示指针指向的底层的值。

func simpleTest() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func main() {
	simpleTest()
}
