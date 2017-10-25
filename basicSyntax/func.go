package main // 程序运行的入口

// 打包导入语句
import (
	"fmt"
)

func main() {
	// 首字母大写的名称是被导出的
	fmt.Println("Hello")
	fmt.Println(add(22, 33))

	b, a := swap("hello", "world")
	fmt.Println(b, a)

	fmt.Println(split(27))
}

// 函数可以没有参数或者接受多个参数，类型在变量名之后
func add(x int, y int) int {
	return x + y
}

// 同一类型可以写成这样
func add2(x, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, int) {
	return y + x, 5
}

// 命名返回。没有参数的`return`语句返回各个返回变量的当前值，称之为`裸`返回
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
