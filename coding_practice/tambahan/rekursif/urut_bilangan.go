package main

import "fmt"

func main() {
	var i, n int
	i = 1
	fmt.Scan(&n)
	if n %2 == 0 {
		n -= 1
	}
	urutBilangan(n, i)
}

func urutBilangan(n, i int) {
	if i == n {
		fmt.Print(i)
	} else {
		fmt.Print(i, " ")
		urutBilangan(n, i+2)
	}
}