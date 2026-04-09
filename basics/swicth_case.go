package main

import "fmt"

func main() {

	// Switch statment (default)
	// switch expression {
	// case value1:
	// 	// code to execute
	// case value2:
	// 	// code to execute
	// default:
	// 	// code to execute
	// }

	// Switch statment (fallthrough)
	// switch expression {
	// case value1:
	// 	// code to execute
	// 	fallthrough
	// case value2:
	// 	// code to execute
	// 	fallthrough
	// default:
	// 	// code to execute
	// }

	// Switch statment (multiple conditions)
	// switch expression {
	// case value1, value2, value3:
	// 	// code to execute
	// case value4, value5, value6:
	// 	// code to execute
	// default:
	// 	// code to execute
	// }

	// Switch statment (break)
	// switch expression {
	// case value1:
	// 	// code to execute
	// 	break
	// case value2:
	// 	// code to execute
	// 	break
	// default:
	// 	// code to execute
	// }

	// day := "Monday"

	// switch day {
	// case "Monday":
	// 	fmt.Println("It's Monday")
	// 	break
	// case "Tuesday":
	// 	fmt.Println("It's Tuesday")
	// 	break
	// case "Wednesday":
	// 	fmt.Println("It's Wednesday")
	// 	break
	// case "Thursday":
	// 	fmt.Println("It's Thursday")
	// 	break
	// case "Friday":
	// 	fmt.Println("It's Friday")
	// 	break
	// case "Saturday":
	// 	fmt.Println("It's Saturday")
	// 	break
	// case "Sunday":
	// 	fmt.Println("It's Sunday")
	// 	break
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number > 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 20")
	// case number > 20 && number < 30:
	// 	fmt.Println("Number is between 20 and 30")
	// default:
	// 	fmt.Println("Number is greater than 30")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float64")
	default:
		fmt.Println("Unknown")
	}
}
