package main

import "fmt"

func main() {

	// ... ellipsis
	// func functionName(param1,type1, param2, type2, param3 ...type3) returnType{
	// function body
	// }

	statement, total := sum("The sum of 1,2,3,4,5 is", 1, 2, 3, 4, 5)
	fmt.Println(statement, total)

	sequence, total := sum1(1, 2, 3, 4, 5)
	fmt.Println("Sequence :", sequence, "Total :", total)

	numbers := []int{1, 2, 3, 4, 5}
	sequence, total = sum1(3, numbers...)
	fmt.Println("Sequence :", sequence, "Total :", total)
}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}

func sum1(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
