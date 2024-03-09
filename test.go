package main

import "fmt"

func konsekutif(bilangan int) {
	konsekutif := false

	for bilangan >= 10 {
		if konsekutif && (bilangan%10)-(bilangan/10%10) != 1 || ((bilangan/10%10)-(bilangan%10) != 1) {
			konsekutif = true
			break
		}
		bilangan /= 10
	}

	if konsekutif {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	konsekutif(bilangan)
}
