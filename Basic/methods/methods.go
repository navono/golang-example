package method

import "fmt"

// methods 是一个具有特殊接收器类型（receiver）的函数，接收器类型可以是 struct 类型也可以使 non-struct 类型，
// 接收器（receiver）在 method 内可以被访问

// func (t Type) methodName(parameter list) {
// }

// 有了 func 还需要 method 的理由是：
//  1. Go 不是纯面向对象的语言，也没有类的概念，类型上附带 method 可以达到类似类的行为
//  2. 相同名字的 method 可以定义在不同的类型上，而函数做不到

type employee struct {
	name     string
	age      int
	salary   int
	currency string
}

// e 就是 receiver
func (e employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
	fmt.Print("\n")
}

// pointer receivers vs value receivers
// 简单说就是指针类型的接收器在改变接收器数据后，是对调用者生效的

/*
Method with value receiver
*/
func (e employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *employee) changeAge(newAge int) {
	e.age = newAge
}

// 匿名字段的方法调用
type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("\nFull address: %s, %s\n", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

func testAnounymouseFieldMethod() {
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress()
}

// method 中的值接收器（value receivers）vs 函数中的值参数（value arguments）
//   1. 当函数有一个值参数，那么它将只能接受值参数
//   2. 当方法有一个值接收器，那么它可以接受值和指针的接收器

// 上述规则对于指针类型也适用

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func testValueArgsAndValueReceivers() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	 compilation error, cannot use p (type *rectangle) as type rectangle
	 in argument to area
	*/
	// area(p)

	p.area() //calling value receiver with a pointer
}

// method 作用于 non-struct 上时，要求 method 的接收器类型的定义与 method 的定义
// 在同一个包

// 非法，因为 int 不在 main 包
// package main
// func (a int) add(b int) {
// }

// 合法
// package main
// import "fmt"
// type myInt int
// func (a myInt) add(b myInt) myInt {
//     return a + b
// }

func init() {
	fmt.Println()
	fmt.Println("===> enter methods package")

	emp1 := employee{
		name:     "ping",
		salary:   5000,
		currency: "$",
	}

	emp1.displaySalary()

	emp2 := employee{
		name: "ping",
		age:  30,
	}

	fmt.Printf("Employee name before change: %s", emp2.name)
	emp2.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", emp2.name)

	fmt.Printf("\n\nEmployee age before change: %d", emp2.age)
	// (&emp2).changeAge(51)
	emp2.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d\n", emp2.age)

	testAnounymouseFieldMethod()
	testValueArgsAndValueReceivers()

	fmt.Println("<=== exit methods package")
	fmt.Println()
}
