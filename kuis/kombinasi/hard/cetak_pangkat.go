package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	var i, suku int

	fmt.Scan(&n, &m)
	suku = 1
	for i = n; i <= m; i++ {
		if suku%2 != 0 {
			fmt.Printf("%.2f ", math.Pow(2, float64(i)))
		} else if suku%2 == 0 {
			fmt.Printf("%.2f ", math.Pow(2, float64(-i)))
		}
		suku++
	}
}
