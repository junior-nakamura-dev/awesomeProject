package main

import "fmt"

func main() {
	fmt.Println(sum(5, 4, 1))
}

func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}
