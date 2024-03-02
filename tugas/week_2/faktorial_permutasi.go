package main

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(x, y int) int {
	var xFactorial, yFactorial int
	var hasilPermutasi int
	if x >= y {
		xFactorial = factorial(x)
		yFactorial = factorial(x - y)

	} else if x < y {
		xFactorial = factorial(y)
		yFactorial = factorial(y - x)
	}
	hasilPermutasi = xFactorial / yFactorial

	return hasilPermutasi
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	permutation := permutation(x, y)

	fmt.Println(factorial(x), factorial(y), permutation)
}