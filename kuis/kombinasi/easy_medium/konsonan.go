package main

import "fmt"

func main() {
	var huruf byte
	var jumKonsonan int

	fmt.Scanf("%c", &huruf)
	for huruf != '.' {
		if huruf != 'a' && huruf != 'e' && huruf != 'i' && huruf != 'o' && huruf != 'u' {
            jumKonsonan++
        }
        fmt.Scanf("%c", &huruf)
	}
	fmt.Println(jumKonsonan)
}