package main

import "fmt"

func main() {
	var panjang, lebar, luasPersegiPanjang int
	fmt.Scan(&panjang, &lebar)
	hitungLuas(panjang, lebar, &luasPersegiPanjang)
	fmt.Println(luasPersegiPanjang)
}

func hitungLuas(p, l int, luas *int) {
	*luas = p * l
}
