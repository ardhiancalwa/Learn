package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	recPrint(n, s)
}

func recPrint(n int, s string) {
	if n == 1 {
        fmt.Println(s)
    } else {
        recPrint(n-1, s)
        fmt.Println(s)
    }
}

/**
TRACING TEST
n = 3
s = WOW
print = WOW
----------
n = 2
s = WOW
print = WOW
----------
n = 1
print = WOW
----------
*/
