package main

import "fmt"

func menu() {
	fmt.Println("-------------------------")
	fmt.Println("         M E N U         ")
	fmt.Println("-------------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-------------------------")
}

func hitungJumlah() {
	var jumlah, a, b int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	jumlah = a + b
	fmt.Println("Hasil penjumlahan:", jumlah)
}

func hitungKali() {
	var kali, a, b int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	kali = a * b
	fmt.Println("Hasil perkalian:", kali)
}

func hitungBagi() {
	var a, b int
	var bagi float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	bagi = float64(a) / float64(b)
	fmt.Println("Hasil pembagian:", bagi)
}

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		}
		if pilih == 4 {
			break
		}
	}
}
