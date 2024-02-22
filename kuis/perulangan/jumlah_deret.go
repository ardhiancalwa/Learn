package main

import "fmt"

func main() {
	var i, n, m int
	var deret float64

	deret = 0.0
	fmt.Scan(&n, &m)
	for i = n; i <= m; i++ {
		deret += 2.0/float64(i)
	}
	fmt.Printf("%.2f", deret)

}
