package main

import "fmt"

// Procedure untuk mengurutkan dua bilangan secara membesar
func mengurutkan(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for a != b {
		mengurutkan(&a, &b)
		fmt.Println(a, b)
		fmt.Scan(&a, &b)
	}
}
