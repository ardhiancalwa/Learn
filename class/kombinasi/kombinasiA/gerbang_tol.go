	package main

	import "fmt"

	func main() {
		var tipe byte
		var A, B, C int

		fmt.Scanf("%c", &tipe)

		for tipe == 'A' || tipe == 'B' || tipe == 'C' {
			if tipe == 'A' {
				A++
			}
			if tipe == 'B' {
				B++
			}
			if tipe == 'C' {
				C++
			}
			fmt.Scanf(" %c", &tipe)
		}

		fmt.Println("Tipe A =", A)
		fmt.Println("Tipe B =", B)
		fmt.Println("Tipe C =", C)
	}
