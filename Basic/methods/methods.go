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
	(&emp2).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d\n", emp2.age)

	fmt.Println("<=== exit methods package")
	fmt.Println()
}
