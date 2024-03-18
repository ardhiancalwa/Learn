package main

import "fmt"

func main() {
	var n, i int
	fmt.Scan(&n)
	i = 1
	if n%2 == 0 {
        n -= 1
    }
	ganjil(n, i)
}

func ganjil(n, i int) {
	if i == n {
		fmt.Print(i)
	} else {
		fmt.Print(i, " ")
		ganjil(n, i+2)
	}
}
