package main

import "fmt"

func diskon(belanja int, member bool) float64 {
	var totBelanja float64
	if belanja > 100000 && member {
		totBelanja = float64(belanja) - (float64(belanja) * 0.1)
	} else if belanja > 100000 && !member {
		totBelanja = float64(belanja) - (float64(belanja) * 0.05)
	} else {
		totBelanja = float64(belanja)
	}
	return totBelanja
}

func main() {
	var belanja int
	var member bool

	fmt.Scan(&belanja, &member)
	fmt.Println(diskon(belanja, member))
}
