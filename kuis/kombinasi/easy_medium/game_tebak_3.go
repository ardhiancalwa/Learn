package main

import "fmt"

func main() {
	var tebak1, tebak2, tebak3 byte
	var pemain1, pemain2, pemain3 int

	fmt.Scanf("%c%c%c", &tebak1, &tebak2, &tebak3)
	for tebak1 != '.' && tebak2 != '.' && tebak3 != '.' {
		if (tebak1 != tebak2) && (tebak1 != tebak3) && (tebak2 == tebak3) {
			pemain1++
		}
		if (tebak2 != tebak1) && (tebak2 != tebak3) && (tebak1 == tebak3) {
			pemain2++
		}
		if (tebak3 != tebak1) && (tebak3 != tebak2) && (tebak1 == tebak2) {
			pemain3++
		}
		fmt.Scanf("%c%c%c", &tebak1, &tebak2, &tebak3)
	}
	fmt.Println(pemain1, pemain2, pemain3)
}