package firstclassfunction

import "fmt"

func simpleTest() {
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)
}

type add func(a int, b int) int

func testDefinedType() {
	var a add = func(a int, b int) int {
		return a + b
	}

	s := a(5, 6)
	fmt.Println("Sum", s)
}

// 高阶函数至少满足以下一个条件：
// 1. 有一个或多个函数作为参数
// 2. 返回一个结果为函数的值

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 6))
}

func testFuncAsParam() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

func simple2() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func testFuncAsReturn() {
	s := simple2()
	fmt.Println(s(60, 7))
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func testClosuer() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

func init() {
	fmt.Println()
	fmt.Println("===> enter firstclassfunction package")

	// testDefinedType()

	// testFuncAsParam()
	// testFuncAsReturn()

	// testClosuer()

	testPractice()

	fmt.Println("<=== exit firstclassfunction package")
	fmt.Println()
}
