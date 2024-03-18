package main

import "fmt"

func main() {
	var n, i int
	i = 1
	fmt.Scan(&n)
	faktor(n, i)
}

func faktor(n, i int) {
	if i == n {
		fmt.Print(i)
	} else {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
		faktor(n, i+1)
	}
}
