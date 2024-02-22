package main
import "fmt"

func main()  {
	var a, jumlah, i int
	jumlah = 0
	for i = 1; i <= 100; i++ {
		fmt.Scan(&a)
		jumlah += a
	}
	fmt.Println(jumlah)
	
}