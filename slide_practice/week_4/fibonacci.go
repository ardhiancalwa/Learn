package main

import "fmt"

func fibonnaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}

	return fibonnaci(i-1) + fibonnaci(i-2)
}

func main() {
	var i int
	for i = 0; i <= 10; i++ {
		fmt.Print(fibonnaci(i), " ")
	}
}
