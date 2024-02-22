package main

import "fmt"

func main() {
	var kondisi string
	var jumBarang int
	var bagus, cacat int

	fmt.Scan(&kondisi, &jumBarang)
	for kondisi != "selesai" && jumBarang != 0 {
		if kondisi == "bagus" {
            bagus += jumBarang
        }
        if kondisi == "cacat" {
            cacat += jumBarang
        }
        fmt.Scan(&kondisi, &jumBarang)
	}
	fmt.Println(bagus, cacat)
}