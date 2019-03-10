package intf

import "fmt"

// Type Assertion
// 语法：i.(T)
// 获取接口 i 的具体类型（concrete type）T 的值
func assert(i interface{}) {
	s, ok := i.(int) // get the underlying int value from i
	fmt.Println(s, ok)
}

func testAssert() {
	// var s interface{} = 56
	var s interface{} = "Ping"
	assert(s)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
