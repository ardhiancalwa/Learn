package main

import "fmt"

func fibonacci(n int) int {
	var jumlahFibonacci, bil1, bil2, bil2Lama, i int
	jumlahFibonacci = 0
	bil1 = 0
	bil2 = 1

	for i = 1; i <= n; i++ {
		jumlahFibonacci += bil1
		bil2Lama = bil2
		bil2 = bil1 + bil2
		bil1 = bil2Lama
	}
	return jumlahFibonacci
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}
