package intf

import "fmt"

// VowelsFinder interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

// MyString custom string type
type MyString string

// FindVowels MyString implements
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

type salaryCalculator interface {
	CalculateSalary() int
}

type permanent struct {
	empID    int
	basicpay int
	pf       int
}

type contract struct {
	empID    int
	basicpay int
}

// CalculateSalary returns salary of permanent employee is sum of basic pay and pf
func (p permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// CalculateSalary returns salary of contract employee is the basic pay alone
func (c contract) CalculateSalary() int {
	return c.basicpay
}

/*
	total expense is calculated by iterating though the SalaryCalculator slice and summing
	the salaries of the individual employees
*/
func totalExpense(s []salaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n", expense)
}

func testPractiseInterface() {
	pemp1 := permanent{1, 5000, 20}
	pemp2 := permanent{2, 6000, 30}
	cemp1 := contract{3, 3000}

	employees := []salaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}

func init() {
	fmt.Println()
	fmt.Println("===> enter interfaces package")

	name := MyString("Sam Anderson")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	testPractiseInterface()

	testAssert()

	findType("Na")
	findType(77)
	findType(77.2)

	p := person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)

	testImplements()
	testEmbededInterface()

	fmt.Println("<=== exit interfaces package")
	fmt.Println()
}
