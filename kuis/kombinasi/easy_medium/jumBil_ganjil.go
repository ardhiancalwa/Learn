package main

import "fmt"

func main() {
	var bil, jumBil int

	fmt.Scan(&bil)
	for bil >= 0 {
		if bil%2 != 0 && bil%3 == 0 {
			jumBil += bil
		}
		fmt.Scan(&bil)
	}
	fmt.Println(jumBil)
}