package main

import "fmt"

func main() {
	var pemain1, pemain2, pemain3 int
	var tebak1, tebak2, tebak3 int

	fmt.Scan(&tebak1, &tebak2, &tebak3)
	if tebak1 >= 1 && tebak1 <= 10 && tebak2 >= 1 && tebak2 <= 10 && tebak3 >= 1 && tebak3 <= 10 {
		for tebak1 != 0 && tebak2 != 0 && tebak3 != 0 {
			if (tebak1%2 != 0 && tebak2%2 == 0 && tebak3%2 == 0) || (tebak1%2 == 0 && tebak2%2 != 0 && tebak3%2 != 0) {
				pemain1++
			}
			if (tebak1%2 == 0 && tebak2%2 != 0 && tebak3%2 == 0) || (tebak1%2 != 0 && tebak2%2 == 0 && tebak3%2 != 0) {
				pemain2++
			}
			if (tebak1%2 == 0 && tebak2%2 == 0 && tebak3%2 != 0) || (tebak1%2 != 0 && tebak2%2 != 0 && tebak3%2 == 0) {
				pemain3++
			}
			fmt.Scan(&tebak1, &tebak2, &tebak3)
		}
	}
	fmt.Println(pemain1, pemain2, pemain3)
}
