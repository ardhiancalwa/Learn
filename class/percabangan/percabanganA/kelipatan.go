package main

import "fmt"

func main() {
	var bil int
	var kelipatan string

	fmt.Scan(&bil)
	if bil%3 == 0 && bil%5 == 0 {
		kelipatan = "Kelipatan 3 \nKelipatan 5"
	} else if bil%3 == 0 {
		kelipatan = "Kelipatan 3"
	} else if bil%5 == 0 {
		kelipatan = "Kelipatan 5"
	}
	fmt.Println(kelipatan)
}
