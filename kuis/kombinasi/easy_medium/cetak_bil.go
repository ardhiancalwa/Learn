package main

import "fmt"

func main() {
	var i, n, m int
	var bilangan string

	fmt.Scan(&bilangan)
	fmt.Scan(&n, &m)
	for i = n; i <= m; i++ {
		if bilangan == "genap" {
			if i%2 == 0 {
                fmt.Print(i, " ")
            }
		} else if bilangan == "ganjil" {
			if i%2!= 0 {
                fmt.Print(i, " ")
            }
		}
	}
}
