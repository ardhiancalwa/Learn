package main

import "fmt"

func main() {
	var poin, jumPoin, gold, platinum, silver int
	var poinGold, poinPlatinum, poinSilver int
	var rataGold, rataPlatinum, rataSilver float64
	fmt.Scan(&poin)

	jumPoin = 0
	gold = 0
	platinum = 0
	silver = 0
	if poin > 500 {
		gold = 1
	}
	if jumPoin <= 500 {
		jumPoin += poin
		for jumPoin <= 500 {
			if poin > 200 {
				poinGold += poin
				gold++
			} else if poin >= 100 {
				poinPlatinum += poin
				platinum++
			} else if poin >= 50 {
				poinSilver += poin
				silver++
			}
			fmt.Scan(&poin)
			jumPoin += poin
		}
	}

	rataGold = float64(poinGold) / float64(gold)
	rataPlatinum = float64(poinPlatinum) / float64(platinum)
	rataSilver = float64(poinSilver) / float64(silver)

	if gold == 0 {
		rataGold = 0
	}
	if platinum == 0 {
		rataPlatinum = 0
	}
	if silver == 0 {
		rataSilver = 0
	}
	fmt.Println("Gold: ", rataGold)
	fmt.Println("Platinum: ", rataPlatinum)
	fmt.Println("Silver: ", rataSilver)
}
