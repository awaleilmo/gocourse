package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22.0 / 7.0
	fmt.Println("Pi: ", p)

	// Overflow with signed integers
	var maxInt int64 = 9223372036854775807 // max value of int64
	fmt.Println("Max Int64: ", maxInt)

	maxInt = maxInt + 1
	fmt.Println("Max Int64 + 1: ", maxInt)

	// Overflow with unsigned integers
	var maxUint uint64 = 18446744073709551615 // max value of uint64
	fmt.Println("Max Uint64: ", maxUint)

	maxUint = maxUint + 1
	fmt.Println("Max Uint64 + 1: ", maxUint)

	// Underflow with floating point numbers
	var smallFLoat float64 = 1.0e-323
	fmt.Println("Small Float: ", smallFLoat)

	smallFLoat = smallFLoat / math.MaxFloat64
	fmt.Println("Small Float: ", smallFLoat)
}
