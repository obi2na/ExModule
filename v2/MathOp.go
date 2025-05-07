package ExModule

import "fmt"

// This add function just does what it says
// it adds to values and returns the result of the addition
func Add[T Number](a, b T) T {
	fmt.Println("successs")
	return a + b
}

type Number interface {
	int8 | int16 | int32 | int64 | int | uint8 |
		uint16 | uint32 | uint64 | uint | float32 | float64
}
