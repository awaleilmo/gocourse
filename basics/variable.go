package basics

import "fmt"

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "jhon"
	// var name1 = "jane"

	// count := 10
	// lastName := "Smith"

	fmt.Println(middleName)
	// Default values
	// Numeric Types: 0
	// String: ""
	// Boolean: false
	// Pointer: nil
	// Slice: nil
	// Map: nil
	// Channel: nil
	// Struct: zero value for each field
	// Array: zero value for each element
	// Interface: nil
	// Function: nil

	// ---- SCOPE

	// fmt.Println(firstName)

}

func printname() {
	firstname := "Michel"
	fmt.Println(firstname)
}
