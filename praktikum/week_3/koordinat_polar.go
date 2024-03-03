package main

import (
	"fmt"
	"math"
)

func panjangX(R, T float64) float64 {
	var X float64
	var derajatSudutX float64
	derajatSudutX = T * math.Pi / 180
	X = R * math.Cos(derajatSudutX)
	return X
}
func panjangY(R, T float64) float64 {
	var Y float64
	var derajatSudutY float64
	derajatSudutY = T * math.Pi / 180
	Y = R * math.Sin(derajatSudutY)
	return Y
}

func main() {
	var R, T float64

	fmt.Scan(&R, &T)
	fmt.Printf("%.2f %.2f", panjangX(R, T), panjangY(R, T))
}
