package math

// When First letter is capitalize it means it is a public method/function/variable and it could be used in another place
func Sum[T int | float64](a T, b T) T {
	return a + b
}

func sum[T int | float64](a T, b T) T {
	return a + b
}
