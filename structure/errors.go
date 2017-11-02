package main

import (
	"fmt"
	"time"
)

// Go 程序使用 error 值来表示错误状态。
// 与 fmt.Stringer 类似， error 类型是一个内建接口
// type error interface {
// 	Error() string
// }

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
