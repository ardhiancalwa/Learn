package main

import "fmt"

func main() {
	var diskon, cashback, totBelanja int
	var i, pembeli int

	fmt.Scan(&pembeli)
	for i = 1; i <= pembeli; i++ {
		fmt.Scan(&totBelanja)
		if (totBelanja%3 == 0 && totBelanja%100 != 0) || (totBelanja%300 == 0) {
			cashback++
		}
		if (totBelanja%4 == 0 && totBelanja%100 != 0) || (totBelanja%400 == 0) {
			diskon++
		}
	}
	fmt.Println(diskon, cashback)
}
