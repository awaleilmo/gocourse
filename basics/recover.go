package main

import "fmt"

func main() {
	process()
	fmt.Println("End Process")
}

func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong")

}
