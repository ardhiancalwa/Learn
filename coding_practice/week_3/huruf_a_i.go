package main

import "fmt"

func sum(kar byte, jumlah *int) {
	/*  I.S. Terdefinisi kar byte dan jumlah int
	F.S. jumlah huruf selain  'a' dan 'i' terdefinisi di jumlah int */
	for kar != '.' {
		if kar != 'a' && kar != 'i' {
			*jumlah++
		}
		fmt.Scanf("%c", &kar)
	}
}

func main() {
	var jumlah int
	var kar byte
	fmt.Scanf("%c", &kar)
	sum(kar, &jumlah)
	fmt.Println(jumlah)
}
