package main

import "fmt"

func main() {
	var bil, digit, max int

	fmt.Scan(&bil)
	max = 0
	for bil > 0 {
		digit = bil % 10
		if digit > max {
			max = digit
		}
		bil = bil / 10
	}
	fmt.Println(max)
}