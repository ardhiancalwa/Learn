package main

import "fmt"

func main() {
	var a, b, c int
	var min int

	fmt.Scan(&a, &b, &c)
	min = a
	if b <= min {
		min = b
	} else {
		min = min 
	}
	if c <= min {
        min = c
    } else {
        min = min 
    }
	fmt.Println(min)
	
}