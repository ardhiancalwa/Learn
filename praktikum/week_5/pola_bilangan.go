package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(polaSuku(N))
}

// Fungsi rekursif untuk menghitung nilai suku ke-N
func polaSuku(n int) int {
	if n == 1 || n == 2 || n == 3 {
		return 1
	}
	return polaSuku(n-1) + polaSuku(n-2) + polaSuku(n-3)
}
