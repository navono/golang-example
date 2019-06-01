package intf

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// implemented using value receiver
func (p person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

// 对象的类型与接口进行比较
func findType2(i interface{}) {
	switch v := i.(type) {
	case describer:
		Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}
