package interfaces

import (
	"fmt"
)

type describer interface {
	Describe()
}

type person struct {
	name string
	age  int
}

func (p person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

// 对象的类型与接口进行比较
func findType2(i interface{}) {
	switch v := i.(type) {
	case describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}
