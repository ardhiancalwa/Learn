package main

import "fmt"

func main() {
	var i, n, m int
	var jumKali int

	fmt.Scan(&n, &m)
	for i = 1; i <= n; i++ {
		jumKali += m
	}
	fmt.Println(jumKali)
}