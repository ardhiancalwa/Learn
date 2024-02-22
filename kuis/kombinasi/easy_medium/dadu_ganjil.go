package main

import "fmt"

func main() {
	var dadu1, dadu2, dadu3 int
	var jumGanjil int

	fmt.Scan(&dadu1, &dadu2, &dadu3)
	for dadu1+dadu2+dadu3 != 10 {
		if dadu1%2 != 0 && dadu2%2 != 0 && dadu3%2 != 0 {
			jumGanjil++
		}
		fmt.Scan(&dadu1, &dadu2, &dadu3)
	}
	fmt.Println(jumGanjil)
}
