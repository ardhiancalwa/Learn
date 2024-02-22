package main

import "fmt"

func main() {
	var a, b, c int
	var max int

	fmt.Scan(&a, &b, &c)
	max = a
	if b >= max {
		max = b
	} else {
		max = max 
	}
	if c >= max {
        max = c
    } else {
        max = max 
    }
	fmt.Println(max)
	
}