package main

import "fmt"

func main() {
	var x, y, i int
	var jumKetemu int
	const hari int = 365

	fmt.Scan(&x, &y)
	for i = 1; i <= hari; i++ {
		if i%x == 0 && i%y != 0 {
			jumKetemu++
		}
	}
	fmt.Println(jumKetemu)
}