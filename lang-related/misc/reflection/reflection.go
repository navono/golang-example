package reflection

import (
	"fmt"
	"reflect"
)

type order struct {
	ordID    int
	customID int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

// func createQuery(o order) string {
// 	i := fmt.Sprintf("insert into order values(%d, %d)",
// 		o.ordID, o.customID)
// 	return i
// }

func createQuery(q interface{}) {
	// t := reflect.TypeOf(q)
	// k := t.Kind()
	// v := reflect.ValueOf(q)

	// fmt.Println("Type ", t)
	// fmt.Println("Kind ", k)
	// fmt.Println("Value ", v)

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func testCreateQuery() {
	o := order{1234, 567}
	// fmt.Println(createQuery(o))
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)
}

func init() {
	fmt.Println()
	fmt.Println("===> enter reflection package")

	testCreateQuery()

	fmt.Println("<=== exit reflection package")
	fmt.Println()
}
