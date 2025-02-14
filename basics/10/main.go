package main

import "fmt"

func main() {
	double := func() int {
		return sum(5, 4, 1) * 2
	}

	fmt.Println(double())

	triple := func() int {
		return sum(5, 4, 1) * 3
	}()
	fmt.Println(triple)

}

func sum(numbers ...int) int {

	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}
