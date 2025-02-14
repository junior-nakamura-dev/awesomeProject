package main

import "fmt"

func main() {

	defer fmt.Println("It will be the last line executed because of the defer instruction")
	fmt.Println("It will be the first line printed")
	fmt.Println("It will be the second line printed")

}
