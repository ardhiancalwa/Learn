package main
import "fmt"

func main()  {
	var i, n, bilangan, jumlah int

	fmt.Scan(&n)
	jumlah = 0
	i = 1
	for i <= n{
		fmt.Scan(&bilangan)
		jumlah += bilangan
		i++
	}
	fmt.Println(jumlah)
	
}