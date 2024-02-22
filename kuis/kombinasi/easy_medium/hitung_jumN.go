package main

import "fmt"

func main() {
	var i, n int
	var jumN, total int

	jumN = 1
	fmt.Scan(&n)
	for i = 1; jumN <= n; i++ {
		if i%2 != 0 {
            total += i
			jumN++
        }
	}
	fmt.Println(total)
}