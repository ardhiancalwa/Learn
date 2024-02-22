package main

import "fmt"

func main() {
	var harga int
	var suwirAyam, cakue, atiAmpela, tambahPorsi bool

	fmt.Scan(&suwirAyam, &cakue, &atiAmpela, &tambahPorsi)
	harga = 6000

	if suwirAyam && cakue && atiAmpela && tambahPorsi {
		harga += 3000 + 1500 + 4500 + 4000
	} else if cakue && atiAmpela && tambahPorsi {
		harga += 1500 + 4500 + 4000
	} else if suwirAyam && atiAmpela && tambahPorsi {
		harga += 3000 + 4500 + 4000
	} else if suwirAyam && cakue && tambahPorsi {
		harga += 3000 + 1500 + 4000
	} else if suwirAyam && cakue && atiAmpela {
		harga += 3000 + 1500 + 4500
	} else if suwirAyam && cakue {
		harga += 3000 + 1500
	} else if suwirAyam && atiAmpela {
		harga += 3000 + 4500
	} else if suwirAyam && tambahPorsi {
		harga += 3000 + 4000
	} else if cakue && atiAmpela {
		harga += 1500 + 4500
	} else if cakue && tambahPorsi {
		harga += 1500 + 4000
	} else if atiAmpela && tambahPorsi {
		harga += 4500 + 4000
	} else if suwirAyam {
		harga += 3000
	} else if cakue {
		harga += 1500
	} else if atiAmpela {
		harga += 4500
	} else if tambahPorsi {
		harga += 4000
	} else {
		harga += 0
	}

	fmt.Println(harga)
}
