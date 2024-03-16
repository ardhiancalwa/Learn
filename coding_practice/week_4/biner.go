package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	biner(N)
}

func biner(N int) {
	if N > 0 {
		biner(N / 2) 
		fmt.Print(N % 2)
	}
}
