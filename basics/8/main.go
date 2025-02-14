package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := sum(50, 1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func sum(a, b int) (int, error) {

	if a+b > 50 {
		return 0, errors.New("Sum is greater than 50")
	}
	return a + b, nil
}
