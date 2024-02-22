package main

import "fmt"

func main() {
	var huruf byte
	var jumAI int

	fmt.Scanf("%c", &huruf)
	for huruf!= '.' {
		if huruf == 'a' || huruf == 'i' {
            jumAI++
        }
        fmt.Scanf("%c", &huruf)
	}
	fmt.Println(jumAI)
}