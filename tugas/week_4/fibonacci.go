package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Print(fibonacci(i), " ")
	}
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
