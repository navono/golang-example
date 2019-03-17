package errorhandling

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
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s\n", e.radius, e.err)
}

func circleAreaWithTypeError(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func testTypedError() {

	radius := -20.0
	area, err := circleAreaWithTypeError(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero\n", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle1 %0.2f\n", area)
}
