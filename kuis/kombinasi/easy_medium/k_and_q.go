package main

import "fmt"

func main() {
	var huruf byte
	var hasil bool
	var i int

	for i = 1; i <= 10; i++ {
		fmt.Scanf("%c", &huruf)
		if huruf == 'k' || huruf == 'q' {
			hasil = true
			break
		}
	}
	
	fmt.Println(hasil)
}
