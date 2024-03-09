package main

import "fmt"

func denom_uang(uang int) {
	var lembar10, lembar2, lembar1 int
	lembar10 = uang / 10000
	uang = uang % 10000
	lembar2 = uang / 2000
	uang = uang % 2000
	lembar1 = uang / 1000
	fmt.Println(lembar10, "lembar 10000")
	fmt.Println(lembar2, "lembar 2000")
	fmt.Println(lembar1, "lembar 1000")
}

func main() {
	var uang int
	fmt.Scan(&uang)
	denom_uang(uang)
}
