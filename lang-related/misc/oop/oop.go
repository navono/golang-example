package oop

import (
	"fmt"
	employee2 "golang_example/lang-related/misc/oop/employee"
)

func init() {
	fmt.Println()
	fmt.Println("===> enter oop package")

	// e := employee.Employee{
	// 	FirstName:   "Sam",
	// 	LastName:    "Adolf",
	// 	TotalLeaves: 30,
	// 	LeavesTaken: 20,
	// }
	e := employee2.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()

	fmt.Println("<=== exit oop package")
	fmt.Println()
}
