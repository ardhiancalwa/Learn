package main
import "fmt"

func main()  {
	var bilangan, jumlah int

	fmt.Scan(&bilangan)
	jumlah = 0
	for bilangan >= 0 {
		jumlah += bilangan
		fmt.Scan(&bilangan)
	}
	fmt.Println(jumlah)
	
}