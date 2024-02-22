package main

import "fmt"

func main() {
	var prima bool
	var bil, i, jumFaktor int

	fmt.Scan(&bil)
	for i = 1; i <= bil; i++ {
		if bil%i == 0 {
			jumFaktor++
		}
	}
	prima = jumFaktor == 3
	fmt.Println(prima)
}