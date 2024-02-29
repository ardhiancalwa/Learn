package main

import "fmt"

func jumlahKelipatan(m, n, x int) int {
	var hasilJumlah, i int

	if m <= n {
		for i = m; i <= n; i++ {
			if i%x == 0 {
				hasilJumlah += i
			}
		}
	} else if m >= n {
		for i = n; i <= m; i++ {
			if i%x == 0 {
				hasilJumlah += i
			}
		}
	}

	return hasilJumlah
}

func main() {
	var bil1, bil2, x int
	fmt.Scan(&bil1, &bil2, &x)
	fmt.Println(jumlahKelipatan(bil1, bil2, x))
}
