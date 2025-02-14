package main

import "fmt"

// array
func main() {

	myArray := [3]int{10, 20, 30}

	for index, value := range myArray {
		fmt.Printf("Index: %d, Value: %v\n", index, value)
	}

}
