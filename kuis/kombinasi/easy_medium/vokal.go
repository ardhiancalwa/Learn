package main

import "fmt"

func main() {
	var huruf byte
	var vokal int

	fmt.Scanf("%c", &huruf)
	for huruf != '.' {
		if huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
			vokal++
		}
		fmt.Scanf("%c", &huruf)
	}
	fmt.Println(vokal)
}
