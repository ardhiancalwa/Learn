package main

import "fmt"

func main() {
	var bil, i, jumBil int
	var prima bool

	fmt.Scan(&bil)
	for i = 2; i <= bil; i++ {
		if bil%i == 0 {
			jumBil++
		}
	}
	prima = jumBil == 1 || jumBil == bil
	fmt.Println(prima)
}
