package main

import "fmt"

//buatlah prosedur saja

func terkecil(a, b, c int) {
	/* I.S. terdefinisi 3 bilangan bulat a, b, dan c yang berbeda
	   F.S. print bilangan terkecil dari 3 bilangan bulat a,b,c */
	var min int
	min = a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	fmt.Println(min)
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	terkecil(a, b, c)
}
