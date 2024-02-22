package main
import "fmt"

func main()  {
	var i, bilangan int
	var jumlah float64
	
	jumlah = 1.0
	i = 1
	for {
		fmt.Scan(&bilangan)
		i++
		if bilangan == 0 {
			break
		}
		jumlah *= 1.0 / float64(bilangan)
	}
	fmt.Printf("%.5f",jumlah)
	
}