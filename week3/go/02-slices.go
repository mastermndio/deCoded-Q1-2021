package main

import "fmt"

func main() {
	//Our First Slice
	x := []int{5, 10, 1, 532425}

	//Reference item ion slice
	fmt.Println(x[3])

	//Change item in slice
	x[2] = 56

	//Add item to slice
	x = append(x, 76)

	fmt.Println(x)
}
