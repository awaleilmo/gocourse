package main

import "fmt"

func main() {

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	for _, v := range myMap {
		fmt.Println(v)
	}

	for k := range myMap {
		fmt.Println(k)
	}

	message := "Hello World!"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Character: %c\n", i, v)
	}

}
