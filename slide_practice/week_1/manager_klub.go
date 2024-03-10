package main

import "fmt"

func main() {
	var hasil string
	var hari, kalah int

	for hari = 1; hari <= 5; hari++ {
		fmt.Scan(&hasil)
		if hasil == "kalah" {
			kalah++
		}
	}
	if kalah == 5 {
		fmt.Println("dipecat")
	} else {
		fmt.Println("tidak dipecat")
	}

}