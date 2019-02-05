package types

import "fmt"

func init() {
	fmt.Println()
	fmt.Println("===> enter types package (map)")

	// map 是引用类型
	// 不能使用 == 进行比较，== 只能与 nil 进行判断

	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000

	// personSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// }
	// personSalary["mike"] = 9000

	fmt.Println("personSalary map contents:", personSalary)

	// value, ok := map[key]

	fmt.Println("All items of a map")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}

	fmt.Println("map before deletion", personSalary)
	delete(personSalary, "steve")
	fmt.Println("map after deletion", personSalary)

	fmt.Println("<=== exit types package (map)")
	fmt.Println()
}
