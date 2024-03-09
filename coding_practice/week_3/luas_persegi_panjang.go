package main

import "fmt"

func hitungLuas(p, l int, luas *int) {
	/* I.S.: Terdefinisi panjang (p) dan lebar (l) persegi panjang
	   F.S.: luas berisi hasil perhitungan luas persegi panjang */
	*luas = p * l
}

func main() {
	var panjang, lebar, luasPersegiPanjang int
	fmt.Scan(&panjang, &lebar)
	hitungLuas(panjang, lebar, &luasPersegiPanjang)
	fmt.Println(luasPersegiPanjang)
}
