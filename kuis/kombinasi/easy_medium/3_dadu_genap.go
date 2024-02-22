package main

import "fmt"

func main() {
	var tebak1, tebak2, tebak3 int
	var genap int

	fmt.Scan(&tebak1, &tebak2, &tebak3)
	for tebak1+tebak2+tebak3 != 12 {
		if tebak1%2 == 0 && tebak2%2 == 0 && tebak3%2 == 0 {
			genap++
		}
		fmt.Scan(&tebak1, &tebak2,&tebak3)
	}
	fmt.Println(genap)
}