package main

import (
	"fmt"
	"math"
)

func main()  {
	var bil1, bil2, hasil int
	var pangkat string

	fmt.Scan(&bil1, &bil2, &hasil)
	if math.Pow(float64(bil1), float64(bil2)) == float64(hasil) {
		pangkat = "benar"
	} else {
		pangkat = "salah"
	}
	fmt.Println(pangkat)
}