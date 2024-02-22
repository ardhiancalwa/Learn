package main

import "fmt"

func main() {
	var bil1, bil2 int
	var jumBil int

	fmt.Scan(&bil1, &bil2)
	for bil1 != 0 && bil2 != 0 {
		if bil1%2 != 0 && bil2%2 == 0 {
			jumBil += bil1 + bil2
		}
		fmt.Scan(&bil1, &bil2)
	}
	fmt.Println(jumBil)
}
