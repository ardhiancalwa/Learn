package main
import "fmt"

func main()  {
	var bilangan, d1, d2, d3, d4 int
	var hasil bool

	fmt.Scan(&bilangan)
	d1 = bilangan / 1000
	d2 = bilangan % 1000 / 100
	d3 = bilangan % 100 / 10
	d4 = bilangan % 10

	hasil = (bilangan >= 1000 && bilangan <= 9999) && d1 < d2 && d2 < d3 && d3 < d4
	fmt.Println(d1, d2, d3, d4, hasil)
	
}