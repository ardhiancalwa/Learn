package main

import "fmt"

func main() {
	var A, B, C byte
	var jumA, jumB, jumC int
	var tebakA, tebakB, tebakC int
	var i, n int

	A = 'A'
	B = 'B'
	C = 'C'

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&tebakA, &tebakB, &tebakC)
		if tebakA%2 != 0 {
			jumA += tebakA
		}
		if tebakB%2 != 0 {
			jumB += tebakB
		}
		if tebakC%2 != 0 {
			jumC += tebakC
		}
	}
	if jumA > jumB && jumA > jumC {
		fmt.Printf("%c %d", A, jumA)
	} else if jumB > jumA && jumB > jumC {
		fmt.Printf("%c %d", B, jumB)
	} else if jumC > jumA && jumC > jumB {
		fmt.Printf("%c %d", C, jumC)
	} else {
		fmt.Println("Imbang")
	}
}
