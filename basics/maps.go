package main

import (
	"fmt"
	"maps"
)

func main() {

	// var map1 = map[string]int{
	// 	"one": 1,
	// 	"two": 2,
	// 	"three": 3,
	// }

	// fmt.Println(map1)

	// map2 := map[string]int{
	// 	"four": 4,
	// 	"five": 5,
	// 	"six": 6,
	// }

	// fmt.Println(map2)

	// map3 := make(map[string]int)
	// map3["seven"] = 7
	// map3["eight"] = 8
	// map3["nine"] = 9

	// fmt.Println(map3)

	// fmt.Println("Value of key 'one':", map1["one"])
	// fmt.Println("Value of key 'four':", map2["four"])
	// fmt.Println("Value of key 'seven':", map3["seven"])

	// fmt.Println("Length of map1:", len(map1))
	// fmt.Println("Length of map2:", len(map2))
	// fmt.Println("Length of map3:", len(map3))

	// fmt.Println("Keys of map1:", map1)
	// fmt.Println("Keys of map2:", map2)
	// fmt.Println("Keys of map3:", map3)

	// fmt.Println("Values of map1:", map1)
	// fmt.Println("Values of map2:", map2)
	// fmt.Println("Values of map3:", map3)

	// fmt.Println("Key-Value pairs of map1:", map1)
	// fmt.Println("Key-Value pairs of map2:", map2)
	// fmt.Println("Key-Value pairs of map3:", map3)

	// fmt.Println("Key 'one' exists in map1:", "one" in map1)
	// fmt.Println("Key 'four' exists in map2:", "four" in map2)
	// fmt.Println("Key 'seven' exists in map3:", "seven" in map3)

	// fmt.Println("Key 'ten' exists in map1:", "ten" in map1)
	// fmt.Println("Key 'eleven' exists in map2:", "eleven" in map2)
	// fmt.Println("Key 'twelve' exists in map3:", "twelve" in map3)

	// fmt.Println("Key 'one' exists in map1:", map1["one"])
	// fmt.Println("Key 'four' exists in map2:", map2["four"])
	// fmt.Println("Key 'seven' exists in map3:", map3["seven"])

	// fmt.Println("Key 'ten' exists in map1:", map1["ten"])
	// fmt.Println("Key 'eleven' exists in map2:", map2["eleven"])
	// fmt.Println("Key 'twelve' exists in map3:", map3["twelve"])

	myMap := make(map[string]int)

	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 10

	fmt.Println("Default Value:", myMap)
	fmt.Println("Value when key is present:", myMap["key1"])
	fmt.Println("value when key is not present:", myMap["key2"])
	myMap["code"] = 11
	fmt.Println("Value when key is updated:", myMap)

	delete(myMap, "key1")
	fmt.Println("Value when key is deleted:", myMap)

	myMap["key1"] = 12
	myMap["key2"] = 13
	myMap["key3"] = 14
	fmt.Println("Value when key is added:", myMap)

	// clear(myMap)
	// fmt.Println("Value when map is cleared:", myMap)

	value, unknownvalue := myMap["key1"]
	fmt.Println("Value when key is present:", value)
	fmt.Println("Value when key is not present:", unknownvalue)

	mymap2 := map[string]int{"a": 1, "b": 2}
	mymap3 := map[string]int{"a": 1, "b": 2}

	if maps.Equal(mymap2, mymap3) {
		fmt.Println("Maps are equal")
	}

	for _, v := range mymap2 {
		fmt.Println(v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value")
	} else {
		fmt.Println("The map is not initialized to nil value")
	}

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println("The map is initialized to non-nil value", myMap4)

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println("The map is initialized to non-nil value", myMap5)

}
