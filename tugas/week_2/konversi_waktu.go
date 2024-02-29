package main

import "fmt"

func konversiWaktu(jam, menit, detik int) int {
	var konversi int
	konversi = jam*3600 + menit*60 + detik
	return konversi
}

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	fmt.Println("Hasil konversi =", konversiWaktu(jam, menit, detik), "detik")
}
