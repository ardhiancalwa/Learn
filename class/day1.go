package main

import (
	"fmt"
	// "io/ioutil"
)

func main() {
	// 1
	// var topi, sabuk, pulpen, tali, pisau, lengkap bool
	// var i, N int

	// fmt.Scan(&N)

	// lengkap = true
	// for i = 1; i <= N && lengkap; {
	// 	fmt.Scan(&topi, &pulpen, &tali, &pisau, &sabuk)
	// 	lengkap = topi && pisau && pulpen && tali && sabuk
	// 	i++
	// }

	// // data, err := ioutil.ReadFile("input.txt")
	// // if err != nil {
	// //     fmt.Println("Gagal membaca file:", err)
	// //     return
	// // }

	// // fmt.Println("Isi file teks:")
	// // fmt.Println(string(data))

	// fmt.Println(lengkap)

	// 2
	// var bil1, bil2, hasil int
	// fmt.Print("Masukkan bil1 : ")
	// fmt.Scan(&bil1)
	// fmt.Print("Masukkan bil2 : ")
	// fmt.Scan(&bil2)
	// hasil = bil1 + bil2

	// fmt.Println("Jadi hasil penjumlahannya adalah", hasil)

	// 3
	// var b1, b2, hasil int8
	// b1 = 120
	// b2 = 8
	// hasil = b1 + b2

	// fmt.Println(hasil) // output -128

	// 4
	// var a, b, c int
	// a = 71
	// b = 21

	// c = a ^ b
	// fmt.Println(c)

	// 5
	var total, u10, u5, u1 int
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&total)

	u10 = total / 10000
	u5 = total % 10000 / 5000
	u1 = total % 10000 % 5000 / 1000

	fmt.Println(u10, u5, u1)
}
