package main
import "fmt"

func main()  {
	var bilangan, digit, jumlah int

	jumlah = 0
	fmt.Scan(&bilangan)
	for bilangan > 0 {
		digit = bilangan % 10
		fmt.Println(digit)
		jumlah += digit
		bilangan /= 10
	}
	fmt.Println(jumlah)
}