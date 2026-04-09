package main

import (
	"fmt"
	"slices"
)

func main() {

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{9, 8, 7}

	// slice := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice1:", slice1)

	sliceCopy := make([]int, len(slice1))
	fmt.Println("SliceCopy (before copy):", sliceCopy)
	copy(sliceCopy, slice1)
	fmt.Println("SliceCopy (after copy):", sliceCopy)

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice1:", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slices are equal")
	}

	twoD := make([][]int, 4)
	for i := 0; i < 4; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
