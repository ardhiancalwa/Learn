package main

import "fmt"

func main() {
	var aktivitas string
	var belajar, main int

	fmt.Scan(&aktivitas)
	for aktivitas != "tidur" {
		if aktivitas == "belajar" {
			belajar++
		} else if aktivitas == "main" {
			main++
        }
		fmt.Scan(&aktivitas)
	}
	fmt.Println(belajar, main)
}