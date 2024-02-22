package main

import "fmt"

func main() {
	var x, y, a, b int
	var i int

	fmt.Scan(&x, &y, &a, &b)
	for i = x; i <= y; i++ {
		if i/100 != a && i%10 != b {
			fmt.Println(i)
		}
	}
}
