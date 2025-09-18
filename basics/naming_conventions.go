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
	// pascal case
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// structs, interface, enums

	// snake_case
	// Eg. calculate_area, user_info, new_http_request

	// UPPERCASE
	// Use case is Constants

	// mixedCase
	// Eg. javaScript, htmlDocument. isValid

	const MAXRETRIES = 5
	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)

}
