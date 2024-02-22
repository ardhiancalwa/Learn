package main

import "fmt"

func main() {
	var skorA, skorB, skorC int
	var jumA, jumB, jumC int
	var max int

	fmt.Scan(&skorA, &skorB, &skorC)
	for skorA != 0 && skorB != 0 && skorC != 0 {
		jumA += skorA
		jumB += skorB
		jumC += skorC
		fmt.Scan(&skorA, &skorB, &skorC)
	}
    max = jumA
	if jumB > max {
		max = jumB
	} else {
		max = max
	}
	if jumC > max {
        max = jumC
    } else {
		max = max
	}
	fmt.Println(jumA, jumB, jumC, max)
}