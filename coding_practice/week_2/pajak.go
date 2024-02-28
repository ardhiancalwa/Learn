package main

import "fmt"

func hargaBarangKenaPajak(asal string, harga int) float64 {
	// Mengembalikan harga barang setelah terkena pajak
	// Barang dalam negeri terkena pajak 8%, luar negeri 18%
	var hargaBarang float64
	if asal == "dalam" {
		hargaBarang = float64(harga) - (float64(harga) * 0.08)
	} else if asal == "luar" {
		hargaBarang = float64(harga) - (float64(harga) * 0.18)
	}
	return hargaBarang
}

func main() {
	var hbarang int
	var asal string
	var hbarang2 float64
	fmt.Scan(&asal, &hbarang)
	hbarang2 = hargaBarangKenaPajak(asal, hbarang)
	fmt.Printf("%.2f", hbarang2)
}
