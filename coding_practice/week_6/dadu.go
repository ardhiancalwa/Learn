package main

import "fmt"

type HasilGame struct {
	TotalKemunculanGenap int
}

func game() HasilGame {
	var d1, d2, d3 HasilGame
	var genap HasilGame
	fmt.Scan(&d1.TotalKemunculanGenap, &d2.TotalKemunculanGenap, &d3.TotalKemunculanGenap)
	for d1.TotalKemunculanGenap+d2.TotalKemunculanGenap+d3.TotalKemunculanGenap != 15 {
		if d1.TotalKemunculanGenap%2 == 0 && d2.TotalKemunculanGenap%2 == 0 && d3.TotalKemunculanGenap%2 == 0 {
			genap.TotalKemunculanGenap++
		}
		fmt.Scan(&d1.TotalKemunculanGenap, &d2.TotalKemunculanGenap, &d3.TotalKemunculanGenap)
	}
	return genap
}

func main() {
	hasil := game()
	fmt.Println(hasil.TotalKemunculanGenap)
}
