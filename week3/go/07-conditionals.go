package main

import "fmt"

func main() {
	// Example Coditionals
	x := 100

	// iVjjjf x < 100 {
	// 	fmt.Println("Number is less than 100")
	// } else if x >= 100 && x <= 200 {
	// 	fmt.Println("Number is between 100 and 200")
	// } else {
	// 	fmt.Println("Number must be greater than 200")
	// }

	if x <= 100 {
		fmt.Println("Number is less than 100")
	}

	if x >= 100 && x <= 200 {
		fmt.Println("Number is between 100 and 200")
	}

	if x > 200 {
		fmt.Println("Number must be greater than 200")
	}

}
