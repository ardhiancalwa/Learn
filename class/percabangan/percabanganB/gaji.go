package main

import "fmt"

func main() {
	var jabatan string
	var masaKerja, jumAnak, gaji_pokok int

	gaji_pokok = 0
	jumAnak = 0
	fmt.Scan(&jabatan, &masaKerja, &jumAnak)
	if jabatan == "Staf" {
		if masaKerja < 5 {
			gaji_pokok += 4000
			jumAnak *= 0
		} else if masaKerja > 10 {
			gaji_pokok += 5000
			jumAnak *= 100
		} else if masaKerja >= 5 && masaKerja <= 10 {
			gaji_pokok += 4000
			jumAnak *= 100
		}
	}

	if jabatan == "Manager" {
		if masaKerja > 10 {
			gaji_pokok += 10000
			jumAnak *= 300
		} else {
			gaji_pokok += 8500
			jumAnak *= 300
		}
	}

	if jabatan == "Direktur" {
		gaji_pokok += 20000
		jumAnak *= 500
	}

	fmt.Println(gaji_pokok, "+", jumAnak, "=", gaji_pokok+jumAnak)
}
