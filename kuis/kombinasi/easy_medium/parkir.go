package main

import "fmt"

func main() {
	var kendaraan byte
	var bayar, jumBayarM, jumBayarB int

	fmt.Scanf("%c %d\n", &kendaraan, &bayar)

	jumBayarB = 0
	jumBayarM = 0

	for kendaraan != 'x' && bayar != 0 {
		if kendaraan == 'm' {
			jumBayarM += bayar
		} 
		if kendaraan == 'b' {
			jumBayarB += bayar
		}
		fmt.Scanf("%c %d\n",&kendaraan, &bayar)
	}
	fmt.Println(jumBayarB, jumBayarM)
}