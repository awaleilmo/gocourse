package main

import "fmt"

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	arr[0] = 10
	fmt.Println(arr)

	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println(fruits)

	fmt.Println("The first fruit is", fruits[0])
	fmt.Println("The second fruit is", fruits[1])
	fmt.Println("The third fruit is", fruits[2])

	originalArr := [3]int{223, 11, 90}
	var copiedArr *[3]int

	copiedArr = &originalArr
	copiedArr[0] = 100
	fmt.Println("Original Array:", originalArr)
	fmt.Println("Copied Array:", copiedArr)

	for i := 0; i < len(originalArr); i++ {
		fmt.Println("Element at index", i, "is", originalArr[i])
	}

	for i, v := range originalArr {
		fmt.Printf("Index %d, Value %d\n", i, v)
	}

	for _, v := range originalArr {
		fmt.Printf("Value %d\n", v)
	}

	a, b := someFunction()
	fmt.Println(a, b)

	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println("Arrays1 is equal to Array2", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

}

func someFunction() (int, int) {
	return 1, 2
}
