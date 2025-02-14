package main

import "fmt"

func main() {

	salaryMap := map[string]int{}
	salaryMap["Junior"] = 10000
	salaryMap["Nakamura"] = 12000

	for name, salary := range salaryMap {
		fmt.Printf("The %s salary is %d USD \n", name, salary)
	}

	anotherWayDeclareMap := map[string]int{"Jun": 50000}
	fmt.Println(anotherWayDeclareMap)
	delete(anotherWayDeclareMap, "Jun")
	fmt.Println(anotherWayDeclareMap)

	anotherWayDeclareMapWithMake := make(map[string]int)
	anotherWayDeclareMapWithMake["Junior"] = 10000
	fmt.Println(anotherWayDeclareMapWithMake)

}
