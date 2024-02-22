package main

import "fmt"

func main() {
	var skor1, skor2, skor3, skor4, skor5, skor6 int
	var i, n, jumA, jumB int
	var A, B, imbang byte

	A = 'A'
	B = 'B'
	imbang = '0'

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&skor1, &skor2, &skor3, &skor4, &skor5, &skor6)
			jumA += skor1+skor2+skor3
			jumB += skor4+skor5+skor6
	}
	fmt.Print(jumA, jumB, " ")
	if jumA > jumB {
		fmt.Printf("%c",A)
	} else if jumA < jumB {
		fmt.Printf("%c",B)
	} else {
		fmt.Printf("%c", imbang)
	}
}