package main

import "fmt"

func main() {
	var k, p int
	fmt.Scan(&k, &p)
	hitungSewaBus(k, p)
}

func hitungSewaBus(k, p int) {
	var jumBus int
	jumBus = p / k
	if p%k > 0 {
		jumBus++
	}
	fmt.Println(jumBus)
}
