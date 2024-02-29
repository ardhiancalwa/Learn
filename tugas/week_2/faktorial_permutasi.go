package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(x, y int) int {
	if x >= y {
		return faktorial(x) / faktorial(x-y)
	}
	return faktorial(y) / faktorial(y-x)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(faktorial(x), faktorial(y), permutasi(x, y))
}
