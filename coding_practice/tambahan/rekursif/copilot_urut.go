package main

import "fmt"

// Fungsi rekursif untuk mencetak bilangan ganjil hingga batas tertentu
func printOddNumbers(n int) {
	if n > 0 {
		printOddNumbers(n - 1) // Panggilan rekursif
		if n%2 != 0 {
			fmt.Println(n) // Cetak bilangan ganjil
		}
	}
}

func main() {
	var limit int
	fmt.Scan(&limit)
	printOddNumbers(limit) // Memanggil fungsi rekursif
}
