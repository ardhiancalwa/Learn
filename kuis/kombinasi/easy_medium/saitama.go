package main

import "fmt"

func main() {
	var ying, yang int
	var A, B float64
	var i, hari int

	A = 0.0
	B = 0.0
	fmt.Scan(&hari)
	for i = 1; i <= hari; i++ {
		fmt.Scan(&ying, &yang)
		if yang < ying {
			A += 0.25
		} else if yang > ying {
			B += 0.15
		}
		if i%3 == 0 {
			A -= A * 0.05
		}
	}
	fmt.Println(A, B)
}
