package deferpkg

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
}

func testDefferedMethod() {
	p := person{"John", "Smith"}
	defer p.fullName()
	fmt.Println("Welcome ")
}

func printA(a int) {
	fmt.Println("value of a in deferred function", a)
}

func testArgEvaluation() {
	a := 5
	defer printA(a)

	a = 10
	fmt.Println("value of a before deferred function call", a)
}

// 多个 defer 调用时，执行的时 LIFO（Last In First Out）队列
func testStackOfDefers() {
	name := "Naveen"
	fmt.Printf("Orignal String: %s\n", string(name))
	fmt.Printf("Reversed String: ")

	// 最后输出
	defer fmt.Printf("\n")

	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
