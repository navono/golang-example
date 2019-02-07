package types

import (
	"fmt"
	"unsafe"
)

// bool
// Numeric Types
// 	int8, int16, int32, int64, int
// 	uint8, uint16, uint32, uint64, uint
// 	float32, float64
// 	complex64, complex128
// 	byte: an alias of uint8
// 	rune: an alias of int32
// string

func init() {
	fmt.Println()
	fmt.Println("===> enter types package")

	var a = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))     //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d\n", b, unsafe.Sizeof(b)) //type and size of b

	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println("int converted: ", sum)

	k := 10
	var m = float64(k) //this statement will not work without explicit conversion
	fmt.Println("float converted:", m)

	fmt.Println("<=== exit types package")
	fmt.Println()
}
