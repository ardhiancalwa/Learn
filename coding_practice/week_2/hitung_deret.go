package main

import "fmt"

func hitungDeret(n, m int) float64 {
	var jumlahDeret, suku float64
	var i int

	jumlahDeret = 0.0
	suku = 1.0
	for i = n; i <= m; i++ {
		jumlahDeret += suku * (3.0 / float64(i))
		suku *= -1
	}
	return jumlahDeret
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Printf("%.2f\n", hitungDeret(n, m))
}
