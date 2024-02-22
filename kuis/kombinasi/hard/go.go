package main

import "fmt"

func main() {
	var huruf byte
	var jumGo int

	fmt.Scanf("%c", &huruf)
	for huruf != '.' {
		if huruf == 'g' {
			fmt.Scanf("%c", &huruf)
			if huruf == 'o' {
				jumGo++
			}
		} else {
			fmt.Scanf("%c", &huruf)
		}
	}
	fmt.Println(jumGo)
}
