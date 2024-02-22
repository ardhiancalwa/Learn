package main

import "fmt"

func main() {
	var anggota1, anggota2, anggota3 byte
	var poin, i int

	for i = 1; i <= 10; i++ {
		fmt.Scanf("%c%c%c", &anggota1, &anggota2, &anggota3)
		if anggota1 == 'O' || anggota2 == 'O' || anggota3 == 'O' {
			poin++
		}

		if anggota1 == 'T' && anggota2 == 'T' && anggota3 == 'T' {
			break 
		}
	}
	fmt.Println("Poin =",poin)
}
