package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringer 是一个可以用字符串描述自己的类型。
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
