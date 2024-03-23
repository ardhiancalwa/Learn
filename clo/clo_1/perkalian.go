package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(kali(m, n))
}

func kali(m, n int) int {
	if n == 1 {
		return m
	} else {
		return m + kali(m, n-1)
	}
}