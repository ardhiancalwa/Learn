package main

import "fmt"

func main() {
	var poin int
	
	fmt.Scan(&poin)
	if poin > 200 {
		fmt.Println("Gold User")
	} else if poin >= 100 {
		fmt.Println("Platinum User")
	} else if poin >= 50 {
		fmt.Println("Silver User")
	}
}
