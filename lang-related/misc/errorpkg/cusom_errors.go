package errorpkg

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// return 0, errors.New("Area calculation failed, radius is less than zero")

		// Or we can use Errorf
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}

	return math.Pi * radius * radius, nil
}

func testCustomError() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

// user struct type
type areaError struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

// old version
// func circleAreaWithTypeError(radius float64) (float64, error) {
// 	if radius < 0 {
// 		return 0, &areaError{"radius is negative", radius}
// 	}
// 	return math.Pi * radius * radius, nil
// }

// func testTypedError() {

// 	radius := -20.0
// 	area, err := circleAreaWithTypeError(radius)
// 	if err != nil {
// 		if err, ok := err.(*areaError); ok {
// 			fmt.Printf("Radius %0.2f is less than zero\n", err.radius)
// 			return
// 		}
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Area of rectangle1 %0.2f\n", area)
// }

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err, length, width}
	}
	return length * width, nil
}

func testTypedError() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}
