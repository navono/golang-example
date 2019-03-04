package intf

import "fmt"

type salaryDisplayer interface {
	DisplaySalary()
}

type leaveCalculator interface {
	CalculateLeavesLeft() int
}

type employOperations interface {
	salaryDisplayer
	leaveCalculator
}

type employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func testEmbededInterface() {
	e := employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp employOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}
