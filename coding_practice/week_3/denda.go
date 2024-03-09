package main

import "fmt"

func denda(telat int) {
	/*
		I.S terdefinisi bilangan bulat hariTelambat yang menyatakan hari keterlambatan
		F.S menampilkan jumlah uang denda atau string "cabut keanggotaan" jika keterlambatan melebihi 10 hari
	*/
	if telat <= 5 {
		telat *= 1000
		fmt.Println(telat)
	} else if telat >= 6 && telat <= 10 {
		telat *= 2000
		fmt.Println(telat)
	} else if telat > 10 {
		fmt.Println("cabut keanggotaan")
	}

	

}

func main() {
	var hariTerlambat int
	fmt.Scan(&hariTerlambat)
	denda(hariTerlambat)
}
