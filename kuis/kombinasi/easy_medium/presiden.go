package main

import "fmt"

func main() {
	var x, y, z int
	var jumX, jumY, jumZ int
	var i, n int
	var positif, netral, negatif byte

	positif = '+'
	netral = '='
	negatif = '-'

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&x, &y, &z)
		jumX += x
		jumY += y
		jumZ += z
	}
	if jumX > jumY && jumX > jumZ {
		fmt.Printf("%c %d", positif, jumX)
    } else if jumY > jumX && jumY > jumZ {
		fmt.Printf("%c %d", netral, jumY)
	} else if jumZ >jumX && jumZ > jumY {
		fmt.Printf("%c %d", negatif, jumZ)
	}
}