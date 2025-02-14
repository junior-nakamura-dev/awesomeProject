package main

import "fmt"

func main() {

	var a interface{} = 10
	var b interface{} = "Test"

	showTypeAndValue(a)
	showTypeAndValue(b)

	a = "I changed from int to string"
	showTypeAndValue(a)
}

func showTypeAndValue(t interface{}) {
	fmt.Printf("The type is %T and the value is %v \n", t, t)
}
