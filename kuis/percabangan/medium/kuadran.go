package main

import "fmt"

func main() {
	var x, y, kuadran int

	fmt.Scan(&x, &y)
	if x > 0 && y > 0 {
		kuadran = 1
	} else if x < 0 && y > 0 {
		kuadran = 2
	} else if x < 0 && y < 0 {
		kuadran = 3
	} else if x > 0 && y < 0 {
		kuadran = 4
	} 

	fmt.Println(kuadran)
}
