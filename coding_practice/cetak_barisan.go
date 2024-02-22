package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	for n <= m {
		fmt.Println(n)
		n++
	}
}