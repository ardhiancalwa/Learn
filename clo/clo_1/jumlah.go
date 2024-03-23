package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(jumlahTiapDigit(n, 0))
}

func jumlahTiapDigit(n, hasil int) int {
	if n == 0 {
		return hasil
	} else {
		return jumlahTiapDigit(n/10, hasil+n%10)
	}
}
