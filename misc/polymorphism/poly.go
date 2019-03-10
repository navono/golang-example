package polymorphism

import (
	"fmt"
)

// Income represent the imaginary organisation incom from projects
type Income interface {
	calculate() int
	source() string
}

// FixedBilling represent a project
type FixedBilling struct {
	projectName  string
	biddedAmount int
}

// TimeAndMaterial represent another project
type TimeAndMaterial struct {
	projectName string
	noOfHours   int
	hourlyRate  int
}

func (fb FixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling) source() string {
	return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
	return tm.projectName
}

func calculateNetIncome(ic []Income) {
	var netIncome = 0
	for _, income := range ic {
		fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
		netIncome += income.calculate()
	}
	fmt.Printf("Net income of organisation = $%d\n", netIncome)
}

func init() {
	fmt.Println()
	fmt.Println("===> enter polymorphism package")

	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	incomeStreams := []Income{project1, project2, project3}
	calculateNetIncome(incomeStreams)

	fmt.Println("<=== exit polymorphism package")
	fmt.Println()
}
