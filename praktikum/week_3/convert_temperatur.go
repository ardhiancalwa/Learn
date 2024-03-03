package main

import (
	"fmt"
)

func reamur(C float64) float64 {
	return C * 4 / 5
}

func fahrenheit(C float64) float64 {
	return (C * 9 / 5) + 32
}

func kelvin(C float64) float64 {
	return C + 273
}

func main() {
	var C, R, F, K float64
	var i, step, tempAkhir int
	fmt.Scan(&C, &tempAkhir, &step)

	fmt.Printf("%10s %10s %10s %10s\n", "Celsius", "Reamur", "Fahrenheit", "Kelvin")

	// fmt.Println("   Celsius     Reamur Fahrenheit     Kelvin")
	for i = int(C); i <= tempAkhir; i += step {
		R = reamur(C)
		F = fahrenheit(C)
		K = kelvin(C)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
		C = float64(i) + float64(step)
	}

}
