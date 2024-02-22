package main

import "fmt"

func main() {
	var tebakA, tebakB, tebakC int
	var i, n int
	var jumA, jumB, jumC int
	var A, B, C byte

	A = 'A'
	B = 'B'
	C = 'C'

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&tebakA, &tebakB, &tebakC)
        jumA += tebakA
        jumB += tebakB
        jumC += tebakC
	}

	if jumA < jumB && jumA < jumC {
	    fmt.Printf("%c %d", A, jumA)
	} 
	if jumB < jumA && jumB < jumC {
		fmt.Printf("%c %d", B, jumB)
	} 
	if jumC < jumA && jumC < jumB {
		fmt.Printf("%c %d", C, jumC)
	}
}