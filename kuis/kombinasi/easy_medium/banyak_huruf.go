package main

import "fmt"

func main() {
	var huruf byte
	var jumHuruf int

	fmt.Scanf("%c", &huruf)
	for huruf!= '.' {
		if huruf != 'a' && huruf != 'i' {
			jumHuruf++
		}
		fmt.Scanf("%c", &huruf)
	}
	fmt.Println(jumHuruf)
}