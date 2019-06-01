package employee

import (
	"fmt"
)

// Struct Instead of Classes

// Employee represent a employee
// type Employee struct {
// 	FirstName   string
// 	LastName    string
// 	TotalLeaves int
// 	LeavesTaken int
// }

// Change to lowercase
type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

// LeavesRemaining calculate and displays the number of remaining leaves
func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, e.TotalLeaves-e.LeavesTaken)
}

// New to substitute constructor
func New(firstName, lastName string, totalLeave, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}
