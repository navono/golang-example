package main

import (
	"fmt"
	"io"
	"strings"
)

// io 包指定了 io.Reader 接口， 它表示从数据流结尾读取。

func main() {
	r := strings.NewReader("Hello, Reader!")

	// 每次 8 字节的速度读取它的输出
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
