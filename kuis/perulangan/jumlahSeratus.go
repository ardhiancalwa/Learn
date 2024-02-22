package main
import "fmt"

func main()  {
	var bilangan, jumlah, n int

	bilangan = 0
	jumlah = 0
	n = 0
	for jumlah <= 100 {
		fmt.Scan(&bilangan)
		jumlah += bilangan
		n++
	}
	fmt.Println(jumlah, n)
}