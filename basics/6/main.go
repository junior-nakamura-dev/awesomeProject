package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	//get the first 2 values
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])

	//ignore the first 2 values
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)
	//when append goLang duplicate the slice size
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
