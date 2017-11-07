package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go 没有类。然而，仍然可以在结构体类型上定义方法。
// 方法接收者 出现在 func 关键字和方法名之间的参数中。

// 方法可以与命名类型或命名类型的指针关联
// 有两个原因需要使用指针接收者。
// 首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
// 其次，方法可以修改接收者指向的值。
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// 可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	a := Android{
		Model:  "m",
		Person: struct{ Name string }{Name: "t"},
	}
	a.Person.Talk()

	a1 := &Android{
		Model: "m",
		Person: Person{
			Name: "you",
		},
	}
	a1.Person.Talk()

	b := new(Android)
	b.Talk()
}

// Embedded Types
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// is-a releationship
// Android is a Person
type Android struct {
	Person
	Model string
}
