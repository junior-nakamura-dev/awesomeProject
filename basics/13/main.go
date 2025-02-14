package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sumPointer(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	a := 10
	b := 20

	fmt.Println("(No pointer) sum:", sum(a, b))
	fmt.Println("(No pointer) a :", a)
	fmt.Println("(No pointer) b :", b)

	//Pointer make value mutable
	fmt.Println("(Pointer) sum:", sumPointer(&a, &b))
	fmt.Println("(Pointer) a:", a)
	fmt.Println("(Pointer) b:", b)

}
