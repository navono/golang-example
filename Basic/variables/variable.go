package variables

import (
	"fmt"
)

func init() {
	fmt.Println()
	fmt.Println("===> enter variable package")

	var age int // declaration
	fmt.Println("declare age: ", age)

	age = 29 // assignment
	fmt.Println("assign age: ", age)

	var age2 int = 29 // variable declaration with initial value
	fmt.Println("declare with intial age: ", age2)

	var age3 = 29
	fmt.Println("inferred age: ", age3)

	var width, height int = 100, 50 //declaring multiple variables
	// or var width, height = 100, 50
	fmt.Println("multiple variables declaring: width is", width, ",height is", height)

	// different types in a single statement
	var (
		name   = "ping"
		gender string
	)
	fmt.Println("different types in a single statement: name:", name, ",gender: ", gender)

	golang, ver := "Go", 2
	fmt.Println("short hand declaration: lang:", golang, ",ver:", ver)

	fmt.Println("<=== exit variable package")
	fmt.Println()
}
