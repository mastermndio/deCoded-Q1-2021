package main

import "fmt"

func main() {
	// For numbers 0-100, print each number
	// for each number that is divisible by 3, print "Fizz"
	// for each number that is divisible by 5, print "Buzz"
	// for each number that is divisible by 3 and 5, print "Fizz Buzz"

	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}
}
