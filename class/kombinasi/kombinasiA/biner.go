package main

import "fmt"

func main() {
	var bil, digit int
	var  hasil string
	fmt.Scan(&bil)

	for bil > 0 {
		digit = bil % 2
		if digit == 0 {
			hasil = "0" + hasil
		} else if digit == 1 {
			hasil = "1" + hasil
		}
		bil /= 2
	}
	fmt.Print(hasil)
}
