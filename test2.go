package main

import "fmt"

func cariGo(kar byte) {
	var jumKar int
	var karBefore byte
	karBefore = kar 
	fmt.Scanf("%c", &kar)
	for kar != '.' {
		if karBefore == 'g' && kar == 'o' {
			jumKar++
		}
		karBefore = kar 
		fmt.Scanf("%c", &kar)
	}
	fmt.Println(jumKar)
}

func main() {
	var kar byte
	fmt.Scanf("%c", &kar)
	cariGo(kar)
}
