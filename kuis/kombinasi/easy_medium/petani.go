package main

import "fmt"

func main() {
	var i, n, apel, jeruk, mangga int
	var buah string

	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&buah)
		if buah == "apel" {
			apel++
		} else if buah == "jeruk" {
			jeruk++
        } else if buah == "mangga" {
			mangga++
		}
	}
	fmt.Println(apel, jeruk, mangga)
}