package main

import "fmt"

func lowToUpper(kar byte) byte {
	var toUpper byte
	if kar >= 'a' && kar <= 'z' {
		toUpper = kar - 32
	}
	return toUpper
}

func main() {
	var kar byte
	fmt.Scanf("%c",&kar)
	fmt.Printf("%c", lowToUpper(byte(kar)))
}
