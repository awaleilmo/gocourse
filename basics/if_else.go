package main

import "fmt"

func main() {

	// if condition {
	// 	block of code
	// }

	// age := 25
	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// }

	// if condition {
	// 	block of code
	// } else if condition {
	// 	block of code
	// } else {
	// 	block of code
	// }

	// temperature := 25
	// if temperature > 30 {
	// 	fmt.Println("It's hot outside")
	// } else if temperature > 20 {
	// 	fmt.Println("It's warm outside")
	// } else {
	// 	fmt.Println("It's cold outside")
	// }

	// score := 85

	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade C")
	// } else {
	// 	fmt.Println("Grade F")
	// }

	// if condition1 {
	// 	block of code
	// 	if condition2 {
	// 		block of code
	// 	} else if condition3 {
	// 		block of code
	// } else {
	// 	block of code
	// }

	// num := 18
	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println("Number is divisible by both 2 and 3.")
	// 	} else {
	// 		fmt.Println("Number is divisible by 2 but not 3.")
	// 	}
	// } else {
	// 	fmt.Println("Number is not divisible by 2.")
	// }

	// || or
	// && and

	if 10%2 == 0 && 10%5 == 0 {
		fmt.Println("Number is divisible by both 2 and 5.")
	}

	if 10%2 == 0 || 10%5 == 0 {
		fmt.Println("Number is divisible by both 2 and 5.")
	}

}
