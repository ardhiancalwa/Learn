package main
import "fmt"

func main()  {
	var keliling, radius float64
	const pi float64 = 3.14
	fmt.Scan(&radius)

	keliling = 2 * pi * radius 
	fmt.Println(keliling)
}