package main

import "fmt"

func hitungDeret(n, m int) float64 {
	var jumlahDeret float64
	var i int

	jumlahDeret = 0.0

	if n%2 != 0 {
		for i = n; i <= m; i++ {
			if i%2 == 0 {
				jumlahDeret -= 3.0 / float64(i)
			} else {
				jumlahDeret += 3.0 / float64(i)
			}
		}
	}
	if n%2 == 0 {
		for i = n; i <= m; i++ {
			if i%2 == 0 {
				jumlahDeret += 3.0 / float64(i) 
			} else {
				jumlahDeret -= 3.0 / float64(i)
			}

		}

	}
	return jumlahDeret
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Printf("%.2f\n", hitungDeret(n, m))
}
