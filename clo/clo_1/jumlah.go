package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	fmt.Println(jumlah(bil, 0))

}

func jumlah(n, hasil int) int {
	if n == 0 {
		return hasil
	} else {
		return jumlah(n/10, hasil+n%10)
	}
}
