package main

import "fmt"

func main() {
	var x, y, z, i int
	var result int
	const day int = 365

	fmt.Scan(&x, &y, &z)
	for i = 1; i <= day; i++ {
		if i%x == 0 && i%y == 0{
			if i%z != 0 {
				result++
			}
		} 
	}
	fmt.Println(result)
}