package main

import (
	"fmt"
)

func main() {
	var nama, tempat_lahir, prodi, tgl_lahir, fakultas string
	var nim, angkatan, kelas, umur int

	nama = "Ardhian Calwa Nugraha"
	nim = 103012300404
	tempat_lahir = "Pangkalan Bun"
	tgl_lahir = "16-05-2005"
	umur = 18
	prodi = "S1 Informatika"
	fakultas = "Informatika"
	angkatan = 47
	kelas = 07

	fmt.Println("Nama: ", nama)
	fmt.Println("Nim: ", nim)
	fmt.Println("Tanggal Lahir: ", tgl_lahir)
	fmt.Println("Tempat Lahir: ", tempat_lahir)
	fmt.Println("Umur: ", umur)
	fmt.Println("Program Studi: ", prodi)
	fmt.Println("Fakultas: ", fakultas)
	fmt.Println("Angkatan: ", angkatan)
	fmt.Println("Kelas: ", kelas)

	fmt.Println("Telkom Univerisity")
	fmt.Println("Sinergi Bangun Negeri")

	var luas, panjang, lebar int
	panjang = 5
	lebar = 3
	luas = panjang * lebar

	fmt.Println(luas)

	fmt.Println("Nama : Ardhian Calwa Nugraha")
	fmt.Println("Kota Asal : Pangkalan Bun")
	fmt.Println("Kota Domisili : Trenggalek")	

	var plus, times, n1, n2, n3, n4 int
	n1 = 1
	n2 = 2
	n3 = 2
	n4 = 3
	plus = n1 + n2
	times = n3 * n4

	fmt.Println(plus)
	fmt.Println(times)
}