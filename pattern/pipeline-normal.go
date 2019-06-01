package pattern

import (
	"fmt"
)

func pipelineNormal() {
	multiply := func(values []int, multiplier int) []int {
		multipliedValues := make([]int, len(values))
		for i, v := range values {
			multipliedValues[i] = v * multiplier
		}
		return multipliedValues
	}

	add := func(values []int, additive int) []int {
		addedValues := make([]int, len(values))
		for i, v := range values {
			addedValues[i] = v + additive
		}
		return addedValues
	}

	number := []int{1, 2, 3, 4}
	for _, v := range add(multiply(number, 2), 1) {
		fmt.Println(v)
	}
}
