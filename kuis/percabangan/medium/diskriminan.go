package main
import "fmt"

func main()  {
	var a, b, c float64
	var D float64

	
	fmt.Scan(&a, &b, &c)
	D = b*b - 4*a*c
	if D >= 0 {
		fmt.Println("memiliki akar real")
	} else {
		fmt.Println("tidak memiliki akar real")
	}
}