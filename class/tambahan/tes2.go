package main

import "fmt"

func main() {
	var n  ame, kota_lahir string
	var nim, umur int

	fmt.Print("Masukkan nama keren kamu: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan nim aneh kamu: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan umur kamu: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan kota lahir kamu: ")
	fmt.Scan(&kota_lahir)

	fmt.Println("Umur saya :", umur)
	fmt.Println("NIM saya :", nim)
	fmt.Println("Nama saya :", name)
	fmt.Println("Kota lahir saya :", kota_lahir)

	// var hasil, bil1, bil2 int

	// fmt.Print("Masukkan Nilai Bilangan 1: ")
	// fmt.Scan(&bil1)
	// fmt.Print("Masukkan Nilai Bilangan 2: ")
	// fmt.Scan(&bil2)

	// // rumus
	// hasil = bil1 + bil2

	// fmt.Println(hasil)
}
