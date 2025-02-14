package main

func main() {

	for i := 0; i < 5; i++ {
		println("I = ", i)
	}

	index := 5

	for index != 10 {
		println("Index = ", index)
		index++
	}

	number := []string{"One", "two", "three", "four", "five", "six", "seven"}

	for index, value := range number {
		println("Index = ", index, value)
	}

}
