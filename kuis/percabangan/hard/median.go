package main

import "fmt"

func main() {
	var a, b, c int
	var max, min int

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
	if a > min && a < max {
		fmt.Println(a)
	} else if b > min && b < max {
		fmt.Println(b)
	} else if c > min && c < max {
		fmt.Println(c)
	}
	
}