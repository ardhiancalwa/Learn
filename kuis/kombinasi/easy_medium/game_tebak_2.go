package main

import "fmt"

func main() {
	var pemain1, pemain2, pemain3 int
	var tebak1, tebak2, tebak3 int

	fmt.Scan(&tebak1, &tebak2, &tebak3)
	if tebak1 >= 1 && tebak1 <= 5 && tebak2 >= 1 && tebak2 <= 5 && tebak3 >= 1 && tebak3 <= 5 {
		for tebak1 != 0 && tebak2 != 0 && tebak3 != 0 {
			if (tebak1 != tebak2) && (tebak1 != tebak3) && (tebak2 == tebak3) {
				pemain1++
			}
			if (tebak2 != tebak1) && (tebak2 != tebak3) && (tebak1 == tebak3) {
				pemain2++
			}
			if (tebak3 != tebak1) && (tebak3 != tebak2) && (tebak1 == tebak2) {
				pemain3++
			}
			fmt.Scan(&tebak1, &tebak2, &tebak3)
		}
	}
	fmt.Println(pemain1, pemain2, pemain3)
}
