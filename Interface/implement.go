package intf

import "fmt"

type describer interface {
	Describe()
}

type address struct {
	state   string
	country string
}

// implemented using pointer receiver
func (a *address) Describe() {
	fmt.Printf("\nState %s Country %s\n", a.state, a.country)
}

func testImplements() {
	var d1 describer
	p1 := person{"Sam", 25}
	d1 = p1

	p2 := person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 describer
	a := address{"Washington", "USA"}

	/* compilation error if the following line is
	   uncommented
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a // This works since Describer interface
	// is implemented by Address pointer in line 22
	d2.Describe()
}
