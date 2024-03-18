package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetakRerataRekursif(N, 0, 1) // Memulai rekursi dengan i=1 dan jumlah=0
}

// Fungsi rekursif tail-end untuk mencetak rata-rata
func cetakRerataRekursif(N, jumlah, i int) {
	var rata float64
	if i > N { // Kondisi dasar rekursi
		return
	}
	jumlah += i
	rata = float64(jumlah) / float64(i)
	fmt.Println(jumlah, rata)
	cetakRerataRekursif(N, jumlah, i+1) // Pemanggilan rekursif dengan nilai i dan jumlah yang diperbarui
}
