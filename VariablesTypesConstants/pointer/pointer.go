package pointer

import "fmt"

// Go dose not support pointer arithmetic

// Do not pass a pointer to an array as a argument to a function.
// Use slice instead.
func modifyWithPointer(arr *[3]int) {
	(*arr)[0] = 90
	// arr[x] is shorthand for (*arr)[x].
	// So (*arr)[0] in the above program can be replaced by arr[0]
	// so equals to
	// arr[0] = 90
}

func modify(arr []int) {
	arr[0] = 90
}

func change(val *int) {
	*val = 55
}

func init() {
	fmt.Println()
	fmt.Println("===> enter pointer package")

	b := 233
	// get the address of b
	var a = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)

	*a++
	fmt.Println("new value of b is", b)

	change(a)
	fmt.Println("value of a after change function call is", *a)

	li := [3]int{89, 90, 91}
	fmt.Println("before modifyWithPointer, li", li)
	modifyWithPointer(&li)
	fmt.Println("after modifyWithPointer, li", li)

	li = [3]int{89, 90, 91}
	fmt.Println("before modify, li", li)
	modify(li[:])
	fmt.Println("after modify, li", li)

	fmt.Println("<=== exit pointer package")
	fmt.Println()
}
