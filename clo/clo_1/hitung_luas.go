package main

import "fmt"

func main() {
	var panjang, lebar, luasPersegiPanjang int
	fmt.Scan(&panjang, &lebar)
	hitungLuas(panjang, lebar, &luasPersegiPanjang)
	fmt.Println(luasPersegiPanjang)
}

func hitungLuas(panjang, lebar int, luasPersegiPanjang *int) {
    *luasPersegiPanjang = panjang * lebar
}
