package main

import "fmt"

func konversi(celcius float64) {
	var reamur, fahrenheit, kelvin float64
	reamur = (celcius * 4) / 5
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = celcius + 273.15
	fmt.Println(reamur,"R", fahrenheit,"F", kelvin,"K")
}

func main() {
	var celcius float64
	fmt.Scan(&celcius)
	konversi(celcius)
}
