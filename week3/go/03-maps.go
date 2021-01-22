package main

import "fmt"

func main() {
	//Example of a map
	x := map[string]string{"Name": "Aaron", "Age": "30"}

	for k, i := range x {
		fmt.Printf("%v %v\n", k, i)
	}
}
