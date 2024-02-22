package main

import (
	"fmt"
	"math"
)

func main()  {
	var jenis string
	var x, y, radius float64

	fmt.Scan(&jenis, &x, &y)
	radius = math.Sqrt(math.Pow(x,2) + math.Pow(y,2))
	if jenis == "tempur" &&  radius <= 5{
		fmt.Println("tembak")
	} else {
		fmt.Println("tidak tembak")
	}
}