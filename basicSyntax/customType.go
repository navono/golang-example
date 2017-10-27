package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

// 结构体文法
// 表示通过结构体字段的值作为列表来新分配一个结构体
var (
	v1 = Vertex{1, 2}  // 类型为Vertex
	v2 = Vertex{X: 1}  // Y: 0 被省略
	v3 = Vertex{}      // X:0 Y:0
	p1 = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	v := Vertex{1, 3}
	fmt.Println(v)

	v.X = 10
	fmt.Println(v)

	p := &v
	p.Y = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, p1)

	// 数组
	var a [2]string
	a[1] = "Yo"
	fmt.Println(a[0], a[1])
}
