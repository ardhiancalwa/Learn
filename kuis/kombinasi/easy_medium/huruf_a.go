package main

import "fmt"

func main() {
	var huruf byte
	var jumA int

	fmt.Scanf("%c", &huruf)
	for huruf!= '.' {
		if huruf == 'a' {
            jumA++
        }
        fmt.Scanf("%c", &huruf)
	}
	fmt.Println(jumA)
}