package main

import (
	"fmt"
	"math/cmplx"
)

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte  // uint8 的别名

// rune  // int32 的别名

// float32 float64

// complex64 complex128

// int，uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// 变量在定义时没有明确的初始化时会赋值为 零值 。

	// 零值是：

	// 数值类型为 0 ，
	// 布尔类型为 false ，
	// 字符串为 "" （空字符串）。
	var i int
	var fl float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, fl, b, s)

	// 表达式 T(v) 将值 v 转换为类型 T 。
	k := 43.34
	l := int(k)
	fmt.Printf("%v %v\n", k, l)

	// 类型推导
	m := l
	fmt.Printf("%T(%v)\n", m, m)
}

// 常量的定义与变量类似，只不过使用 const 关键字。
// 常量可以是字符、字符串、布尔或数字类型的值。
// 常量不能使用 := 语法定义。
const Pi = 3.14
