package main

import (
	"awesomeProject/18/math"
	"fmt"
)

func main() {
	result := math.Sum(10, 30)
	fmt.Println("Result: ", result)

	//result1 := math.sum(10, 30) It not works because "sum" is not public like "Sum" this rule applies for struct, vairables, method, function and so on...
}
