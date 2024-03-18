package main

import "fmt"

func main() {
	var x, y, i int
	fmt.Scan(&x, &y)
	fmt.Println(pangkat(x,y,i))
}

func pangkat(x, y, i int) int {
	if i < y {
		return x * pangkat(x, y, i+1)
	}
	return 1
}
