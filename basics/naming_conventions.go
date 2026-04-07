package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg, MyStruct, MyInterface, MyEnum

	// camelCase
	// Eg, myStruct, myInterface, myEnum

	// snake_case
	// Eg, my_struct, my_interface, my_enum

	// kebab-case
	// Eg, my-struct, my-interface, my-enum

	// Structs, interfaces, enums

	// UPPERCASE
	// Use case is Constants

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)
}
