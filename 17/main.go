package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func compare[T comparable](n1 T, n2 T) bool {
	if n1 == n2 {
		return true
	}
	return false
}

func main() {
	m1 := map[string]int{"a": 100, "b": 200}
	m2 := map[string]float64{"a": 50.5, "b": 50.53}
	m3 := map[string]MyNumber{"a": 100, "b": 100}
	fmt.Println("M1 sum: ", sum(m1))
	fmt.Println("M2 sum: ", sum(m2))
	fmt.Println("M3 sum: ", sum(m3))
	fmt.Println(compare(10, 10.0))
}
