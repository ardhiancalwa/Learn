package main

import "fmt"

func main() {
	var n, widest, leaf, i  int

	fmt.Scan(&n)
	widest = leaf
	for i = 1; i <= n; i++ {
		fmt.Scan(&leaf)
		if leaf > widest {
			widest = leaf
		} else {
			widest = widest
		}
	}
	fmt.Println(widest)
}