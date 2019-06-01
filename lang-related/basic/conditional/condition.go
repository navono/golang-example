package condition

import "fmt"

func number() int {
	num := 15 * 5
	return num
}

func init() {
	fmt.Println()
	fmt.Println("===> enter condition package")

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}

	fmt.Println("<=== exit condition package")
	fmt.Println()
}
