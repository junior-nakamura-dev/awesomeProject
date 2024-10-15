package main

import (
	"awesomeProject/model"
	"fmt"
	"time"
)

func main() {

	address := model.Address{
		Street: "Street X",
		Number: 100,
		City:   "Maringa",
	}
	client := model.NewClient("Junior", address, time.Date(1994, 9, 2, 0, 0, 0, 0, time.UTC))

	fmt.Println(address)

	fmt.Println("Name: ", client.Name)
	fmt.Println("Address: ", client.Address)
	fmt.Println("Birthdate: ", client.Birthdate)
	fmt.Println("Age: ", client.GetAge())
	client.Address.Street = "New Stret"
	fmt.Println("Address: ", client.Address)

}

func mathOperations(number1 int, number2 int) (sumResult int, subtractResult int, divideResult int, multiplyResult int) {
	sumResult = number1 + number2
	subtractResult = number2 - number1
	divideResult = number1 / number2
	multiplyResult = number1 * number2
	return
}

//func main() {
//
//	values := make([]int, 0)
//	values = append(values, 5)
//	values = append(values, 2)
//	values = append(values, 6)
//	values = append(values, 7)
//	values = append(values, 8)
//	fmt.Println(values)
//
//	for _, value := range values {
//		fmt.Println(value)
//	}
//
//	for i := 0; i < len(values); i++ {
//		fmt.Println(values[i])
//	}
//
//	lastValue := values[len(values)-1]
//
//	fmt.Println("lastValue", lastValue)
//	fmt.Println("first3", values[:3])
//	fmt.Println("from3ToFinal", values[3:])
//
//	//MAP
//
//	cityMaps := make(map[string]int)
//	cityMaps["PR"] = 500
//	cityMaps["SP"] = 1000
//	fmt.Println(cityMaps)
//
//	fmt.Println(cityMaps["SP"])
//
//	value, exists := cityMaps["RJ"]
//
//	if exists {
//		fmt.Println("Value: ", value)
//	} else {
//		fmt.Println("Error: cityMaps[\"RJ\"]")
//	}
//
//	if value, exists := cityMaps["PR"]; exists {
//		fmt.Println("Key: ", "PR")
//		fmt.Println("Value: ", value)
//	}
//
//	delete(cityMaps, "PR")
//
//	for key, value := range cityMaps {
//		fmt.Println("Key: ", key, "Value: ", value)
//	}
//
//}

//
//func main() {
//	values := []int{2, 7, 11, 15}
//	target := 9
//
//	result := twoSumComplement(values, target)
//
//	fmt.Println(result)
//}
//
//func twoSumComplement(nums []int, target int) []int {
//	hashMap := make(map[int]int)
//	for i, num := range nums {
//		numberToFind := target - num
//		if hashMapIndex, ok := hashMap[numberToFind]; ok {
//			fmt.Println("ok", ok)
//			return []int{hashMapIndex, i}
//		}
//		hashMap[num] = i
//	}
//	return []int{}
//}

//func main() {
//	weight := 80
//	var test string
//	var testInt int
//	var testFloat float32
//	const man bool = true
//	const name string = "Test"
//	fmt.Println("Weight: ", weight)
//	fmt.Println("Man: ", man)
//	fmt.Println("Name: ", name)
//	weight = weight * 2
//	fmt.Println("Weight * 2: ", weight)
//	fmt.Println("test: ", test)
//	fmt.Println("testInt: ", testInt)
//	fmt.Println("testFloat: ", reflect.TypeOf(testFloat))
//	fmt.Println("Name: ", name+" "+strconv.FormatInt(int64(weight), 10))
//
//}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
