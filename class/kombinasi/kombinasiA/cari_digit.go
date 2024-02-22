package main

import "fmt"

func main() {
	var x, n, digit int
	var nilai bool

	fmt.Scan(&x, &n)
	if x <= n {
		for n > 0 {
			digit = n % 10
			if digit == x {
				nilai = true
			}
			n /= 10
		}
	} else {
		nilai = false
	}
	fmt.Println(nilai)
}
