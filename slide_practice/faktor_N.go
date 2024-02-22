package main

import "fmt"

func main() {
	var N, d int
	var s bool

	fmt.Scan(&N)
	for d = 1; d <= N; d++ {
		if N%d == 0 {
			s = true
		} else {
			s = false
		}
		fmt.Println(d, s)
	}
	
}