package main

import "fmt"

func main() {
	var A, B, C int
	var jumA, jumB, jumC int
	var total int
	var i, jumMobil int
	var tipeMobil string

	fmt.Scan(&jumMobil)
	for i = 1; i <= jumMobil; i++ {
		fmt.Scan(&tipeMobil)
		if tipeMobil == "A" {
			jumA += 10000
			A++
		} else if tipeMobil == "B" {
			jumB += 23000
            B++
		} else if tipeMobil == "C" {
			jumC += 45000
            C++
		}
	}
	total = jumA + jumB + jumC
	fmt.Println(total, A, B, C)
}