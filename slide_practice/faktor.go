package main

import "fmt"

func main() {
	var x, y int
	var faktor bool

	fmt.Scan(&x, &y)
	faktor = y%x == 0
	fmt.Println(faktor)
}