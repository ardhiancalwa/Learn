package main
import "fmt"

func main()  {
	var a, b, c float64
	var D float64

	
	fmt.Scan(&a, &b, &c)
	D = b*b - 4*a*c
	if D >= 0 {
		fmt.Println("memiliki titik potong dengan sumbu-x")
	} else {
		fmt.Println("tidak memiliki titik potong dengan sumbu-x")
	}
}