package main

import (
	"fmt"
	"math"
)

func panjangX(R, T float64) float64 {
	var X float64
	T = T * math.Pi / 180
	X = R * math.Cos(T)
	return X
}
func panjangY(R, T float64) float64 {
	var Y float64
	T = T * math.Pi / 180
	Y = R * math.Sin(T)
	return Y
}

func main() {
	var R, T float64

	fmt.Scan(&R, &T)
	fmt.Printf("%.2f %.2f", panjangX(R, T), panjangY(R, T))
}
