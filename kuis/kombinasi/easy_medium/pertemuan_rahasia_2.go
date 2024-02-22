package main

import "fmt"

func main() {
	var x, y, z int
	var i, jumPertemuan int
	const hari int = 365

	fmt.Scan(&x, &y, &z)
	for i = 1; i <= hari; i++ {
		if i%x == 0 && i%y == 0 && i%z != 0 {
			jumPertemuan++
		}
	}
	fmt.Println(jumPertemuan)
}