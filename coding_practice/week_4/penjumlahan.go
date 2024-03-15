package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(jumlah(n))
}

func jumlah(n int) int {
	if n == 1 {
        return 1
    } else {
        return n + jumlah(n-1)
    }
}

/**
TRACING TEST
n = 3
jumlah(n-1) = 2
n = 2
jumlah(n-1) = 1
n = 1
*/
