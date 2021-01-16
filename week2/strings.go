package main

import (
	"fmt"
	"strings"
)

func main() {
	//This is an example of a string
	fmt.Println("Hello")

	var ourString string
	ourString = "This is a Test"

	fmt.Println(ourString)

	//This is an example of a Rune
	var ourRune rune
	ourRune = 'g'
	fmt.Printf("Type: %T, \v Value: %v\n", ourRune, ourRune)

	//Example of needing special characters
	fmt.Println("Don't Touch That, tha's my food")
	fmt.Println("She said, \"thats my food too\"")

	Example of string literal
	fmt.Println(`Hello my name is Aaron\n,
				and I am the coolest person`)

	//example of length
	fmt.Println(len("Hello World"))

	//String Formatting
	name := 'A'
	age := 30
	fmt.Printf("Hello my name is %c and I am %d years old", name, age)

	//String Method
	//Uppercase all
	fmt.Println(strings.ToUpper("Gopher dbhdbhbdwhbdwhj"))

	//Example Split
	s := "Why Are We Here"
	sSplit := strings.Split(s, "W")
	fmt.Println(sSplit)
}
