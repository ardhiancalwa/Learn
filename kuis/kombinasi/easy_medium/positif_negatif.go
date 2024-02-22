package main

import "fmt"

func main() {
	var bil int
	var positif, negatif int

	fmt.Scan(&bil)
	for bil != 0 {
		if bil > 0 {
			positif++
		} else if bil < 0 {
			negatif++
        }
		fmt.Scan(&bil)
	}
	fmt.Println(positif, negatif)
}