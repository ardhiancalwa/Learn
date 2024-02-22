package main

import "fmt"

func main() {
	var x, yAwal, yBaru, z int
	var totTabungan int
	var i int

	fmt.Scan(&x, &yAwal, &z)
	yBaru = 0
	for i = 1; i <= z; i++ {
		if i%2 == 0 || i%3 == 0 {
			if totTabungan == 0 {
				totTabungan = x
			} else {
				yBaru += yAwal
				totTabungan += x + yBaru
			}
		}
	}
	fmt.Println(totTabungan)
}
