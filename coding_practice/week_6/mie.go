package main

import "fmt"

type warungMie struct {
	tipeMie string
	topping string
	pedas   bool
}

func main() {
	var mie warungMie
	fmt.Scan(&mie.tipeMie, &mie.pedas)
	hitungHarga(mie)
}
func hitungHarga(mie warungMie) {
	var harga int
	var topping int
	if mie.tipeMie == "kwetiau" {
		harga += 8000
	} else if mie.tipeMie == "bihun" {
		harga += 7000
	} else if mie.tipeMie == "kuning" {
		harga += 9000
	}
	for mie.topping != "0" && topping != 5 {
		fmt.Scan(&mie.topping)
		if mie.topping == "ayam" {
			harga += 5000
			topping++
		}
		if mie.topping == "telur" {
			harga += 3000
			topping++
		}
		if mie.topping == "sayur" {
			harga += 3000
			topping++
		}
		if mie.topping == "pangsit" {
			harga += 5000
			topping++
		}
		
	}
	fmt.Println(harga)
}
