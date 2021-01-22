package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	total := add(5, 8)
	fmt.Println(total)
}
