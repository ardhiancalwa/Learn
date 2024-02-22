package main

import "fmt"

func main() {
	var bil, i, n, jumlah int

	jumlah = 0 //inisialisasi
	fmt.Scan(&n)
	i = 1        //inisialisasi
	for i <= n { //kondisi
		fmt.Scan(&bil)
		jumlah += bil
		i++
	}
	fmt.Println(jumlah) // terminasi
}
