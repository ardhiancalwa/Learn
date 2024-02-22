package main

import "fmt"

func main() {
	var N, ember, tangki int
	var penuh bool

	tangki = 0
	ember = 0
	fmt.Scan(&N)
	for {
		fmt.Scan(&ember)
		tangki += ember
		penuh = tangki >= N
		fmt.Print(penuh)
		if tangki >= N {
			break
		}
	}
}
