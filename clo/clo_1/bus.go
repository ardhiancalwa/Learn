package main

import "fmt"

func main() {
	var k, p int 
	fmt.Scan(&k, &p)
	hitungSewaBus(k, p)
}

func hitungSewaBus(k, p int) {
	var total int
	if p/k != 0 {
		total = (p/k) + 1
	} else {
		total = (p/k)
	}
	fmt.Println(total)
}
