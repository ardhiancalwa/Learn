package main
import "fmt"

func main()  {
	var i, jumlah int
	
	jumlah = 0
	i = 1
	for {
		jumlah += i
		i++
		if i > 1000 {
			break
		}
	}
	fmt.Println(jumlah) 
	
}