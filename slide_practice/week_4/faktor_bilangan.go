package main

import "fmt"

func main() {
	var N, i int
	fmt.Scan(&N)
	i = 1
	faktorBilangan(N, i)
}

func faktorBilangan(N, i int) {
	if i == N {
		fmt.Print(N)
	} else {
		if N%i == 0 {
			fmt.Print(i, " ")
		}
		faktorBilangan(N, i+1)
	}
}
