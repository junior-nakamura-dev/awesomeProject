package main

import "fmt"

func main() {

	a := 10
	fmt.Println("a: ", a)
	fmt.Println("&a: ", &a)

	var pointer *int = &a
	fmt.Println("pointer: ", pointer)

	b := &a
	fmt.Println("b: ", b)
	*b = 30
	fmt.Println("*b: ", *b)
	fmt.Println("a: ", a)
}
