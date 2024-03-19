package main

import "fmt"

func main() {
	var fib, N int
	fmt.Scan(&N)	
	fib = fibonacci(N)
	fmt.Println(fib)
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
