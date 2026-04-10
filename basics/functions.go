package main

import "fmt"

func main() {

	// fmt.Println(add(1, 2))

	// greet := func() {
	// 	fmt.Println("Hello World!")
	// }

	// greet()

	// operation := add

	// result := operation(10, 20)

	// fmt.Println(result)

	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println(result)

	// Returning and using a function
	multiplier := createMultiplier(9)
	fmt.Println("6 * 9 = ", multiplier(6))

}

func add(a int, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
