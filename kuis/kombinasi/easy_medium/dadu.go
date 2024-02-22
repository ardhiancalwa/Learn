package main

import "fmt"

func main() {
	var dadu1, dadu2, dadu3 int
	var jumGenap, jumGanjil, jumDadu int

	fmt.Scan(&dadu1, &dadu2, &dadu3)
	for dadu1+dadu2+dadu3 <= 15 {
		if dadu1%2 == 0 && dadu2%2 == 0 && dadu3%2 == 0 {
			jumGenap++
		}
		if dadu1%2 != 0 && dadu2%2 != 0 && dadu3%2 != 0 {
			jumGanjil++
		}
		fmt.Scan(&dadu1, &dadu2, &dadu3)
	}
	jumDadu = jumGenap+jumGanjil
	fmt.Println(jumDadu)
}
