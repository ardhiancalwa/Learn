package main

import "fmt"

func main() {
	var P, S, X, Y, B int
	var jumP, jumS int

	fmt.Scan(&P, &S, &X, &Y, &B)
	B *= 30
	for i := 1; i <= B; i++ {
		if i%X != 0 && i%Y != 0 {
			if i%2 == 0 {
				jumS += S
			} else if i%2 != 0 {
				jumP += P
			}
		} else {
			jumP += 0
			jumS += 0
		}
	}
	fmt.Println(jumP, jumS)
}
