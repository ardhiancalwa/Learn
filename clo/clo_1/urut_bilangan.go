package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sortAscending(a, b, c)
}

func sortAscending(a, b, c int) {
	if a > b {
        if a > c {
            fmt.Println(a)
        } else {
            fmt.Println(c)
        }
    } else {
        if b > c {
            fmt.Println(b)
        } else {
            fmt.Println(c)
        }
    }
}
