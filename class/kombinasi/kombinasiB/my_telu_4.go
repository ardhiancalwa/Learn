package main

import "fmt"

func main() {
	var poin, jumPoin, gold, platinum, silver int
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
				gold++
			} else if poin >= 100 {
				platinum++
			} else if poin >= 50 {
				silver++
			}
			fmt.Scan(&poin)
			jumPoin += poin
		}
	}
	fmt.Println("Gold: ", gold)
	fmt.Println("Platinum: ", platinum)
	fmt.Println("Silver: ", silver)
}