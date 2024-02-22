package main

import "fmt"

func main() {
	var suara string
	var i, hari int
	var lima_suara bool
	var JA, BE, CI, JE, LO bool

	fmt.Scan(&hari)
	for i = 1; i <= hari; i++ {
		fmt.Scan(&suara)
		if suara == "JA" {
			JA = true
		} else if suara == "BE" {
			BE = true
		} else if suara == "CI" {
			CI = true
		} else if suara == "JE" {
			JE = true
		} else if suara == "LO" {
			LO = true
		}
	}
	if JA && BE && CI && JE && LO {
        lima_suara = true
    } else {
        lima_suara = false
    }
	fmt.Println(lima_suara)
}