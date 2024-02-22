package main

import "fmt"

func main() {
	var num int

	for num = 1; num <= 100; num++ {
		if num%15 == 0 {
			fmt.Println("algoritma pemograman")
		} else if num%5 == 0 {
			fmt.Println("pemograman")
		} else if num%3 == 0 {
			fmt.Println("algoritma")
		} else {
			fmt.Println(num)
		}
	}
}
