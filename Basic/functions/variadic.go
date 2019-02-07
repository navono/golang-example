package functions

import (
	"fmt"
)

// 如果函数的最后一个参数被标记为 `...T`，那么此函数可以接收任意数量
// 类型为 T 的参数

func find(num int, nums ...int) {
	// 实际上变参是作为一个 slice 传入的
	// 例如 find(89, []int{})
	// 因为 find 的第二个参数是变参，实际传入的 slice，因此是引用关系
	// 改变其参数的值实际上会改变传入的参数的值

	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in", nums)
	}

	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func init() {
	fmt.Println()
	fmt.Println("===> enter functions package")

	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}
	// 如果直接传入 slice，则会导致错误，因为类型不正确
	// 应为 nums 已经是一个 slice
	// find(89, []int{nums})
	// find(89, nums)

	// 所以应该用以下语法糖方式传入一个 slice，这样就不用通过编译器再次创建
	// 新的 slice，而是直接以参数作为 slice
	find(89, nums...)

	welcom := []string{"hello", "world"}
	change(welcom...)
	fmt.Println(welcom)

	fmt.Println("<=== exit functions package")
	fmt.Println()
}
