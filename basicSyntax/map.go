package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// k-v容器
var m map[string]Vertex

func main() {
	// 需要用`make`来创建
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
