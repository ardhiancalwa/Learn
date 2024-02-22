package main

import "fmt"

func main() {
	var A, B, i int
	var tim byte
	var stop bool

	fmt.Scanf("%c", &tim)
	A = 1
	B = 1
	stop = false
	i = 1
	for  i <= 10 && !stop {
		if tim == 'A'{
			A++
		}
		if tim == 'B'{
            B++
        }
		if A == 4 || B == 4 {
			stop = true
		}
		i++
		fmt.Scanf("%c", &tim)
	}
	fmt.Println(A, B)
	if A == 4 {
		fmt.Println("A")	
	} else if B == 4 {
		fmt.Println("B")
	} else if A >= 5 || B >= 5 {
		fmt.Println("tidak valid")
	} 
}