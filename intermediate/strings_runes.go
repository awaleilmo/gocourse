package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawmessage variable is", len(rawMessage))
	fmt.Println("The first character in message var is", message[0]) // ASCII

	greeting := "Hello "
	name := "Awal"
	fmt.Println(greeting + name)

	str1 := "Apple"          // A has an ASCII value of 65
	str2 := "banana"         // b has an ASCII value of 98
	str3 := "app"            // a has an ASCII value of 97
	fmt.Println(str1 < str2) // false because 65 < 98
	fmt.Println(str3 < str1) // true because 97 < 65

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("Hexadecimal value: %x\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := 'N'
	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("Type of cstr is %T\n", cstr)
}
