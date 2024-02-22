package main

import "fmt"

func main() {
	var a1, a2, a3 int
	var sahabat, tidak int

	fmt.Scan(&a1, &a2, &a3)
	for a1 != 0 && a2 != 0 && a3 != 0 {
		if (a1 != a2 && a2 != a3 && a1 != a3) {
			if a1+a2 == a3 || a1+a3 == a2 || a2+a3 == a1 {
				sahabat++
			} else {
				tidak++
			}
		} else {
			tidak++
		}
		fmt.Scan(&a1, &a2, &a3)
	}
	fmt.Println(sahabat, tidak)
}