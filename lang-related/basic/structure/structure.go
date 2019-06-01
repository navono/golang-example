package structure

import "fmt"

// Structs are value types and are comparable if each of their fields are comparable.
// Two struct variables are considered equal if their corresponding fields are equal.

// Struct variables are not comparable if they contain fields which are not comparable

// named structure
// type Employee struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// equals to
type employee struct {
	firstName, lastName string
	age, salary         int
}

// anonymous structures
var anonymousEmployee struct {
	firstName, lastName string
	age                 int
}

func structure() {
	//creating structure using field names
	emp1 := employee{
		firstName: "sam",
		age:       25,
		salary:    100,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// creating anonymous structure
	emp3 := struct {
		firstname, lastname string
		age, salary         int
	}{
		firstname: "Andreah",
		lastname:  "Nikola",
		age:       22,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	emp3.salary = 6000
	fmt.Println("Employee 3 after change field", emp3)

	// pointers to a struct
	emp4 := &employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 4", emp4)
	// (*emp4).salary = 1000
	// equals to
	emp4.salary = 1000
	fmt.Println("Employee 4 after change field", emp4)
}

type address struct {
	city, state string
}

// anonymous fields.
// default the name of a anonymous field is the name of its type
type person struct {
	string
	int
}

// nested struct
type person2 struct {
	name string
	age  int
	addr address
}

// promoted fileds
type person3 struct {
	name string
	age  int
	address
}

func init() {
	fmt.Println()
	fmt.Println("===> enter structure package")

	structure()

	p := person{"ping", 30}
	fmt.Println("Person", p)
	p.int = 31
	fmt.Println("Person after change field", p)

	var p2 person2
	p2.name = "ping"
	p2.age = 30
	p2.addr = address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Person2", p2)

	var p3 person3
	p3.name = "ping"
	p3.age = 31
	p3.address = address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Person3", p3)
	fmt.Println("Person3 city", p3.city)

	fmt.Println("<=== exit structure package")
	fmt.Println()
}
