package arrayslices

import "fmt"

// 数组的大小是类型的一部分
// 数组时值类型，数组作为函数的参数传入时，也是作为值类型

// slice 是数组的一个封装，它更便捷、更灵活、更强大；slice 并不直接拥有数据，
// 它只是数组的一个引用

// array 是固定长度的数组，使用前必须确定数组长度
// slice 是不定长的

// 声明数组时，方括号内写明了数组的长度或者 ...,声明 slice 时候，方括号内为空
// 作为函数参数时，数组传递的是数组的副本，而 slice 传递的是指针

func init() {
	fmt.Println()
	fmt.Println("===> enter arrayslices package")

	var a [3]int // int array with length
	a[0] = 12

	fmt.Println("int array with length 3. a:", a)

	b := [2]int{12, 3}
	fmt.Println("short hand declaration. b:", b)

	c := [...]int{30, 51, 52}
	fmt.Println("compiler determine the length. c:", c)

	// fmt.Println("The size of the array is a part of the type. b == c", b == c)

	d := [...]string{"USA", "China", "Germany", "Canada"}
	dCopy := d // d copy of a is assigned to dCopy
	dCopy[0] = "Singapore"
	fmt.Println("Arrays are value types. d:", d)
	fmt.Println("Arrays are value types. dCopy:", dCopy)

	fmt.Println("length of d:", len(d))
	fmt.Println("capacity of d:", cap(d))

	fmt.Println("use range to iterate d:")
	for i, v := range d { //range returns both the index and value
		fmt.Printf("%d the element of d is %v\n", i, v)
	}

	e := [5]int{76, 77, 78, 79, 80}
	f := e[1:4] //creates a slice from e[1] to e[3]
	fmt.Println("original array e:", e)
	fmt.Println("creates a slice from e[1] to e[4], [1, 4):", f)

	fmt.Println("length of f:", len(f))
	fmt.Println("capacity of f:", cap(f))

	g := make([]int, 5, 5)
	fmt.Println("create slice using make. g:", g)

	g = append(g, 3)
	fmt.Println("append a new element to g:", g, "len:", len(g), "cap:", cap(g))
	h := []int{100, 200}
	g = append(g, h...)
	fmt.Println("append another array(h) to g:", g, "len:", len(g), "cap:", cap(g))

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

	fmt.Println("<=== exit arrayslices package")
	fmt.Println()
}

func countries() []string {
	// Memory Optimisation
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))

	// 拷贝需要的元素，原始的数组就可以被回收
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
