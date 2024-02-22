package main

import "fmt"

func main() {
	var huruf byte
	var jumVokal, jumKonsonan int

	fmt.Scanf("%c", &huruf)
	for huruf!= '.' {
		if huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
            jumVokal++
        } else {
            jumKonsonan++
        }
        fmt.Scanf("%c", &huruf)
	}
	fmt.Println(jumVokal, jumKonsonan)
}