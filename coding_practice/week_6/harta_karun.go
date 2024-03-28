package main

import "fmt"

type koordinat struct {
    x int
	y int
}

func main() {
	var koordinat1, koordinat2 koordinat
	var i int = 1
	var isFindHarta bool = false

	fmt.Scan(&koordinat1.x, &koordinat2.y)
	for i <= 3 && isFindHarta == !true {
		if koordinat1.x*koordinat2.y == 2 {
			fmt.Println("Yay Menang")
			isFindHarta = true
		} else {
			fmt.Println("Coba Lagi")
			i++
		}
		fmt.Scan(&koordinat1.x, &koordinat2.y)

	}
}
