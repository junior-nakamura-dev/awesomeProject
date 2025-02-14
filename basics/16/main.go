package main

import "fmt"

func main() {

	var myVar interface{} = "Test"
	fmt.Println(myVar.(string))

	var res, ok = myVar.(int)

	fmt.Printf("The res is %v and the ok is %v ", res, ok)

}
