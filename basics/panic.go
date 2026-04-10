package main

import "fmt"

func main() {

	//panic(interface{})

	//Example of a valid input
	process(10)

	//Example of a invalid input
	process(-1)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.Println("After Pasnic")
		// defer fmt.Println("Deferred 3")
	}
	fmt.Println("Processing input:", input)
}
