package main
import "fmt"

func main()  {
	var bil, jumlah int

	fmt.Scan(&bil)
	jumlah = 0
	for bil > 0 {
		jumlah += bil 
	}
	fmt.Println(jumlah)
}