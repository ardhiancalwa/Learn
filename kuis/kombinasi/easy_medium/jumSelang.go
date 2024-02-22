package main

import "fmt"

func main() {
	var n, m int
	var total, nilaiAwal float64

	fmt.Scan(&n, &m)
	nilaiAwal = 1.0
	for i := n; i <= m; i++ {
		total += float64(nilaiAwal) * 3.0 / float64(i)
		nilaiAwal *= -1
	}

	fmt.Printf("%.2f\n", total)
}
