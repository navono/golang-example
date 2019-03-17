package deferpkg

import "fmt"

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) {
	defer finished()

	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func init() {
	fmt.Println()
	fmt.Println("===> enter defer package")

	// nums := []int{78, 109, 2, 563, 300}
	// largest(nums)

	// testDefferedMethod()

	// testArgEvaluation()

	testStackOfDefers()

	fmt.Println("<=== exit defer package")
	fmt.Println()
}
