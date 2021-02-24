package main

import "fmt"

func main() {
	// Here is an example of a variable
	name := "Maria"
	fmt.Println(name)

	// Data Types
	// Example of a String
	var ourString string
	ourString = "njsdfanfkj 42314"
	fmt.Printf("Type: %T\n", ourString)

	// Example of int
	var ourInt int
	ourInt = 55
	fmt.Printf("Type: %T\n", ourInt)

	// Example of a boolean
	var ourBool bool
	ourBool = true
	fmt.Printf("Type: %T\n", ourBool)

	// Example of a Float
	var ourFloat float32
	ourFloat = 5.5436
	fmt.Printf("Type: %T\n", ourFloat)

	// Collection
	//Example of a Slice
	var ourSlice []int
	ourSlice = []int{5, 4, 2, 875, 424}
	fmt.Println(ourSlice)

	// rune byte
	var ourrune rune
	ourrune = 'h'
	fmt.Printf("Type: %v\n", ourrune)
}
